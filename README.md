# Liferay Upgrade Action

`lgdd/liferay-upgrade-action@v2` create a new branch and pull request if a new Liferay version is available.

This action uses another action you might find useful: https://github.com/lgdd/get-liferay-info-action

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

permissions:
  contents: write
  pull-requests: write

jobs:
  liferay-upgrade:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - uses: lgdd/liferay-upgrade-action@v2
        with:
          java-distribution: 'zulu'
          java-version: '11'
```

In this example we run the every monday to follow Liferay weekly release schedule. Of course, you can change the frequency as well as the event list you want this action to be triggered by.

More information about [Github Actions Events](https://docs.github.com/en/actions/using-workflows/events-that-trigger-workflows).

## v2

In v1, the checkout step was done by default inside that action. Even if you could disable it with an input, it doesn't feel like a good practice to include that in a custom action.

So **In v2 you need to add the checkout step first**:

```diff
steps:
+     - uses: actions/checkout@v3
      - uses: lgdd/liferay-upgrade-action@v2
        with:
          java-distribution: 'zulu'
          java-version: '11'
```

If you were already using the checkout action in v1, you can now remove the input in v2:
```diff
steps:
      - uses: actions/checkout@v3
      - uses: lgdd/liferay-upgrade-action@v2
        with:
          java-distribution: 'zulu'
          java-version: '11'
-         checkout: false
```

## License

[MIT](LICENSE)
