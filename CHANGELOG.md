# Changelog

## [2.0.1](https://github.com/lgdd/liferay-upgrade-action/compare/v2.0.0...v2.0.1) (2023-06-06)


### Bug Fixes

* gradle build & getting output ([d534a7d](https://github.com/lgdd/liferay-upgrade-action/commit/d534a7d6afcd4f0da3049d5ccd2b4eef3c45bf35))
* gradle.properties path ([c963b43](https://github.com/lgdd/liferay-upgrade-action/commit/c963b430e64855329a74732aa7ee148c3309585b))
* remove working-directory for setup-go ([06ed407](https://github.com/lgdd/liferay-upgrade-action/commit/06ed407191753cc7590c95df9016f9a839efb3c3))
* use working-directory to properly find workspace files ([5f3dcdb](https://github.com/lgdd/liferay-upgrade-action/commit/5f3dcdb56c68c29f431fee27c4a14b27ad59918f))
* wrong upgrade branch name reference in yaml ([5d7bd63](https://github.com/lgdd/liferay-upgrade-action/commit/5d7bd63798e241fa8434a40230b27e3eac0eed5e))
* wrong workspace path for git add ([95d981c](https://github.com/lgdd/liferay-upgrade-action/commit/95d981c151468f82d9086195a5aac8f43a4bd303))

## [2.0.0](https://github.com/lgdd/liferay-upgrade-action/compare/v1.0.2...v2.0.0) (2023-06-06)


### ⚠ BREAKING CHANGES

* Checkout step included in this action has been removed. You need to add it first in your own workflow before using this action.

### Miscellaneous Chores

* remove the checkout step ([c1f1b0f](https://github.com/lgdd/liferay-upgrade-action/commit/c1f1b0f0c6dc6c960dffc2a730d2885991aebc89))

## [1.0.2](https://github.com/lgdd/liferay-upgrade-action/compare/v1.0.1...v1.0.2) (2023-06-05)


### Bug Fixes

* fix output syntax ([47f9d01](https://github.com/lgdd/liferay-upgrade-action/commit/47f9d0134488467bbffb8f3e1cf22ddc4a1eea66))
* incorrect output keys ([cf9657f](https://github.com/lgdd/liferay-upgrade-action/commit/cf9657f828478c62618b365198eb50906eac34e9))
* missuse of run (script instead) ([0877cfb](https://github.com/lgdd/liferay-upgrade-action/commit/0877cfbf5bad31f4905d3b6edd650f845aef9c82))
* pass input to get-liferay-info ([b26ce31](https://github.com/lgdd/liferay-upgrade-action/commit/b26ce31dd0fafa4857377a385eb4902423180b4f))
* step outputs pass as env vars ([037ced1](https://github.com/lgdd/liferay-upgrade-action/commit/037ced1ac2560ea8e09d7c8f261806b6863880c9))

## [1.0.1](https://github.com/lgdd/liferay-upgrade-action/compare/v1.0.0...v1.0.1) (2023-06-04)


### Bug Fixes

* check branch step shell syntax ([51e447c](https://github.com/lgdd/liferay-upgrade-action/commit/51e447c46887d0981854b0b401762866f1f9e26f))

## 1.0.0 (2023-06-04)


### Features

* adapt pull request message to highlight build success or failure ([b952ea0](https://github.com/lgdd/liferay-upgrade-action/commit/b952ea09de6de3a1ff08d50c8f1def8059dde074))
* add checkout input ([bf4fe1b](https://github.com/lgdd/liferay-upgrade-action/commit/bf4fe1ba39cb60b9576c4c29e1e9d242171686c1))
* add step to remove branch if pull request fails ([06dd133](https://github.com/lgdd/liferay-upgrade-action/commit/06dd133ce934487270431ff30f38804382305c3e))
* add upgrade action ([9cc8310](https://github.com/lgdd/liferay-upgrade-action/commit/9cc8310e4326d049803b746ce2f157fcca1874a2))


### Bug Fixes

* **ci:** add missing checkout step ([be7333e](https://github.com/lgdd/liferay-upgrade-action/commit/be7333e3fd7bdd8636a2eab068324986b5bb7e3c))
* delete origin only ([c7892f6](https://github.com/lgdd/liferay-upgrade-action/commit/c7892f6c9c73f71994fd21e1fe818ccbea18a39a))
* **docs:** wording ([36eae60](https://github.com/lgdd/liferay-upgrade-action/commit/36eae6071193e8bd462fe71479ae7bd33c611162))
* inputs usage instead of env vars ([458ac7e](https://github.com/lgdd/liferay-upgrade-action/commit/458ac7e2e27b6485dab776e99f4c6938c6e07aab))
* make checkout fetch all tags and branches ([6b6ca1b](https://github.com/lgdd/liferay-upgrade-action/commit/6b6ca1be9d18ad2bed3d3a8e05fab2e6f3e2a814))
* only delete branch with create branch succeeded and create pr failed ([01cf9f6](https://github.com/lgdd/liferay-upgrade-action/commit/01cf9f6c9001f9634b87b9e878089aefa60e61e0))
* token for github cli ([0a39c77](https://github.com/lgdd/liferay-upgrade-action/commit/0a39c77b05b49611c545f24821082b79740d301b))
