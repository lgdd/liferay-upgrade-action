package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"

	gh "github.com/cli/go-gh/v2"
)

func main() {
	printExpectedEnvVariables("WORKSPACE_DIRECTORY", "GITHUB_REF_NAME",
		"LFR_CURRENT_PRODUCT_NAME", "LFR_LATEST_PRODUCT_NAME",
		"LFR_LATEST_PRODUCT_VERSION_NAME", "NO_UPGRADE_BRANCH")

	workspacePath := os.Getenv("WORKSPACE_DIRECTORY")
	mainBranchName := os.Getenv("GITHUB_REF_NAME")
	currentProductName := os.Getenv("LFR_CURRENT_PRODUCT_NAME")
	latestProductName := os.Getenv("LFR_LATEST_PRODUCT_NAME")
	latestProductVersionName := os.Getenv("LFR_LATEST_PRODUCT_VERSION_NAME")
	upgradeBranchName := os.Getenv("UPGRADE_BRANCH_NAME")
	noUpgradeBranch, _ := strconv.ParseBool(os.Getenv("NO_UPGRADE_BRANCH"))

	gitConfigUser()
	gitFetchAll()

	if currentProductName == latestProductName {
		fmt.Println("Liferay workspace is already set to the latest " + latestProductName)
		os.Exit(0)
	}

	gitSwitchBranch(noUpgradeBranch, upgradeBranchName)
	gitMergeMainIntoUpgrade(mainBranchName, upgradeBranchName)

	updateGradleProperties(workspacePath, currentProductName, latestProductName)
	gradleBuildResult := runGradleAndGetResult()
	gitCommitAndPush(workspacePath, upgradeBranchName)

	pullRequestTitle := "[Liferay Upgrade] To " + latestProductVersionName
	pullRequestBody := createPullRequestBodyMarkdown(gradleBuildResult)
	createOrEditPullRequest(mainBranchName, upgradeBranchName, pullRequestTitle, pullRequestBody)
}

func createPullRequestBodyMarkdown(gradleBuildResult string) string {
	var pullRequestBodyBuilder strings.Builder

	if strings.Contains(gradleBuildResult, "BUILD FAILED") {
		pullRequestBodyBuilder.WriteString("❌ Build failed with output:")
	} else {
		pullRequestBodyBuilder.WriteString("✅ Build succeeded with output:")
	}

	pullRequestBodyBuilder.WriteString("```")
	pullRequestBodyBuilder.WriteString(gradleBuildResult)
	pullRequestBodyBuilder.WriteString("```")

	return pullRequestBodyBuilder.String()
}

func runGradleAndGetResult() string {
	runCmd("./gradlew", "build", "-S", ">", "gradle-out.txt", "2>", "gradle-err.txt")
	var gradleResultBuilder strings.Builder
	gradleResultBuilder.WriteString(getFileContentAsString("gradle-out.txt"))
	gradleResultBuilder.WriteString("\n")
	gradleResultBuilder.WriteString(getFileContentAsString("gradle-err.txt"))
	os.Remove("gradle-out.txt")
	os.Remove("gradle-err.txt")
	return gradleResultBuilder.String()
}

func updateGradleProperties(path, currentProductName, latestProductName string) {
	read, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	newContents := strings.Replace(string(read), currentProductName, latestProductName, -1)

	err = os.WriteFile(path, []byte(newContents), 0)
	if err != nil {
		panic(err)
	}
}

func gitConfigUser() {
	runCmd("git", "config", "user.name", "github-actions[bot]")
	runCmd("git", "config", "user.email", "41898282+github-actions[bot]@users.noreply.github.com")
}

func gitFetchAll() {
	runCmd("git", "fetch", "--all")
	runCmd("git", "pull", "--all")
}

func gitMergeMainIntoUpgrade(mainBranchName, upgradeBranchName string) {
	runCmd("git", "merge", "origin/"+mainBranchName, "-Xtheirs", "-m", "\"chore: merge '"+mainBranchName+"' into '"+upgradeBranchName+"'\"", "--allow-unrelated-histories")
}

func gitSwitchBranch(noUpgradeBranch bool, upgradeBranchName string) {
	if noUpgradeBranch {
		runCmd("git", "switch", "-c", upgradeBranchName)
	} else {
		runCmd("git", "switch", upgradeBranchName)
	}

	cmd := exec.Command("git", "pull", "origin", upgradeBranchName, "--rebase")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil {
		fmt.Println(err.Error())
	}
}

func gitCommitAndPush(path, upgradeBranchName string) {
	runCmd("git", "add", path)

	cmd := exec.Command("git", "diff-index", "--quiet", "HEAD")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil {
		runCmd("git", "commit", "-m", "chore: upgrade liferay cloud images")
		runCmd("git", "push", "-u", "origin", upgradeBranchName)
	}
}

func createOrEditPullRequest(mainBranchName, upgradeBranchName, title, body string) {
	fmt.Println("Run pr edit " + upgradeBranchName)
	stdoutBuffer, stderrBuffer, err := gh.Exec("pr", "edit", upgradeBranchName, "-t", title, "-b", body)
	if err != nil {
		fmt.Println("error: " + stderrBuffer.String())
		// pr edit fails, so no pr exists therefore we can run pr create
		createPullRequest(mainBranchName, upgradeBranchName, title, body)
	} else {
		pullRequestUrl := strings.TrimSuffix(stdoutBuffer.String(), "\n")
		fmt.Println("Run pr reopen " + pullRequestUrl)
		_, stderrBuffer, err := gh.Exec("pr", "reopen", pullRequestUrl)
		if err != nil {
			fmt.Println("error: " + stderrBuffer.String())
			// pr reopen fails, so pr lost track of the branch therefore we can run pr create
			createPullRequest(mainBranchName, upgradeBranchName, title, body)
		} else {
			// pr reopen works, let's comment
			gh.Exec("pr", "comment", pullRequestUrl, "--body", body)
		}
	}
}

func createPullRequest(mainBranchName, upgradeBranchName, title, body string) {
	fmt.Println("Run pr create --base " + mainBranchName + " --head " + upgradeBranchName)
	_, stderrBuffer, err := gh.Exec("pr", "create", "--base", mainBranchName, "--head", upgradeBranchName, "-t", title, "-b", body)
	if err != nil {
		fmt.Println("error: " + stderrBuffer.String())
		panic(err)
	}
}

func printExpectedEnvVariables(keys ...string) {
	for _, key := range keys {
		fmt.Println(os.Getenv(key))
	}
}

func runCmd(command string, args ...string) {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil {
		panic(err)
	}
}

func getFileContentAsString(path string) string {
	fileContent, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	}

	return string(fileContent)
}
