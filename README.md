# Liferay Upgrade Action

GitHub Action to create a new branch and pull request if a new Liferay version is available

## Requirements

This action create a branch, push changes and create a pull request. So make sure to give proper permissions to GitHub Actions in your repository:

- Go to `Settings > Actions > General > Workflow Permissions`
- Select `Read and write permissions`
- Check `Allow GitHub Actions to create and approve pull requests`

More information in [GitHub Actions documentation](https://docs.github.com/en/repositories/managing-your-repositorys-settings-and-features/enabling-features-for-your-repository/managing-github-actions-settings-for-a-repository#configuring-the-default-github_token-permissions).

## Usage

You can use this action in a [GitHub Actions Workflow](https://help.github.com/en/articles/about-github-actions) by a adding a YAML file under `.github/workflows/` with the following content:

```yaml
name: liferay-auto-upgrade
run-name: Liferay Auto Upgrade

on:
  schedule:
    # https://crontab.guru/every-monday
    - cron: '0 0 * * MON'

jobs:
  liferay-upgrade:
    permissions:
      contents: write # to push changes
      pull-requests: write # to create pull requests
    runs-on: ubuntu-latest
    steps:
      - name: Liferay Upgrade
        uses: lgdd/liferay-upgrade-action@v1
        with:
          java-distribution: 'zulu'
          java-version: '11'
```

In this example we run the every monday to follow Liferay weekly release schedule. Of course, you can change the frequency as well as the event list you want this action to be triggered by.

More information about [Github Actions Events](https://docs.github.com/en/actions/using-workflows/events-that-trigger-workflows).

## License

[MIT](LICENSE)
