# Changelog

## [0.0.6](https://github.com/spectrocloud-labs/valid8or/compare/v0.0.5...v0.0.6) (2023-09-01)


### Bug Fixes

* **deps:** update kubernetes packages to v0.28.1 ([#30](https://github.com/spectrocloud-labs/valid8or/issues/30)) ([f94b40d](https://github.com/spectrocloud-labs/valid8or/commit/f94b40d7d9b0be097cad185b0426727380cef822))
* **deps:** update module github.com/onsi/ginkgo/v2 to v2.12.0 ([#31](https://github.com/spectrocloud-labs/valid8or/issues/31)) ([98a7aa7](https://github.com/spectrocloud-labs/valid8or/commit/98a7aa785946534db076a093c0715ac63782d72f))
* **deps:** update module github.com/onsi/gomega to v1.27.10 ([#29](https://github.com/spectrocloud-labs/valid8or/issues/29)) ([8697124](https://github.com/spectrocloud-labs/valid8or/commit/8697124f495004f71ed9ba694ab8116880a4ae7f))
* **deps:** update module sigs.k8s.io/controller-runtime to v0.16.1 ([#33](https://github.com/spectrocloud-labs/valid8or/issues/33)) ([94bf0ad](https://github.com/spectrocloud-labs/valid8or/commit/94bf0ad33b8524a82427b12320ba70493af9ac21))
* NET_ADMIN -&gt; NET_RAW ([122cc80](https://github.com/spectrocloud-labs/valid8or/commit/122cc808ed6b83eb6a33dc38e031d86805440784))


### Other

* **deps:** update actions/checkout digest to f43a0e5 ([#25](https://github.com/spectrocloud-labs/valid8or/issues/25)) ([fa0b3d9](https://github.com/spectrocloud-labs/valid8or/commit/fa0b3d95f74e70226f07114bffcde2e1b270ad33))
* **deps:** update actions/setup-go digest to 93397be ([#26](https://github.com/spectrocloud-labs/valid8or/issues/26)) ([e32d52b](https://github.com/spectrocloud-labs/valid8or/commit/e32d52bc1243451b6a0f6a27f228869da9497761))
* **deps:** update docker/setup-buildx-action digest to 885d146 ([#28](https://github.com/spectrocloud-labs/valid8or/issues/28)) ([f4b1dd1](https://github.com/spectrocloud-labs/valid8or/commit/f4b1dd10a4170f2596bf2cebf64a363d21efcc44))

## [0.0.5](https://github.com/spectrocloud-labs/valid8or/compare/v0.0.4...v0.0.5) (2023-08-31)


### Bug Fixes

* omit conditions for uninstalled plugins ([c9f430d](https://github.com/spectrocloud-labs/valid8or/commit/c9f430d81bfdbb077edfa5d3cc48f314bf831c45))
* preserve VC annotations when updating plugin hashes ([19c9463](https://github.com/spectrocloud-labs/valid8or/commit/19c9463a4ed7b516731fdaf76cf487e682a6a2c4))
* securityContext blocking MTU check w/ ping ([131e5d9](https://github.com/spectrocloud-labs/valid8or/commit/131e5d91015b54b470b61708bfe8675f7eb26a0e))
* update 2+ plugin conditions properly ([a12488f](https://github.com/spectrocloud-labs/valid8or/commit/a12488f9dc376d1b0a8b413791b3aa2e25b185a5))


### Other

* release 0.0.5 ([24e9712](https://github.com/spectrocloud-labs/valid8or/commit/24e9712ffbf6fb8d33333d5d8c063f3935c7ceae))
* update README ([f8254d4](https://github.com/spectrocloud-labs/valid8or/commit/f8254d4a0bca6db90523ace43503cf6d80d3af30))


### Docs

* add valid8or-plugin-network to default values.yaml ([1aada24](https://github.com/spectrocloud-labs/valid8or/commit/1aada2492d518e7ed50934d2cb6e184da5d1e031))

## [0.0.4](https://github.com/spectrocloud-labs/valid8or/compare/v0.0.3...v0.0.4) (2023-08-29)


### Other

* add public validation result utils ([528be5f](https://github.com/spectrocloud-labs/valid8or/commit/528be5f91e8bfb7d6f1530002dd99971e4983a7e))
* release 0.0.4 ([28f8418](https://github.com/spectrocloud-labs/valid8or/commit/28f8418356694ad5c86b97d2b9df9d51c2f6d279))


### Docs

* update chart description ([0a5635f](https://github.com/spectrocloud-labs/valid8or/commit/0a5635f94f949c39592dbe7f20e5301f4836f291))

## [0.0.3](https://github.com/spectrocloud-labs/valid8or/compare/v0.0.2...v0.0.3) (2023-08-29)


### Features

* handle plugin updates via values hashes ([7f485b4](https://github.com/spectrocloud-labs/valid8or/commit/7f485b41f5dfba40e8c08a5da79410dfc0c97e0c))
* log ValidationResult metadata on completion ([0cc38e5](https://github.com/spectrocloud-labs/valid8or/commit/0cc38e5cf464d6f9342865f3a41787dfe9bc3c5c))
* plugin management w/ helm ([537faac](https://github.com/spectrocloud-labs/valid8or/commit/537faac4c3f1c6695f1db34114401a14ad292906))
* update status and handle plugin removal ([bae7e9d](https://github.com/spectrocloud-labs/valid8or/commit/bae7e9dc36a1a22e8f08828421d0cc7e73deb54f))


### Bug Fixes

* increase memory limit for helm upgrade ([660a80d](https://github.com/spectrocloud-labs/valid8or/commit/660a80d57fcae2dc3a16e610699b60a5749e0786))
* update RBAC in helm templates ([6ff735c](https://github.com/spectrocloud-labs/valid8or/commit/6ff735c891e40328bba0524f4b8b240d3e85a6c9))


### Other

* add pull_request test trigger ([3e75bdb](https://github.com/spectrocloud-labs/valid8or/commit/3e75bdbff56bab925ca98b81c57fe9f4b1a60471))
* release 0.0.3 ([5b2473d](https://github.com/spectrocloud-labs/valid8or/commit/5b2473dce974a160b14640e86e88028f66c94f5e))


### Docs

* update README, fix release-please annotations ([c2c96e8](https://github.com/spectrocloud-labs/valid8or/commit/c2c96e8e3e91820826242b36d6760ab1d2530baf))

## [0.0.2](https://github.com/spectrocloud-labs/valid8or/compare/v0.0.1...v0.0.2) (2023-08-25)


### Other

* release 0.0.2 ([40cdd88](https://github.com/spectrocloud-labs/valid8or/commit/40cdd88ebb8b75f9908c5dab6aa29337f5d778d8))

## [0.0.1](https://github.com/spectrocloud-labs/valid8or/compare/v0.0.1...v0.0.1) (2023-08-25)


### Bug Fixes

* helm chart CI ([46f37f0](https://github.com/spectrocloud-labs/valid8or/commit/46f37f0cea87e90e6effb85cb15128ab5970a621))
* release image push repo ([4a2aca6](https://github.com/spectrocloud-labs/valid8or/commit/4a2aca6ecbfeca48ed4dd7566441923815281432))


### Other

* release 0.0.1 ([a23551a](https://github.com/spectrocloud-labs/valid8or/commit/a23551a1984d43d9acbc7de3cacad6ee928cc517))
