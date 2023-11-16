# Changelog

## [0.0.20](https://github.com/spectrocloud-labs/validator/compare/v0.0.19...v0.0.20) (2023-11-16)


### Bug Fixes

* retry all status updates due to controller contention ([#114](https://github.com/spectrocloud-labs/validator/issues/114)) ([35f03a4](https://github.com/spectrocloud-labs/validator/commit/35f03a407a3d0bbcfd76c749908e4b1c9581afac))

## [0.0.19](https://github.com/spectrocloud-labs/validator/compare/v0.0.18...v0.0.19) (2023-11-16)


### Features

* add alertmanager sink ([#107](https://github.com/spectrocloud-labs/validator/issues/107)) ([855e70e](https://github.com/spectrocloud-labs/validator/commit/855e70e69c67cd338f83add9b0b18026e3395184))


### Bug Fixes

* **deps:** update kubernetes packages to v0.28.4 ([#112](https://github.com/spectrocloud-labs/validator/issues/112)) ([fc10444](https://github.com/spectrocloud-labs/validator/commit/fc104445fab89a663ff0e3fee8ea500b1d0a0904))
* ensure plugin removal during Helm uninstall ([#111](https://github.com/spectrocloud-labs/validator/issues/111)) ([0917418](https://github.com/spectrocloud-labs/validator/commit/0917418b6ae3f2940bf8048c0cb09ca4056f21da))


### Docs

* issue template addition ([#109](https://github.com/spectrocloud-labs/validator/issues/109)) ([36ce4a1](https://github.com/spectrocloud-labs/validator/commit/36ce4a1d5630c22b39d481bc45641c5c06e6db04))


### Refactoring

* accept VR in HandleNewValidationResult for flexibility in plugins ([#113](https://github.com/spectrocloud-labs/validator/issues/113)) ([1ead151](https://github.com/spectrocloud-labs/validator/commit/1ead15146156ac278aedb2a77cab0604488fda4f))

## [0.0.18](https://github.com/spectrocloud-labs/validator/compare/v0.0.17...v0.0.18) (2023-11-12)


### Bug Fixes

* **deps:** update module github.com/onsi/ginkgo/v2 to v2.13.1 ([#95](https://github.com/spectrocloud-labs/validator/issues/95)) ([496ecad](https://github.com/spectrocloud-labs/validator/commit/496ecada5655f5760e46f7d647ce381f616ad56f))
* **deps:** update module sigs.k8s.io/yaml to v1.4.0 ([#98](https://github.com/spectrocloud-labs/validator/issues/98)) ([5f35bba](https://github.com/spectrocloud-labs/validator/commit/5f35bbac77502a944d6d5641e1e2f88f98cf7c79))
* retry VR status updates ([21b3808](https://github.com/spectrocloud-labs/validator/commit/21b3808f36a621f89ddc22aa5362d4d7b47265b5))
* SafeUpdateValidationResult not handling all edge cases ([#104](https://github.com/spectrocloud-labs/validator/issues/104)) ([8f34e2f](https://github.com/spectrocloud-labs/validator/commit/8f34e2f677a2b70c3c931491ce8b5cd6ac7abd0b))


### Other

* **deps:** pin codecov/codecov-action action to eaaf4be ([#105](https://github.com/spectrocloud-labs/validator/issues/105)) ([70c3a0d](https://github.com/spectrocloud-labs/validator/commit/70c3a0d834cccc0bc6400c12395b7952c3fb25bb))
* **deps:** pin ncruces/go-coverage-report action to a12281c ([#101](https://github.com/spectrocloud-labs/validator/issues/101)) ([bb5d6ac](https://github.com/spectrocloud-labs/validator/commit/bb5d6ac4b235013eeba198e2ba4fd86c5eeb93a2))

## [0.0.17](https://github.com/spectrocloud-labs/validator/compare/v0.0.16...v0.0.17) (2023-11-10)


### Bug Fixes

* ensure ValidationResult Status resets to successful if all checks pass ([10237c5](https://github.com/spectrocloud-labs/validator/commit/10237c5a17adbc61288c6c6b9b981b072ea0e46f))


### Other

* add coverage report ([#100](https://github.com/spectrocloud-labs/validator/issues/100)) ([5d24704](https://github.com/spectrocloud-labs/validator/commit/5d2470488e1a09ca0d0f623762ae926dea5ba3ef))

## [0.0.16](https://github.com/spectrocloud-labs/validator/compare/v0.0.15...v0.0.16) (2023-11-10)


### Bug Fixes

* **deps:** update golang.org/x/exp digest to 9a3e603 ([#97](https://github.com/spectrocloud-labs/validator/issues/97)) ([17c83e2](https://github.com/spectrocloud-labs/validator/commit/17c83e2f9bc4580eb8b2d9b1f9b4767a42a32d2f))
* ensure State always remains failed if any conditions fail ([4f55823](https://github.com/spectrocloud-labs/validator/commit/4f558234af6c190186f8d82fb4a6e135d83030e8))
* ensure State always remains failed if any conditions fail ([711485f](https://github.com/spectrocloud-labs/validator/commit/711485f130804eab34de871f86c138adcb3bf257))


### Other

* fix Helm chart lint error ([c938c89](https://github.com/spectrocloud-labs/validator/commit/c938c89947371fa3e3e3e2d43044749a413ec443))
* release 0.0.16 ([f031904](https://github.com/spectrocloud-labs/validator/commit/f031904a18fb5854586d58145842689d12028360))

## [0.0.15](https://github.com/spectrocloud-labs/validator/compare/v0.0.14...v0.0.15) (2023-11-10)


### Bug Fixes

* ensure State always remains failed if any conditions fail ([67e8462](https://github.com/spectrocloud-labs/validator/commit/67e846293ce26e8f416fbb24d4b247d38f2a15aa))


### Other

* Merge branch 'main' of https://github.com/spectrocloud-labs/validator ([42cf0ef](https://github.com/spectrocloud-labs/validator/commit/42cf0efc8ee8f8187b260ac0f4d1f8b1b9db6c16))
* release 0.0.15 ([d574854](https://github.com/spectrocloud-labs/validator/commit/d57485496dc90f4b6e421d978bdf0451edfcc59c))

## [0.0.14](https://github.com/spectrocloud-labs/validator/compare/v0.0.13...v0.0.14) (2023-11-10)


### Bug Fixes

* prevent extraneous sink emissions ([22de119](https://github.com/spectrocloud-labs/validator/commit/22de119a3ce93c8fb751473f5250d55446012d83))


### Other

* Merge branch 'main' of https://github.com/spectrocloud-labs/validator ([a2700a2](https://github.com/spectrocloud-labs/validator/commit/a2700a26661d6c7668965cc5b5a377adfd49d109))


### Refactoring

* use header block for msg titles ([a6ce7af](https://github.com/spectrocloud-labs/validator/commit/a6ce7af748dcb6cbd50c5d13e1a2c9a724e9a328))

## [0.0.13](https://github.com/spectrocloud-labs/validator/compare/v0.0.11...v0.0.13) (2023-11-10)


### Bug Fixes

* handle VRs w/ multiple conditions ([8a3a243](https://github.com/spectrocloud-labs/validator/commit/8a3a2431dbd317dce80376c64d336317c688c2a3))
* omit sink state from hash to avoid double-emitting on VR creation ([071b960](https://github.com/spectrocloud-labs/validator/commit/071b9602fee35262a66a9f403ceb878ac4a422b4))


### Other

* release 0.0.12 ([4903786](https://github.com/spectrocloud-labs/validator/commit/49037866402a7f16ef2c15cf172a11626392b9ff))
* release 0.0.13 ([1233488](https://github.com/spectrocloud-labs/validator/commit/1233488756f02cd3de9ee2a0d800cb29522545c8))
* Update default values.yaml ([f9af283](https://github.com/spectrocloud-labs/validator/commit/f9af2832bd652c9ea404d08936a6879930d29753))


### Refactoring

* change devspace port; always use exponential backoff; tidy validation result helpers ([c866429](https://github.com/spectrocloud-labs/validator/commit/c866429772e8d301916eed7ea8908b811cad3e7c))

## [0.0.11](https://github.com/spectrocloud-labs/validator/compare/v0.0.10...v0.0.11) (2023-11-09)


### Features

* add sink interface, Slack sink ([#84](https://github.com/spectrocloud-labs/validator/issues/84)) ([dac2c3a](https://github.com/spectrocloud-labs/validator/commit/dac2c3a83eebbe433790aa439cab1433eadfb0ec))


### Bug Fixes

* **deps:** update golang.org/x/exp digest to 2478ac8 ([#83](https://github.com/spectrocloud-labs/validator/issues/83)) ([0d5200f](https://github.com/spectrocloud-labs/validator/commit/0d5200faf789343c10149d5f1333894a51d13ff8))
* **deps:** update module github.com/go-logr/logr to v1.3.0 ([#77](https://github.com/spectrocloud-labs/validator/issues/77)) ([532fd6f](https://github.com/spectrocloud-labs/validator/commit/532fd6f82cf6a9f8322a74428e0834bb023ff67e))
* **deps:** update module github.com/onsi/gomega to v1.28.1 ([#74](https://github.com/spectrocloud-labs/validator/issues/74)) ([eb7d990](https://github.com/spectrocloud-labs/validator/commit/eb7d990a840d14c5700dffc549de8f27feb0b614))
* **deps:** update module github.com/onsi/gomega to v1.29.0 ([#76](https://github.com/spectrocloud-labs/validator/issues/76)) ([ca58e8c](https://github.com/spectrocloud-labs/validator/commit/ca58e8c622c75ce345550e2115f8311428bfceb5))
* **deps:** update module github.com/onsi/gomega to v1.30.0 ([#82](https://github.com/spectrocloud-labs/validator/issues/82)) ([7bfa8f7](https://github.com/spectrocloud-labs/validator/commit/7bfa8f71aa77db2953aaba698f6acf4f2700a03e))
* **deps:** update module k8s.io/klog/v2 to v2.110.1 ([#78](https://github.com/spectrocloud-labs/validator/issues/78)) ([8a79219](https://github.com/spectrocloud-labs/validator/commit/8a79219a40a4beb47182da4f5aea0d8045dd119f))
* update plugin versions in default values.yaml ([6f7f90d](https://github.com/spectrocloud-labs/validator/commit/6f7f90d15fbeb105df6c5b6c568c35fa4f12121f))


### Other

* add license ([065ef97](https://github.com/spectrocloud-labs/validator/commit/065ef97d16dadf35a54b84fe3cd1026e14f172d5))
* **deps:** update gcr.io/kubebuilder/kube-rbac-proxy docker tag to v0.15.0 ([#72](https://github.com/spectrocloud-labs/validator/issues/72)) ([4f0fc63](https://github.com/spectrocloud-labs/validator/commit/4f0fc630bc3ab969d6411fb1d31d968e313a20be))
* **deps:** update google-github-actions/release-please-action digest to db8f2c6 ([#81](https://github.com/spectrocloud-labs/validator/issues/81)) ([20956a3](https://github.com/spectrocloud-labs/validator/commit/20956a3fa864d5cc0e37349ba0632da61879b8b8))
* **deps:** update helm/chart-testing-action action to v2.6.0 ([#79](https://github.com/spectrocloud-labs/validator/issues/79)) ([3272b22](https://github.com/spectrocloud-labs/validator/commit/3272b226df2586344efd1dcf3f077483ca3f45a7))
* **deps:** update helm/chart-testing-action action to v2.6.1 ([#80](https://github.com/spectrocloud-labs/validator/issues/80)) ([cbb430e](https://github.com/spectrocloud-labs/validator/commit/cbb430e70a8aad6558816d3ce2c5c42cacefcab8))
* **main:** release 0.0.10 ([#70](https://github.com/spectrocloud-labs/validator/issues/70)) ([6c5e4fc](https://github.com/spectrocloud-labs/validator/commit/6c5e4fcc1182862e3902044e543309800e96b31e))
* **main:** release 0.0.10 ([#75](https://github.com/spectrocloud-labs/validator/issues/75)) ([32e4025](https://github.com/spectrocloud-labs/validator/commit/32e4025ba21223c7214e5378bb060769b931e685))
* release 0.0.10 ([65ce207](https://github.com/spectrocloud-labs/validator/commit/65ce2076727fd3d1da49afd884329c55a6394c91))
* release 0.0.11 ([f694577](https://github.com/spectrocloud-labs/validator/commit/f694577bb2b0fae8127935c3f9fd9e90f14fe328))

## [0.0.10](https://github.com/spectrocloud-labs/validator/compare/v0.0.10...v0.0.10) (2023-11-09)


### Features

* add sink interface, Slack sink ([#84](https://github.com/spectrocloud-labs/validator/issues/84)) ([dac2c3a](https://github.com/spectrocloud-labs/validator/commit/dac2c3a83eebbe433790aa439cab1433eadfb0ec))


### Bug Fixes

* **deps:** update golang.org/x/exp digest to 2478ac8 ([#83](https://github.com/spectrocloud-labs/validator/issues/83)) ([0d5200f](https://github.com/spectrocloud-labs/validator/commit/0d5200faf789343c10149d5f1333894a51d13ff8))
* **deps:** update module github.com/go-logr/logr to v1.3.0 ([#77](https://github.com/spectrocloud-labs/validator/issues/77)) ([532fd6f](https://github.com/spectrocloud-labs/validator/commit/532fd6f82cf6a9f8322a74428e0834bb023ff67e))
* **deps:** update module github.com/onsi/gomega to v1.28.1 ([#74](https://github.com/spectrocloud-labs/validator/issues/74)) ([eb7d990](https://github.com/spectrocloud-labs/validator/commit/eb7d990a840d14c5700dffc549de8f27feb0b614))
* **deps:** update module github.com/onsi/gomega to v1.29.0 ([#76](https://github.com/spectrocloud-labs/validator/issues/76)) ([ca58e8c](https://github.com/spectrocloud-labs/validator/commit/ca58e8c622c75ce345550e2115f8311428bfceb5))
* **deps:** update module github.com/onsi/gomega to v1.30.0 ([#82](https://github.com/spectrocloud-labs/validator/issues/82)) ([7bfa8f7](https://github.com/spectrocloud-labs/validator/commit/7bfa8f71aa77db2953aaba698f6acf4f2700a03e))
* **deps:** update module k8s.io/klog/v2 to v2.110.1 ([#78](https://github.com/spectrocloud-labs/validator/issues/78)) ([8a79219](https://github.com/spectrocloud-labs/validator/commit/8a79219a40a4beb47182da4f5aea0d8045dd119f))
* update plugin versions in default values.yaml ([6f7f90d](https://github.com/spectrocloud-labs/validator/commit/6f7f90d15fbeb105df6c5b6c568c35fa4f12121f))


### Other

* add license ([065ef97](https://github.com/spectrocloud-labs/validator/commit/065ef97d16dadf35a54b84fe3cd1026e14f172d5))
* **deps:** update gcr.io/kubebuilder/kube-rbac-proxy docker tag to v0.15.0 ([#72](https://github.com/spectrocloud-labs/validator/issues/72)) ([4f0fc63](https://github.com/spectrocloud-labs/validator/commit/4f0fc630bc3ab969d6411fb1d31d968e313a20be))
* **deps:** update google-github-actions/release-please-action digest to db8f2c6 ([#81](https://github.com/spectrocloud-labs/validator/issues/81)) ([20956a3](https://github.com/spectrocloud-labs/validator/commit/20956a3fa864d5cc0e37349ba0632da61879b8b8))
* **deps:** update helm/chart-testing-action action to v2.6.0 ([#79](https://github.com/spectrocloud-labs/validator/issues/79)) ([3272b22](https://github.com/spectrocloud-labs/validator/commit/3272b226df2586344efd1dcf3f077483ca3f45a7))
* **deps:** update helm/chart-testing-action action to v2.6.1 ([#80](https://github.com/spectrocloud-labs/validator/issues/80)) ([cbb430e](https://github.com/spectrocloud-labs/validator/commit/cbb430e70a8aad6558816d3ce2c5c42cacefcab8))
* **main:** release 0.0.10 ([#70](https://github.com/spectrocloud-labs/validator/issues/70)) ([6c5e4fc](https://github.com/spectrocloud-labs/validator/commit/6c5e4fcc1182862e3902044e543309800e96b31e))
* release 0.0.10 ([65ce207](https://github.com/spectrocloud-labs/validator/commit/65ce2076727fd3d1da49afd884329c55a6394c91))

## [0.0.10](https://github.com/spectrocloud-labs/validator/compare/v0.0.10...v0.0.10) (2023-10-20)


### Bug Fixes

* update plugin versions in default values.yaml ([6f7f90d](https://github.com/spectrocloud-labs/validator/commit/6f7f90d15fbeb105df6c5b6c568c35fa4f12121f))


### Other

* release 0.0.10 ([65ce207](https://github.com/spectrocloud-labs/validator/commit/65ce2076727fd3d1da49afd884329c55a6394c91))

## [0.0.10](https://github.com/spectrocloud-labs/validator/compare/v0.0.9...v0.0.10) (2023-10-20)


### Bug Fixes

* **deps:** update module sigs.k8s.io/controller-runtime to v0.16.3 ([#63](https://github.com/spectrocloud-labs/validator/issues/63)) ([6f79f8a](https://github.com/spectrocloud-labs/validator/commit/6f79f8af8f7a10c70ce403fadeb16d6eb9e13020))


### Other

* **deps:** bump golang.org/x/net from 0.16.0 to 0.17.0 ([#61](https://github.com/spectrocloud-labs/validator/issues/61)) ([eca7abd](https://github.com/spectrocloud-labs/validator/commit/eca7abd8da204cabb28d7fb6ee8c381d1cb60e7b))
* **deps:** update actions/checkout digest to b4ffde6 ([#64](https://github.com/spectrocloud-labs/validator/issues/64)) ([d9bbd21](https://github.com/spectrocloud-labs/validator/commit/d9bbd21fe962de4f14e0e734d697ebd2ceb7e144))
* **deps:** update actions/setup-python digest to 65d7f2d ([#65](https://github.com/spectrocloud-labs/validator/issues/65)) ([bdb95d0](https://github.com/spectrocloud-labs/validator/commit/bdb95d038149ed9eab6c8372018bb771b848157c))
* **deps:** update gcr.io/kubebuilder/kube-rbac-proxy docker tag to v0.14.4 ([#59](https://github.com/spectrocloud-labs/validator/issues/59)) ([78847f5](https://github.com/spectrocloud-labs/validator/commit/78847f54c35f6fc684d333a286e1315cc558e6e2))
* **deps:** update google-github-actions/release-please-action digest to 4c5670f ([#66](https://github.com/spectrocloud-labs/validator/issues/66)) ([2c24d48](https://github.com/spectrocloud-labs/validator/commit/2c24d48035b7ebddbbd20ca182e3352fa6c2f09e))
* enable renovate automerges ([84ad3cd](https://github.com/spectrocloud-labs/validator/commit/84ad3cdee59ed634e5f9577d801dc13701958e46))
* release 0.0.10 ([9a324e9](https://github.com/spectrocloud-labs/validator/commit/9a324e9e666b0da68a1e0c3be225ab19bfd04a6b))


### Refactoring

* valid8or -&gt; validator ([#67](https://github.com/spectrocloud-labs/validator/issues/67)) ([ff88026](https://github.com/spectrocloud-labs/validator/commit/ff8802656d8115dd6afbbfbaf12613c4f205feb5))

## [0.0.9](https://github.com/spectrocloud-labs/validator/compare/v0.0.8...v0.0.9) (2023-10-10)


### Bug Fixes

* **deps:** update golang.org/x/exp digest to 7918f67 ([#55](https://github.com/spectrocloud-labs/validator/issues/55)) ([3f173d4](https://github.com/spectrocloud-labs/validator/commit/3f173d4dc256415d9f447133afc70024d0115021))
* **deps:** update kubernetes packages to v0.28.2 ([#51](https://github.com/spectrocloud-labs/validator/issues/51)) ([f43d5a0](https://github.com/spectrocloud-labs/validator/commit/f43d5a098d6d5923fa540564279defe31701f3c7))
* **deps:** update module github.com/onsi/ginkgo/v2 to v2.12.1 ([#52](https://github.com/spectrocloud-labs/validator/issues/52)) ([516693e](https://github.com/spectrocloud-labs/validator/commit/516693e3d150dfceb0a6be3a5f00bfbe260a1cb6))
* **deps:** update module github.com/onsi/ginkgo/v2 to v2.13.0 ([#57](https://github.com/spectrocloud-labs/validator/issues/57)) ([be32cb8](https://github.com/spectrocloud-labs/validator/commit/be32cb85ea38e8ddfd03e9d6837dddda1967c4b7))
* **deps:** update module github.com/onsi/gomega to v1.28.0 ([#54](https://github.com/spectrocloud-labs/validator/issues/54)) ([e89431e](https://github.com/spectrocloud-labs/validator/commit/e89431e5685fc9ab83dd7dd2a87864f57b835bcb))
* **deps:** update module sigs.k8s.io/controller-runtime to v0.16.2 ([#50](https://github.com/spectrocloud-labs/validator/issues/50)) ([1a07c2a](https://github.com/spectrocloud-labs/validator/commit/1a07c2a134bc26d86a17c04fcf71bae2f601c3e5))


### Other

* **deps:** update actions/checkout digest to 8ade135 ([#53](https://github.com/spectrocloud-labs/validator/issues/53)) ([d42a8a6](https://github.com/spectrocloud-labs/validator/commit/d42a8a6832c15b76eac882a7f35fc3a2f8ba842a))
* **deps:** update docker/build-push-action action to v5 ([#47](https://github.com/spectrocloud-labs/validator/issues/47)) ([40d0053](https://github.com/spectrocloud-labs/validator/commit/40d0053011c9d05582d96ab07c7e743867fbed8d))
* **deps:** update docker/build-push-action digest to 0a97817 ([#45](https://github.com/spectrocloud-labs/validator/issues/45)) ([91a951b](https://github.com/spectrocloud-labs/validator/commit/91a951b03d40db29020af68cb7b1a22db8cad19c))
* **deps:** update docker/login-action action to v3 ([#48](https://github.com/spectrocloud-labs/validator/issues/48)) ([9c940d7](https://github.com/spectrocloud-labs/validator/commit/9c940d7acebf0e237ed5129cc3b25aca596a257f))
* **deps:** update docker/setup-buildx-action action to v3 ([#49](https://github.com/spectrocloud-labs/validator/issues/49)) ([a3fe730](https://github.com/spectrocloud-labs/validator/commit/a3fe730ae38b9655583be333e255e48767c2cf0c))
* release 0.0.9 ([3cbcfa9](https://github.com/spectrocloud-labs/validator/commit/3cbcfa9e3f1e5c75dcfc083df1494aab0c41c31f))

## [0.0.8](https://github.com/spectrocloud-labs/validator/compare/v0.0.7...v0.0.8) (2023-09-06)


### Other

* bump aws plugin version ([bac567e](https://github.com/spectrocloud-labs/validator/commit/bac567e69ac6a957b78f29b5bcec81d2f3b0c6b8))
* **deps:** update actions/upload-artifact digest to a8a3f3a ([#43](https://github.com/spectrocloud-labs/validator/issues/43)) ([fc33646](https://github.com/spectrocloud-labs/validator/commit/fc3364674e45c9dd3fdbf15773a7fbe2f04e3395))

## [0.0.7](https://github.com/spectrocloud-labs/validator/compare/v0.0.6...v0.0.7) (2023-09-06)


### Bug Fixes

* **deps:** update golang.org/x/exp digest to 9212866 ([#41](https://github.com/spectrocloud-labs/validator/issues/41)) ([50ad9cb](https://github.com/spectrocloud-labs/validator/commit/50ad9cbd72f531ab3e29eb43d59f75026efc96b0))
* include pkg/ in dockerfiles ([b45bb22](https://github.com/spectrocloud-labs/validator/commit/b45bb22532adc3e860aea56edf7ab3f3a95519fd))


### Other

* bump plugin versions ([6658d19](https://github.com/spectrocloud-labs/validator/commit/6658d190022815cf56d1a345ae66d46fd260c10c))
* **deps:** update actions/checkout action to v4 ([#39](https://github.com/spectrocloud-labs/validator/issues/39)) ([3c848b6](https://github.com/spectrocloud-labs/validator/commit/3c848b61c3294cdd1ceed376c3e4a48521221b6f))
* update AWS plugin version; fix default values.yaml ([516213b](https://github.com/spectrocloud-labs/validator/commit/516213ba35d80c4e8692b1448b22f6e1b9756c5d))


### Refactoring

* expose helm package ([2a34e0a](https://github.com/spectrocloud-labs/validator/commit/2a34e0ae780b287a8a1a48ee00f5016e667c8304))

## [0.0.6](https://github.com/spectrocloud-labs/validator/compare/v0.0.5...v0.0.6) (2023-09-01)


### Bug Fixes

* **deps:** update kubernetes packages to v0.28.1 ([#30](https://github.com/spectrocloud-labs/validator/issues/30)) ([f94b40d](https://github.com/spectrocloud-labs/validator/commit/f94b40d7d9b0be097cad185b0426727380cef822))
* **deps:** update module github.com/onsi/ginkgo/v2 to v2.12.0 ([#31](https://github.com/spectrocloud-labs/validator/issues/31)) ([98a7aa7](https://github.com/spectrocloud-labs/validator/commit/98a7aa785946534db076a093c0715ac63782d72f))
* **deps:** update module github.com/onsi/gomega to v1.27.10 ([#29](https://github.com/spectrocloud-labs/validator/issues/29)) ([8697124](https://github.com/spectrocloud-labs/validator/commit/8697124f495004f71ed9ba694ab8116880a4ae7f))
* **deps:** update module sigs.k8s.io/controller-runtime to v0.16.1 ([#33](https://github.com/spectrocloud-labs/validator/issues/33)) ([94bf0ad](https://github.com/spectrocloud-labs/validator/commit/94bf0ad33b8524a82427b12320ba70493af9ac21))
* NET_ADMIN -&gt; NET_RAW ([122cc80](https://github.com/spectrocloud-labs/validator/commit/122cc808ed6b83eb6a33dc38e031d86805440784))


### Other

* **deps:** update actions/checkout digest to f43a0e5 ([#25](https://github.com/spectrocloud-labs/validator/issues/25)) ([fa0b3d9](https://github.com/spectrocloud-labs/validator/commit/fa0b3d95f74e70226f07114bffcde2e1b270ad33))
* **deps:** update actions/setup-go digest to 93397be ([#26](https://github.com/spectrocloud-labs/validator/issues/26)) ([e32d52b](https://github.com/spectrocloud-labs/validator/commit/e32d52bc1243451b6a0f6a27f228869da9497761))
* **deps:** update docker/setup-buildx-action digest to 885d146 ([#28](https://github.com/spectrocloud-labs/validator/issues/28)) ([f4b1dd1](https://github.com/spectrocloud-labs/validator/commit/f4b1dd10a4170f2596bf2cebf64a363d21efcc44))
* release 0.0.6 ([10421d5](https://github.com/spectrocloud-labs/validator/commit/10421d59a5fbc4dd840a6aac9f24657ee8d1be7c))

## [0.0.5](https://github.com/spectrocloud-labs/validator/compare/v0.0.4...v0.0.5) (2023-08-31)


### Bug Fixes

* omit conditions for uninstalled plugins ([c9f430d](https://github.com/spectrocloud-labs/validator/commit/c9f430d81bfdbb077edfa5d3cc48f314bf831c45))
* preserve VC annotations when updating plugin hashes ([19c9463](https://github.com/spectrocloud-labs/validator/commit/19c9463a4ed7b516731fdaf76cf487e682a6a2c4))
* securityContext blocking MTU check w/ ping ([131e5d9](https://github.com/spectrocloud-labs/validator/commit/131e5d91015b54b470b61708bfe8675f7eb26a0e))
* update 2+ plugin conditions properly ([a12488f](https://github.com/spectrocloud-labs/validator/commit/a12488f9dc376d1b0a8b413791b3aa2e25b185a5))


### Other

* release 0.0.5 ([24e9712](https://github.com/spectrocloud-labs/validator/commit/24e9712ffbf6fb8d33333d5d8c063f3935c7ceae))
* update README ([f8254d4](https://github.com/spectrocloud-labs/validator/commit/f8254d4a0bca6db90523ace43503cf6d80d3af30))


### Docs

* add validator-plugin-network to default values.yaml ([1aada24](https://github.com/spectrocloud-labs/validator/commit/1aada2492d518e7ed50934d2cb6e184da5d1e031))

## [0.0.4](https://github.com/spectrocloud-labs/validator/compare/v0.0.3...v0.0.4) (2023-08-29)


### Other

* add public validation result utils ([528be5f](https://github.com/spectrocloud-labs/validator/commit/528be5f91e8bfb7d6f1530002dd99971e4983a7e))
* release 0.0.4 ([28f8418](https://github.com/spectrocloud-labs/validator/commit/28f8418356694ad5c86b97d2b9df9d51c2f6d279))


### Docs

* update chart description ([0a5635f](https://github.com/spectrocloud-labs/validator/commit/0a5635f94f949c39592dbe7f20e5301f4836f291))

## [0.0.3](https://github.com/spectrocloud-labs/validator/compare/v0.0.2...v0.0.3) (2023-08-29)


### Features

* handle plugin updates via values hashes ([7f485b4](https://github.com/spectrocloud-labs/validator/commit/7f485b41f5dfba40e8c08a5da79410dfc0c97e0c))
* log ValidationResult metadata on completion ([0cc38e5](https://github.com/spectrocloud-labs/validator/commit/0cc38e5cf464d6f9342865f3a41787dfe9bc3c5c))
* plugin management w/ helm ([537faac](https://github.com/spectrocloud-labs/validator/commit/537faac4c3f1c6695f1db34114401a14ad292906))
* update status and handle plugin removal ([bae7e9d](https://github.com/spectrocloud-labs/validator/commit/bae7e9dc36a1a22e8f08828421d0cc7e73deb54f))


### Bug Fixes

* increase memory limit for helm upgrade ([660a80d](https://github.com/spectrocloud-labs/validator/commit/660a80d57fcae2dc3a16e610699b60a5749e0786))
* update RBAC in helm templates ([6ff735c](https://github.com/spectrocloud-labs/validator/commit/6ff735c891e40328bba0524f4b8b240d3e85a6c9))


### Other

* add pull_request test trigger ([3e75bdb](https://github.com/spectrocloud-labs/validator/commit/3e75bdbff56bab925ca98b81c57fe9f4b1a60471))
* release 0.0.3 ([5b2473d](https://github.com/spectrocloud-labs/validator/commit/5b2473dce974a160b14640e86e88028f66c94f5e))


### Docs

* update README, fix release-please annotations ([c2c96e8](https://github.com/spectrocloud-labs/validator/commit/c2c96e8e3e91820826242b36d6760ab1d2530baf))

## [0.0.2](https://github.com/spectrocloud-labs/validator/compare/v0.0.1...v0.0.2) (2023-08-25)


### Other

* release 0.0.2 ([40cdd88](https://github.com/spectrocloud-labs/validator/commit/40cdd88ebb8b75f9908c5dab6aa29337f5d778d8))

## [0.0.1](https://github.com/spectrocloud-labs/validator/compare/v0.0.1...v0.0.1) (2023-08-25)


### Bug Fixes

* helm chart CI ([46f37f0](https://github.com/spectrocloud-labs/validator/commit/46f37f0cea87e90e6effb85cb15128ab5970a621))
* release image push repo ([4a2aca6](https://github.com/spectrocloud-labs/validator/commit/4a2aca6ecbfeca48ed4dd7566441923815281432))


### Other

* release 0.0.1 ([a23551a](https://github.com/spectrocloud-labs/validator/commit/a23551a1984d43d9acbc7de3cacad6ee928cc517))
