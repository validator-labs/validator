# Changelog

## [0.0.40](https://github.com/validator-labs/validator/compare/v0.0.39...v0.0.40) (2024-05-17)


### Other

* bump plugin versions in helm chart ([#255](https://github.com/validator-labs/validator/issues/255)) ([1d053f0](https://github.com/validator-labs/validator/commit/1d053f0fd38eb405009291593b3532a85010ca00))
* remove old logo from README.md ([650afcd](https://github.com/validator-labs/validator/commit/650afcd26b46e93fa52b1cbb9ebb89f6fc464676))
* update chart docs ([#257](https://github.com/validator-labs/validator/issues/257)) ([16e3c1a](https://github.com/validator-labs/validator/commit/16e3c1a36faf4b5bd5891a0356796265997c7b3d))

## [0.0.39](https://github.com/validator-labs/validator/compare/v0.0.38...v0.0.39) (2024-05-17)


### Other

* bump validator-plugin-oci to v0.0.9 ([#250](https://github.com/validator-labs/validator/issues/250)) ([432e35f](https://github.com/validator-labs/validator/commit/432e35fb4128f27a5ef2d4b0a916d2849554c040))
* migrate from spectrocloud-labs to validator-labs ([#251](https://github.com/validator-labs/validator/issues/251)) ([0386d7e](https://github.com/validator-labs/validator/commit/0386d7e690910fb67476c3d9e5bea9133e11b4fc))
* release 0.0.39 ([1c594bd](https://github.com/validator-labs/validator/commit/1c594bda83cec0c749636a351b5279509519a012))
* update kubescape version ([#248](https://github.com/validator-labs/validator/issues/248)) ([7a0dcd4](https://github.com/validator-labs/validator/commit/7a0dcd41f8b626e463810694d3cd990235213700))
* update release.yaml ([44b4672](https://github.com/validator-labs/validator/commit/44b467285c7b8a67de7655765e3c11f7c20c3861))

## [0.0.38](https://github.com/validator-labs/validator/compare/v0.0.37...v0.0.38) (2024-03-14)


### Features

* support pubkey secret creation for oci signature verification ([#240](https://github.com/validator-labs/validator/issues/240)) ([0a77ed3](https://github.com/validator-labs/validator/commit/0a77ed34efa019afd926f26ee4a8a858123ff439))


### Other

* bump validator plugin versions ([#243](https://github.com/validator-labs/validator/issues/243)) ([2987021](https://github.com/validator-labs/validator/commit/2987021f4c59553965ab94159363cf3c1d6d23e1))

## [0.0.37](https://github.com/validator-labs/validator/compare/v0.0.36...v0.0.37) (2024-03-12)


### Features

* update multiple validation rule results at once ([#237](https://github.com/validator-labs/validator/issues/237)) ([5425ad0](https://github.com/validator-labs/validator/commit/5425ad0491be1e4765e3c90ccdce576f257bcdf7))


### Refactoring

* replace custom retries w/ patch helpers; make SinkState a condition ([#233](https://github.com/validator-labs/validator/issues/233)) ([38f1bc5](https://github.com/validator-labs/validator/commit/38f1bc58d1bcbf2cc83d7ee252ae56506859885d))

## [0.0.36](https://github.com/validator-labs/validator/compare/v0.0.35...v0.0.36) (2024-03-08)


### Features

* add alertmanager sink ([#107](https://github.com/validator-labs/validator/issues/107)) ([855e70e](https://github.com/validator-labs/validator/commit/855e70e69c67cd338f83add9b0b18026e3395184))
* add Azure plugin ([#131](https://github.com/validator-labs/validator/issues/131)) ([25073e3](https://github.com/validator-labs/validator/commit/25073e3c6fdc6b2556d3f0d59192b1b6ec65281b))
* add caFile Helm option ([#166](https://github.com/validator-labs/validator/issues/166)) ([2425599](https://github.com/validator-labs/validator/commit/24255997119b75edde4e8f94568a906c7da6b37f))
* add OCI plugin ([#139](https://github.com/validator-labs/validator/issues/139)) ([3ae7b70](https://github.com/validator-labs/validator/commit/3ae7b70920ed27f8ad1325fdbe0461c71ae48a71))
* add sink interface, Slack sink ([#84](https://github.com/validator-labs/validator/issues/84)) ([dac2c3a](https://github.com/validator-labs/validator/commit/dac2c3a83eebbe433790aa439cab1433eadfb0ec))
* expose insecureSkipVerify for Alertmanager sink ([#118](https://github.com/validator-labs/validator/issues/118)) ([8dc7548](https://github.com/validator-labs/validator/commit/8dc7548ed9511abbd382ec2881b473f0b42fd607))
* handle plugin updates via values hashes ([7f485b4](https://github.com/validator-labs/validator/commit/7f485b41f5dfba40e8c08a5da79410dfc0c97e0c))
* implement client to send FinalizeCleanup requests to spectro-cleanup ([#155](https://github.com/validator-labs/validator/issues/155)) ([2ae0348](https://github.com/validator-labs/validator/commit/2ae03480b48d7b1bffb20227889c7f1c2a12f60e))
* log ValidationResult metadata on completion ([0cc38e5](https://github.com/validator-labs/validator/commit/0cc38e5cf464d6f9342865f3a41787dfe9bc3c5c))
* OCI repository support for Helm charts ([#207](https://github.com/validator-labs/validator/issues/207)) ([4b25d79](https://github.com/validator-labs/validator/commit/4b25d7980c9aac07a39294a27ca3540e83dce1fd))
* plugin management w/ helm ([537faac](https://github.com/validator-labs/validator/commit/537faac4c3f1c6695f1db34114401a14ad292906))
* support private helm repos ([#132](https://github.com/validator-labs/validator/issues/132)) ([cb0cf32](https://github.com/validator-labs/validator/commit/cb0cf32e1c8b09bdaa791c1933f36321f1687046))
* support proxy configuration (env vars & CA certificate) ([#137](https://github.com/validator-labs/validator/issues/137)) ([63c3bc8](https://github.com/validator-labs/validator/commit/63c3bc8497766564d7e76a3da2dabc53ba3b7d54))
* update helm chart to enable fast cleanup in spectro-cleanup via gRPC endpoint ([#156](https://github.com/validator-labs/validator/issues/156)) ([6f91745](https://github.com/validator-labs/validator/commit/6f91745c9db6e92071e691a74942c0bb53692384))
* update status and handle plugin removal ([bae7e9d](https://github.com/validator-labs/validator/commit/bae7e9dc36a1a22e8f08828421d0cc7e73deb54f))


### Bug Fixes

* add yaml tags to ValidatorConfig types ([63afc70](https://github.com/validator-labs/validator/commit/63afc70819a4fffc6d90ba40cea12bca2577e743))
* delete plugins in parallel to avoid hitting timeouts ([#141](https://github.com/validator-labs/validator/issues/141)) ([a3fc0dc](https://github.com/validator-labs/validator/commit/a3fc0dc56c7c1964b2ee38b666a866c78711010e))
* **deps:** update golang.org/x/exp digest to 02704c9 ([#161](https://github.com/validator-labs/validator/issues/161)) ([72990b7](https://github.com/validator-labs/validator/commit/72990b7d01a917d0d2bbdedcfa41e019af993b34))
* **deps:** update golang.org/x/exp digest to 0dcbfd6 ([#175](https://github.com/validator-labs/validator/issues/175)) ([31469da](https://github.com/validator-labs/validator/commit/31469da899767e2d8a6766f609a9e280b6e49609))
* **deps:** update golang.org/x/exp digest to 1b97071 ([#187](https://github.com/validator-labs/validator/issues/187)) ([27581fa](https://github.com/validator-labs/validator/commit/27581fa0282aee29a032bfff15a29f23f9ff6390))
* **deps:** update golang.org/x/exp digest to 2478ac8 ([#83](https://github.com/validator-labs/validator/issues/83)) ([0d5200f](https://github.com/validator-labs/validator/commit/0d5200faf789343c10149d5f1333894a51d13ff8))
* **deps:** update golang.org/x/exp digest to 6522937 ([#133](https://github.com/validator-labs/validator/issues/133)) ([531c8ca](https://github.com/validator-labs/validator/commit/531c8ca2647164ac99777601c95ae469cb5c67c3))
* **deps:** update golang.org/x/exp digest to 73b9e39 ([#151](https://github.com/validator-labs/validator/issues/151)) ([bfaf2de](https://github.com/validator-labs/validator/commit/bfaf2de249f6edfb583ff5870be7572ddd2a55ee))
* **deps:** update golang.org/x/exp digest to 7918f67 ([#55](https://github.com/validator-labs/validator/issues/55)) ([3f173d4](https://github.com/validator-labs/validator/commit/3f173d4dc256415d9f447133afc70024d0115021))
* **deps:** update golang.org/x/exp digest to 9212866 ([#41](https://github.com/validator-labs/validator/issues/41)) ([50ad9cb](https://github.com/validator-labs/validator/commit/50ad9cbd72f531ab3e29eb43d59f75026efc96b0))
* **deps:** update golang.org/x/exp digest to 9a3e603 ([#97](https://github.com/validator-labs/validator/issues/97)) ([17c83e2](https://github.com/validator-labs/validator/commit/17c83e2f9bc4580eb8b2d9b1f9b4767a42a32d2f))
* **deps:** update golang.org/x/exp digest to be819d1 ([#171](https://github.com/validator-labs/validator/issues/171)) ([6d5d2a2](https://github.com/validator-labs/validator/commit/6d5d2a2c5fe210647db4d347b83b81fe7486e31b))
* **deps:** update golang.org/x/exp digest to db7319d ([#179](https://github.com/validator-labs/validator/issues/179)) ([04e3a3d](https://github.com/validator-labs/validator/commit/04e3a3d87324910cb8f32dc46865386e20f43378))
* **deps:** update kubernetes packages to v0.28.1 ([#30](https://github.com/validator-labs/validator/issues/30)) ([f94b40d](https://github.com/validator-labs/validator/commit/f94b40d7d9b0be097cad185b0426727380cef822))
* **deps:** update kubernetes packages to v0.28.2 ([#51](https://github.com/validator-labs/validator/issues/51)) ([f43d5a0](https://github.com/validator-labs/validator/commit/f43d5a098d6d5923fa540564279defe31701f3c7))
* **deps:** update kubernetes packages to v0.28.4 ([#112](https://github.com/validator-labs/validator/issues/112)) ([fc10444](https://github.com/validator-labs/validator/commit/fc104445fab89a663ff0e3fee8ea500b1d0a0904))
* **deps:** update kubernetes packages to v0.29.1 ([#153](https://github.com/validator-labs/validator/issues/153)) ([60bc244](https://github.com/validator-labs/validator/commit/60bc244537838596be433c3e2a59ce66fcb80350))
* **deps:** update kubernetes packages to v0.29.2 ([#214](https://github.com/validator-labs/validator/issues/214)) ([3313de5](https://github.com/validator-labs/validator/commit/3313de573b52b48a43765579dc750f9d6c52b951))
* **deps:** update module buf.build/gen/go/spectrocloud/spectro-cleanup/connectrpc/go to v1.14.0-20231213011348-5645e27c876a.1 ([#192](https://github.com/validator-labs/validator/issues/192)) ([31123a2](https://github.com/validator-labs/validator/commit/31123a25e9c4a1b3f99645113fc78dc178f3814e))
* **deps:** update module buf.build/gen/go/spectrocloud/spectro-cleanup/connectrpc/go to v1.15.0-20240205164452-95dfd137cb54.1 ([#217](https://github.com/validator-labs/validator/issues/217)) ([a57c42e](https://github.com/validator-labs/validator/commit/a57c42ea6ba0565b6941b8fe387c85d32a50a9d4))
* **deps:** update module buf.build/gen/go/spectrocloud/spectro-cleanup/protocolbuffers/go to v1.32.0-20231213011348-5645e27c876a.1 ([#193](https://github.com/validator-labs/validator/issues/193)) ([2abdb84](https://github.com/validator-labs/validator/commit/2abdb848f6de7b4b3d8206cd3c5edf8391cea574))
* **deps:** update module connectrpc.com/connect to v1.14.0 ([#165](https://github.com/validator-labs/validator/issues/165)) ([4a7f94a](https://github.com/validator-labs/validator/commit/4a7f94abc464d77d0f162ae9b71e73bd40ff7ee8))
* **deps:** update module github.com/go-logr/logr to v1.3.0 ([#77](https://github.com/validator-labs/validator/issues/77)) ([532fd6f](https://github.com/validator-labs/validator/commit/532fd6f82cf6a9f8322a74428e0834bb023ff67e))
* **deps:** update module github.com/go-logr/logr to v1.4.1 ([#164](https://github.com/validator-labs/validator/issues/164)) ([bfd0488](https://github.com/validator-labs/validator/commit/bfd04887ba3c430e14cdc7964eff21e64cd3e924))
* **deps:** update module github.com/onsi/ginkgo/v2 to v2.12.0 ([#31](https://github.com/validator-labs/validator/issues/31)) ([98a7aa7](https://github.com/validator-labs/validator/commit/98a7aa785946534db076a093c0715ac63782d72f))
* **deps:** update module github.com/onsi/ginkgo/v2 to v2.12.1 ([#52](https://github.com/validator-labs/validator/issues/52)) ([516693e](https://github.com/validator-labs/validator/commit/516693e3d150dfceb0a6be3a5f00bfbe260a1cb6))
* **deps:** update module github.com/onsi/ginkgo/v2 to v2.13.0 ([#57](https://github.com/validator-labs/validator/issues/57)) ([be32cb8](https://github.com/validator-labs/validator/commit/be32cb85ea38e8ddfd03e9d6837dddda1967c4b7))
* **deps:** update module github.com/onsi/ginkgo/v2 to v2.13.1 ([#95](https://github.com/validator-labs/validator/issues/95)) ([496ecad](https://github.com/validator-labs/validator/commit/496ecada5655f5760e46f7d647ce381f616ad56f))
* **deps:** update module github.com/onsi/ginkgo/v2 to v2.13.2 ([#138](https://github.com/validator-labs/validator/issues/138)) ([31746d6](https://github.com/validator-labs/validator/commit/31746d68c3f43ac3be1c0abacb62a35c57a7f1ce))
* **deps:** update module github.com/onsi/ginkgo/v2 to v2.14.0 ([#178](https://github.com/validator-labs/validator/issues/178)) ([b25c95c](https://github.com/validator-labs/validator/commit/b25c95c8b8b8cd2dd7f5e492fe33e39e3f6a1fa6))
* **deps:** update module github.com/onsi/ginkgo/v2 to v2.15.0 ([#182](https://github.com/validator-labs/validator/issues/182)) ([d421bd6](https://github.com/validator-labs/validator/commit/d421bd61e0e427cc95aa3a271f8773cfc91cc4b9))
* **deps:** update module github.com/onsi/ginkgo/v2 to v2.16.0 ([#224](https://github.com/validator-labs/validator/issues/224)) ([0cee706](https://github.com/validator-labs/validator/commit/0cee706f8046bace067ba55232de8af6a983e8df))
* **deps:** update module github.com/onsi/gomega to v1.27.10 ([#29](https://github.com/validator-labs/validator/issues/29)) ([8697124](https://github.com/validator-labs/validator/commit/8697124f495004f71ed9ba694ab8116880a4ae7f))
* **deps:** update module github.com/onsi/gomega to v1.28.0 ([#54](https://github.com/validator-labs/validator/issues/54)) ([e89431e](https://github.com/validator-labs/validator/commit/e89431e5685fc9ab83dd7dd2a87864f57b835bcb))
* **deps:** update module github.com/onsi/gomega to v1.28.1 ([#74](https://github.com/validator-labs/validator/issues/74)) ([eb7d990](https://github.com/validator-labs/validator/commit/eb7d990a840d14c5700dffc549de8f27feb0b614))
* **deps:** update module github.com/onsi/gomega to v1.29.0 ([#76](https://github.com/validator-labs/validator/issues/76)) ([ca58e8c](https://github.com/validator-labs/validator/commit/ca58e8c622c75ce345550e2115f8311428bfceb5))
* **deps:** update module github.com/onsi/gomega to v1.30.0 ([#82](https://github.com/validator-labs/validator/issues/82)) ([7bfa8f7](https://github.com/validator-labs/validator/commit/7bfa8f71aa77db2953aaba698f6acf4f2700a03e))
* **deps:** update module github.com/onsi/gomega to v1.31.0 ([#183](https://github.com/validator-labs/validator/issues/183)) ([6a4ee30](https://github.com/validator-labs/validator/commit/6a4ee30cc58348fd117c1b6384ae97534c606187))
* **deps:** update module github.com/onsi/gomega to v1.31.1 ([#188](https://github.com/validator-labs/validator/issues/188)) ([4a86ea9](https://github.com/validator-labs/validator/commit/4a86ea968446181b33acb3a45d3f17b9917acb8e))
* **deps:** update module github.com/slack-go/slack to v0.12.4 ([#216](https://github.com/validator-labs/validator/issues/216)) ([f281014](https://github.com/validator-labs/validator/commit/f281014d1722518bde0beb0389ec18eb1f87ce7b))
* **deps:** update module github.com/slack-go/slack to v0.12.5 ([#220](https://github.com/validator-labs/validator/issues/220)) ([37cf5a8](https://github.com/validator-labs/validator/commit/37cf5a87111407c723d7c9f2e9972ea37cc2b736))
* **deps:** update module k8s.io/klog/v2 to v2.110.1 ([#78](https://github.com/validator-labs/validator/issues/78)) ([8a79219](https://github.com/validator-labs/validator/commit/8a79219a40a4beb47182da4f5aea0d8045dd119f))
* **deps:** update module k8s.io/klog/v2 to v2.120.0 ([#176](https://github.com/validator-labs/validator/issues/176)) ([62bdb0a](https://github.com/validator-labs/validator/commit/62bdb0a802ea622796c90188c45f6712274a6d2c))
* **deps:** update module k8s.io/klog/v2 to v2.120.1 ([#185](https://github.com/validator-labs/validator/issues/185)) ([fad66e0](https://github.com/validator-labs/validator/commit/fad66e04a584456bfaf9991fd661d070f716afac))
* **deps:** update module sigs.k8s.io/controller-runtime to v0.16.1 ([#33](https://github.com/validator-labs/validator/issues/33)) ([94bf0ad](https://github.com/validator-labs/validator/commit/94bf0ad33b8524a82427b12320ba70493af9ac21))
* **deps:** update module sigs.k8s.io/controller-runtime to v0.16.2 ([#50](https://github.com/validator-labs/validator/issues/50)) ([1a07c2a](https://github.com/validator-labs/validator/commit/1a07c2a134bc26d86a17c04fcf71bae2f601c3e5))
* **deps:** update module sigs.k8s.io/controller-runtime to v0.16.3 ([#63](https://github.com/validator-labs/validator/issues/63)) ([6f79f8a](https://github.com/validator-labs/validator/commit/6f79f8af8f7a10c70ce403fadeb16d6eb9e13020))
* **deps:** update module sigs.k8s.io/yaml to v1.4.0 ([#98](https://github.com/validator-labs/validator/issues/98)) ([5f35bba](https://github.com/validator-labs/validator/commit/5f35bbac77502a944d6d5641e1e2f88f98cf7c79))
* dynamically set cleanup wait time ([#143](https://github.com/validator-labs/validator/issues/143)) ([13d0399](https://github.com/validator-labs/validator/commit/13d039915a4f5d9c66dae3c6938f55fc2ff210d6))
* ensure default helm cache is writable ([85b3286](https://github.com/validator-labs/validator/commit/85b3286fec22fd02c3f7ebee0d3c6fa74895bfdd))
* ensure plugin removal during Helm uninstall ([#111](https://github.com/validator-labs/validator/issues/111)) ([0917418](https://github.com/validator-labs/validator/commit/0917418b6ae3f2940bf8048c0cb09ca4056f21da))
* ensure State always remains failed if any conditions fail ([4f55823](https://github.com/validator-labs/validator/commit/4f558234af6c190186f8d82fb4a6e135d83030e8))
* ensure State always remains failed if any conditions fail ([711485f](https://github.com/validator-labs/validator/commit/711485f130804eab34de871f86c138adcb3bf257))
* ensure State always remains failed if any conditions fail ([67e8462](https://github.com/validator-labs/validator/commit/67e846293ce26e8f416fbb24d4b247d38f2a15aa))
* ensure ValidationResult Status resets to successful if all checks pass ([10237c5](https://github.com/validator-labs/validator/commit/10237c5a17adbc61288c6c6b9b981b072ea0e46f))
* format alertmanager cert properly ([#120](https://github.com/validator-labs/validator/issues/120)) ([f38635f](https://github.com/validator-labs/validator/commit/f38635f4de63a66f645bca3c9d6a239695a1ab2a))
* handle nil VRs w/ error ([#230](https://github.com/validator-labs/validator/issues/230)) ([ff479b5](https://github.com/validator-labs/validator/commit/ff479b5be609f6f0f77d7407b56445a92941e73e))
* handle VRs w/ multiple conditions ([8a3a243](https://github.com/validator-labs/validator/commit/8a3a2431dbd317dce80376c64d336317c688c2a3))
* helm chart CI ([46f37f0](https://github.com/validator-labs/validator/commit/46f37f0cea87e90e6effb85cb15128ab5970a621))
* include pkg/ in dockerfiles ([b45bb22](https://github.com/validator-labs/validator/commit/b45bb22532adc3e860aea56edf7ab3f3a95519fd))
* increase memory limit for helm upgrade ([660a80d](https://github.com/validator-labs/validator/commit/660a80d57fcae2dc3a16e610699b60a5749e0786))
* NET_ADMIN -&gt; NET_RAW ([122cc80](https://github.com/validator-labs/validator/commit/122cc808ed6b83eb6a33dc38e031d86805440784))
* omit conditions for uninstalled plugins ([c9f430d](https://github.com/validator-labs/validator/commit/c9f430d81bfdbb077edfa5d3cc48f314bf831c45))
* omit secret data from ValidatorConfig ([#125](https://github.com/validator-labs/validator/issues/125)) ([e96d4fe](https://github.com/validator-labs/validator/commit/e96d4fe3cee5bb8791ea95dcdae471f111255798))
* omit sink state from hash to avoid double-emitting on VR creation ([071b960](https://github.com/validator-labs/validator/commit/071b9602fee35262a66a9f403ceb878ac4a422b4))
* preserve VC annotations when updating plugin hashes ([19c9463](https://github.com/validator-labs/validator/commit/19c9463a4ed7b516731fdaf76cf487e682a6a2c4))
* prevent extraneous sink emissions ([22de119](https://github.com/validator-labs/validator/commit/22de119a3ce93c8fb751473f5250d55446012d83))
* quote all optional fields in sink secret ([e0a1365](https://github.com/validator-labs/validator/commit/e0a1365d577f329c8dc4d0bca8f6d3eb25e0c9c5))
* release image push repo ([4a2aca6](https://github.com/validator-labs/validator/commit/4a2aca6ecbfeca48ed4dd7566441923815281432))
* remove redundant b64 in vsphere secret ([f7f0555](https://github.com/validator-labs/validator/commit/f7f0555f9ad87a6726406c0b39c172e545ad9067))
* resolve all gosec issues ([#158](https://github.com/validator-labs/validator/issues/158)) ([dbca19b](https://github.com/validator-labs/validator/commit/dbca19bc47ae73a1589ba2b561002a431b881d12))
* retry all status updates due to controller contention ([#114](https://github.com/validator-labs/validator/issues/114)) ([35f03a4](https://github.com/validator-labs/validator/commit/35f03a407a3d0bbcfd76c749908e4b1c9581afac))
* retry VR status updates ([21b3808](https://github.com/validator-labs/validator/commit/21b3808f36a621f89ddc22aa5362d4d7b47265b5))
* SafeUpdateValidationResult not handling all edge cases ([#104](https://github.com/validator-labs/validator/issues/104)) ([8f34e2f](https://github.com/validator-labs/validator/commit/8f34e2f677a2b70c3c931491ce8b5cd6ac7abd0b))
* SafeUpdateValidationResult: update VR spec and status ([#205](https://github.com/validator-labs/validator/issues/205)) ([972eb55](https://github.com/validator-labs/validator/commit/972eb550968c32447aeba5ab7acfcbc566f0d929))
* securityContext blocking MTU check w/ ping ([131e5d9](https://github.com/validator-labs/validator/commit/131e5d91015b54b470b61708bfe8675f7eb26a0e))
* update 2+ plugin conditions properly ([a12488f](https://github.com/validator-labs/validator/commit/a12488f9dc376d1b0a8b413791b3aa2e25b185a5))
* update plugin versions in default values.yaml ([6f7f90d](https://github.com/validator-labs/validator/commit/6f7f90d15fbeb105df6c5b6c568c35fa4f12121f))
* update RBAC in helm templates ([6ff735c](https://github.com/validator-labs/validator/commit/6ff735c891e40328bba0524f4b8b240d3e85a6c9))
* update VRs to support rule addition ([#198](https://github.com/validator-labs/validator/issues/198)) ([c8de386](https://github.com/validator-labs/validator/commit/c8de3861fd4fe639c2cd27aa76a7883a0f1ae6d2))
* use http to make request to gRPC server ([#157](https://github.com/validator-labs/validator/issues/157)) ([1c960f8](https://github.com/validator-labs/validator/commit/1c960f80e2014d87696f993a598ecd49d36fc84b))
* write Helm CA cert to disk ([#169](https://github.com/validator-labs/validator/issues/169)) ([51c7e6d](https://github.com/validator-labs/validator/commit/51c7e6df1d6d88621dada3e8fa2fa22a73d3361e))


### Other

* add coverage report ([#100](https://github.com/validator-labs/validator/issues/100)) ([5d24704](https://github.com/validator-labs/validator/commit/5d2470488e1a09ca0d0f623762ae926dea5ba3ef))
* add license ([065ef97](https://github.com/validator-labs/validator/commit/065ef97d16dadf35a54b84fe3cd1026e14f172d5))
* add public validation result utils ([528be5f](https://github.com/validator-labs/validator/commit/528be5f91e8bfb7d6f1530002dd99971e4983a7e))
* add pull_request test trigger ([3e75bdb](https://github.com/validator-labs/validator/commit/3e75bdbff56bab925ca98b81c57fe9f4b1a60471))
* bump AWS chart version ([c569524](https://github.com/validator-labs/validator/commit/c569524ad2cb0e7ca2995c1fde327abe82c61425))
* bump aws plugin version ([bac567e](https://github.com/validator-labs/validator/commit/bac567e69ac6a957b78f29b5bcec81d2f3b0c6b8))
* bump azure chart version ([9230953](https://github.com/validator-labs/validator/commit/9230953434f15ef1c8bc19658d8bc8e9156df74a))
* bump go version ([#199](https://github.com/validator-labs/validator/issues/199)) ([06d0a9a](https://github.com/validator-labs/validator/commit/06d0a9a9a1b0f60ad97dbb282c3d3a58bce41e52))
* bump plugin versions ([6658d19](https://github.com/validator-labs/validator/commit/6658d190022815cf56d1a345ae66d46fd260c10c))
* Bump vsphere plugin version to 0.0.15 ([#142](https://github.com/validator-labs/validator/issues/142)) ([8b69b33](https://github.com/validator-labs/validator/commit/8b69b33463280ab3f44330f037f2958c41367d9a))
* **deps:** bump golang.org/x/net from 0.16.0 to 0.17.0 ([#61](https://github.com/validator-labs/validator/issues/61)) ([eca7abd](https://github.com/validator-labs/validator/commit/eca7abd8da204cabb28d7fb6ee8c381d1cb60e7b))
* **deps:** pin codecov/codecov-action action to eaaf4be ([#105](https://github.com/validator-labs/validator/issues/105)) ([70c3a0d](https://github.com/validator-labs/validator/commit/70c3a0d834cccc0bc6400c12395b7952c3fb25bb))
* **deps:** pin ncruces/go-coverage-report action to a12281c ([#101](https://github.com/validator-labs/validator/issues/101)) ([bb5d6ac](https://github.com/validator-labs/validator/commit/bb5d6ac4b235013eeba198e2ba4fd86c5eeb93a2))
* **deps:** update actions/checkout action to v4 ([#39](https://github.com/validator-labs/validator/issues/39)) ([3c848b6](https://github.com/validator-labs/validator/commit/3c848b61c3294cdd1ceed376c3e4a48521221b6f))
* **deps:** update actions/checkout digest to 8ade135 ([#53](https://github.com/validator-labs/validator/issues/53)) ([d42a8a6](https://github.com/validator-labs/validator/commit/d42a8a6832c15b76eac882a7f35fc3a2f8ba842a))
* **deps:** update actions/checkout digest to b4ffde6 ([#64](https://github.com/validator-labs/validator/issues/64)) ([d9bbd21](https://github.com/validator-labs/validator/commit/d9bbd21fe962de4f14e0e734d697ebd2ceb7e144))
* **deps:** update actions/checkout digest to f43a0e5 ([#25](https://github.com/validator-labs/validator/issues/25)) ([fa0b3d9](https://github.com/validator-labs/validator/commit/fa0b3d95f74e70226f07114bffcde2e1b270ad33))
* **deps:** update actions/setup-go action to v5 ([#147](https://github.com/validator-labs/validator/issues/147)) ([335d452](https://github.com/validator-labs/validator/commit/335d452d73f4ec5d86ca9cb00d2d5cb9598c0c0b))
* **deps:** update actions/setup-go digest to 93397be ([#26](https://github.com/validator-labs/validator/issues/26)) ([e32d52b](https://github.com/validator-labs/validator/commit/e32d52bc1243451b6a0f6a27f228869da9497761))
* **deps:** update actions/setup-python action to v5 ([#146](https://github.com/validator-labs/validator/issues/146)) ([d8fec57](https://github.com/validator-labs/validator/commit/d8fec5758fb948ee758c5a1deb1f4c4d9fe86d63))
* **deps:** update actions/setup-python digest to 65d7f2d ([#65](https://github.com/validator-labs/validator/issues/65)) ([bdb95d0](https://github.com/validator-labs/validator/commit/bdb95d038149ed9eab6c8372018bb771b848157c))
* **deps:** update actions/upload-artifact action to v4 ([#154](https://github.com/validator-labs/validator/issues/154)) ([7792867](https://github.com/validator-labs/validator/commit/77928676b59c162411491144f42e75a28a07fec6))
* **deps:** update actions/upload-artifact digest to 1eb3cb2 ([#180](https://github.com/validator-labs/validator/issues/180)) ([e1d9cb7](https://github.com/validator-labs/validator/commit/e1d9cb7a3cfa552c84c657cd55845e13f87ddbea))
* **deps:** update actions/upload-artifact digest to 26f96df ([#190](https://github.com/validator-labs/validator/issues/190)) ([43897d9](https://github.com/validator-labs/validator/commit/43897d987f95e9078c88ee42f0badb1e3bc2453e))
* **deps:** update actions/upload-artifact digest to 694cdab ([#186](https://github.com/validator-labs/validator/issues/186)) ([d33add9](https://github.com/validator-labs/validator/commit/d33add9c8a0bde65eae799a0ed4c37af817a8ecc))
* **deps:** update actions/upload-artifact digest to a8a3f3a ([#43](https://github.com/validator-labs/validator/issues/43)) ([fc33646](https://github.com/validator-labs/validator/commit/fc3364674e45c9dd3fdbf15773a7fbe2f04e3395))
* **deps:** update anchore/sbom-action action to v0.15.0 ([#129](https://github.com/validator-labs/validator/issues/129)) ([961377b](https://github.com/validator-labs/validator/commit/961377b591c7c071f85046a8c5211ee6e161a38d))
* **deps:** update anchore/sbom-action action to v0.15.1 ([#145](https://github.com/validator-labs/validator/issues/145)) ([7cbb94c](https://github.com/validator-labs/validator/commit/7cbb94c17599865ea68e7fe1019cb93acad37524))
* **deps:** update anchore/sbom-action action to v0.15.2 ([#170](https://github.com/validator-labs/validator/issues/170)) ([a37185a](https://github.com/validator-labs/validator/commit/a37185abcff714d1eb00d5eee9fef7ef82f18bbf))
* **deps:** update anchore/sbom-action action to v0.15.3 ([#173](https://github.com/validator-labs/validator/issues/173)) ([3a5be4e](https://github.com/validator-labs/validator/commit/3a5be4eb483c051b26a6fa0143d0380708e2c7e5))
* **deps:** update anchore/sbom-action action to v0.15.4 ([#184](https://github.com/validator-labs/validator/issues/184)) ([a58d980](https://github.com/validator-labs/validator/commit/a58d980b2ba3285e8efd59d9b389cc8d66859b78))
* **deps:** update anchore/sbom-action action to v0.15.8 ([#189](https://github.com/validator-labs/validator/issues/189)) ([56d8a75](https://github.com/validator-labs/validator/commit/56d8a754bdfa950ab2e727780b0234dc9d4c6d6d))
* **deps:** update anchore/sbom-action action to v0.15.9 ([#225](https://github.com/validator-labs/validator/issues/225)) ([f157fc8](https://github.com/validator-labs/validator/commit/f157fc849cdc633a95fa16fdc12787ba5bc22ce9))
* **deps:** update azure/setup-helm action to v4 ([#223](https://github.com/validator-labs/validator/issues/223)) ([8b5bb76](https://github.com/validator-labs/validator/commit/8b5bb76fddabdd3adeb37b7232aa1de2e206d4a1))
* **deps:** update codecov/codecov-action digest to 4fe8c5f ([#191](https://github.com/validator-labs/validator/issues/191)) ([7f55aad](https://github.com/validator-labs/validator/commit/7f55aad95486e87774e9c8bffe4453ccce14fa5f))
* **deps:** update codecov/codecov-action digest to 54bcd87 ([#219](https://github.com/validator-labs/validator/issues/219)) ([3ebb8b1](https://github.com/validator-labs/validator/commit/3ebb8b1d98c96764b262a1ef03ad054078136752))
* **deps:** update codecov/codecov-action digest to ab904c4 ([#194](https://github.com/validator-labs/validator/issues/194)) ([053afa3](https://github.com/validator-labs/validator/commit/053afa31e65fe700a4973f502aea2ad63b1e51d9))
* **deps:** update codecov/codecov-action digest to e0b68c6 ([#197](https://github.com/validator-labs/validator/issues/197)) ([69fe200](https://github.com/validator-labs/validator/commit/69fe200a73f51c7f913d605f16c22091a029d91e))
* **deps:** update docker/build-push-action action to v5 ([#47](https://github.com/validator-labs/validator/issues/47)) ([40d0053](https://github.com/validator-labs/validator/commit/40d0053011c9d05582d96ab07c7e743867fbed8d))
* **deps:** update docker/build-push-action digest to 0a97817 ([#45](https://github.com/validator-labs/validator/issues/45)) ([91a951b](https://github.com/validator-labs/validator/commit/91a951b03d40db29020af68cb7b1a22db8cad19c))
* **deps:** update docker/build-push-action digest to 4a13e50 ([#119](https://github.com/validator-labs/validator/issues/119)) ([24b3bfc](https://github.com/validator-labs/validator/commit/24b3bfc927dc5d04fa77b58157bda2f2f18fcb12))
* **deps:** update docker/login-action action to v3 ([#48](https://github.com/validator-labs/validator/issues/48)) ([9c940d7](https://github.com/validator-labs/validator/commit/9c940d7acebf0e237ed5129cc3b25aca596a257f))
* **deps:** update docker/setup-buildx-action action to v3 ([#49](https://github.com/validator-labs/validator/issues/49)) ([a3fe730](https://github.com/validator-labs/validator/commit/a3fe730ae38b9655583be333e255e48767c2cf0c))
* **deps:** update docker/setup-buildx-action digest to 0d103c3 ([#221](https://github.com/validator-labs/validator/issues/221)) ([c8c02f2](https://github.com/validator-labs/validator/commit/c8c02f248c629ea56a87e4cd19d4f86e8bb909a6))
* **deps:** update docker/setup-buildx-action digest to 885d146 ([#28](https://github.com/validator-labs/validator/issues/28)) ([f4b1dd1](https://github.com/validator-labs/validator/commit/f4b1dd10a4170f2596bf2cebf64a363d21efcc44))
* **deps:** update gcr.io/kubebuilder/kube-rbac-proxy docker tag to v0.14.4 ([#59](https://github.com/validator-labs/validator/issues/59)) ([78847f5](https://github.com/validator-labs/validator/commit/78847f54c35f6fc684d333a286e1315cc558e6e2))
* **deps:** update gcr.io/kubebuilder/kube-rbac-proxy docker tag to v0.15.0 ([#72](https://github.com/validator-labs/validator/issues/72)) ([4f0fc63](https://github.com/validator-labs/validator/commit/4f0fc630bc3ab969d6411fb1d31d968e313a20be))
* **deps:** update gcr.io/spectro-images-public/golang docker tag to v1.22 ([#174](https://github.com/validator-labs/validator/issues/174)) ([d9beecf](https://github.com/validator-labs/validator/commit/d9beecfff64c26ddce8206de808667b1ba3e0f24))
* **deps:** update google-github-actions/release-please-action action to v4 ([#144](https://github.com/validator-labs/validator/issues/144)) ([c4d008c](https://github.com/validator-labs/validator/commit/c4d008c391fbf4c0bcd932668ed0684b571fa7fe))
* **deps:** update google-github-actions/release-please-action digest to 4c5670f ([#66](https://github.com/validator-labs/validator/issues/66)) ([2c24d48](https://github.com/validator-labs/validator/commit/2c24d48035b7ebddbbd20ca182e3352fa6c2f09e))
* **deps:** update google-github-actions/release-please-action digest to cc61a07 ([#152](https://github.com/validator-labs/validator/issues/152)) ([908de87](https://github.com/validator-labs/validator/commit/908de87b359b6e22cb16e69f981d793aec18aa71))
* **deps:** update google-github-actions/release-please-action digest to db8f2c6 ([#81](https://github.com/validator-labs/validator/issues/81)) ([20956a3](https://github.com/validator-labs/validator/commit/20956a3fa864d5cc0e37349ba0632da61879b8b8))
* **deps:** update helm/chart-testing-action action to v2.6.0 ([#79](https://github.com/validator-labs/validator/issues/79)) ([3272b22](https://github.com/validator-labs/validator/commit/3272b226df2586344efd1dcf3f077483ca3f45a7))
* **deps:** update helm/chart-testing-action action to v2.6.1 ([#80](https://github.com/validator-labs/validator/issues/80)) ([cbb430e](https://github.com/validator-labs/validator/commit/cbb430e70a8aad6558816d3ce2c5c42cacefcab8))
* enable renovate automerges ([84ad3cd](https://github.com/validator-labs/validator/commit/84ad3cdee59ed634e5f9577d801dc13701958e46))
* fix broken build link in README ([#222](https://github.com/validator-labs/validator/issues/222)) ([9e78ac2](https://github.com/validator-labs/validator/commit/9e78ac28a970b3b07a75c3346304358fbf6e91f5))
* fix Helm chart lint error ([c938c89](https://github.com/validator-labs/validator/commit/c938c89947371fa3e3e3e2d43044749a413ec443))
* **main:** release 0.0.1 ([#9](https://github.com/validator-labs/validator/issues/9)) ([62fc470](https://github.com/validator-labs/validator/commit/62fc470c14810b9dc9699f3846279999dcc290a3))
* **main:** release 0.0.10 ([#68](https://github.com/validator-labs/validator/issues/68)) ([edc5f33](https://github.com/validator-labs/validator/commit/edc5f33fdaf546e8eeedba1f5ba1647f61c789f1))
* **main:** release 0.0.10 ([#70](https://github.com/validator-labs/validator/issues/70)) ([6c5e4fc](https://github.com/validator-labs/validator/commit/6c5e4fcc1182862e3902044e543309800e96b31e))
* **main:** release 0.0.10 ([#75](https://github.com/validator-labs/validator/issues/75)) ([32e4025](https://github.com/validator-labs/validator/commit/32e4025ba21223c7214e5378bb060769b931e685))
* **main:** release 0.0.11 ([#85](https://github.com/validator-labs/validator/issues/85)) ([081ded0](https://github.com/validator-labs/validator/commit/081ded0ac22d0432bcf04dbed2be9d281a6f38b9))
* **main:** release 0.0.13 ([#91](https://github.com/validator-labs/validator/issues/91)) ([15a5342](https://github.com/validator-labs/validator/commit/15a5342624deba32fc3191ea4583befa942fa157))
* **main:** release 0.0.14 ([#92](https://github.com/validator-labs/validator/issues/92)) ([42e88ff](https://github.com/validator-labs/validator/commit/42e88ff8c7cc5927e6ae78644d146aba7b7f4c92))
* **main:** release 0.0.15 ([#93](https://github.com/validator-labs/validator/issues/93)) ([ee7685a](https://github.com/validator-labs/validator/commit/ee7685a5661d3be1f7729cc48f93ad020351c6d7))
* **main:** release 0.0.16 ([#99](https://github.com/validator-labs/validator/issues/99)) ([93f1b88](https://github.com/validator-labs/validator/commit/93f1b885e47b356a6a55d913bf95bcddb7210309))
* **main:** release 0.0.17 ([#102](https://github.com/validator-labs/validator/issues/102)) ([51e757e](https://github.com/validator-labs/validator/commit/51e757edb211684ebeddd0ed26e57b73e8333cfd))
* **main:** release 0.0.18 ([#103](https://github.com/validator-labs/validator/issues/103)) ([315019a](https://github.com/validator-labs/validator/commit/315019a2a1df8d9332dffe8bb179a6ea360e82dc))
* **main:** release 0.0.19 ([#110](https://github.com/validator-labs/validator/issues/110)) ([acb9902](https://github.com/validator-labs/validator/commit/acb99021461f93334fef2146cd1e8f81b7b066ad))
* **main:** release 0.0.2 ([#10](https://github.com/validator-labs/validator/issues/10)) ([3c10801](https://github.com/validator-labs/validator/commit/3c10801af9c9c7801338a5488f6525a43b610aac))
* **main:** release 0.0.20 ([#115](https://github.com/validator-labs/validator/issues/115)) ([e2e98e7](https://github.com/validator-labs/validator/commit/e2e98e7d3f57cc9caf7ba438647ef9e167aa9969))
* **main:** release 0.0.21 ([#116](https://github.com/validator-labs/validator/issues/116)) ([137df6b](https://github.com/validator-labs/validator/commit/137df6b06828fc054fb76eae5bd514571a245ebe))
* **main:** release 0.0.22 ([#123](https://github.com/validator-labs/validator/issues/123)) ([cc07547](https://github.com/validator-labs/validator/commit/cc0754767011bc752819da56052d9639c098ec77))
* **main:** release 0.0.23 ([#124](https://github.com/validator-labs/validator/issues/124)) ([0331832](https://github.com/validator-labs/validator/commit/03318323a8067b75f9c407c00dbf5cf869a08c77))
* **main:** release 0.0.24 ([#126](https://github.com/validator-labs/validator/issues/126)) ([5f55fdf](https://github.com/validator-labs/validator/commit/5f55fdfe46d0c4f79924cd54a3f40ab31efd1796))
* **main:** release 0.0.25 ([#128](https://github.com/validator-labs/validator/issues/128)) ([7b3af63](https://github.com/validator-labs/validator/commit/7b3af63651e50f6dc7393b0de2361d1605d249d1))
* **main:** release 0.0.26 ([#130](https://github.com/validator-labs/validator/issues/130)) ([a697310](https://github.com/validator-labs/validator/commit/a6973108c87c58a0df551cffb7947f910bfe0f14))
* **main:** release 0.0.27 ([#136](https://github.com/validator-labs/validator/issues/136)) ([90c6f41](https://github.com/validator-labs/validator/commit/90c6f412369881144428d2ce71c0f8b97937f926))
* **main:** release 0.0.28 ([#140](https://github.com/validator-labs/validator/issues/140)) ([5c69568](https://github.com/validator-labs/validator/commit/5c69568a635eaebf852915ecdcdc0441b403b7c0))
* **main:** release 0.0.29 ([#148](https://github.com/validator-labs/validator/issues/148)) ([6f32ca8](https://github.com/validator-labs/validator/commit/6f32ca836ef5b6e319b957bef1f0749550f9cb0d))
* **main:** release 0.0.3 ([#18](https://github.com/validator-labs/validator/issues/18)) ([9d10ae6](https://github.com/validator-labs/validator/commit/9d10ae63a5bdfa46537a11c61a7c15a0a7553aaf))
* **main:** release 0.0.30 ([#160](https://github.com/validator-labs/validator/issues/160)) ([ad78ea4](https://github.com/validator-labs/validator/commit/ad78ea4478fea49111fb3f182a8d613a46fe5dc6))
* **main:** release 0.0.31 ([#167](https://github.com/validator-labs/validator/issues/167)) ([5edba38](https://github.com/validator-labs/validator/commit/5edba3885385fb239f74e94777f9badbd4cb1355))
* **main:** release 0.0.32 ([#168](https://github.com/validator-labs/validator/issues/168)) ([1a70912](https://github.com/validator-labs/validator/commit/1a70912d91c91826ce16273e29d165ef9df28700))
* **main:** release 0.0.33 ([#172](https://github.com/validator-labs/validator/issues/172)) ([ba46727](https://github.com/validator-labs/validator/commit/ba467273b87618ff06e90c5bd9a0b07a1b75e231))
* **main:** release 0.0.34 ([#202](https://github.com/validator-labs/validator/issues/202)) ([d0a1b30](https://github.com/validator-labs/validator/commit/d0a1b30efe83ed1dcf6e0e1691e227aa8e4dede1))
* **main:** release 0.0.35 ([#211](https://github.com/validator-labs/validator/issues/211)) ([c271c73](https://github.com/validator-labs/validator/commit/c271c739fa8622eda636feda97aafda934b419dd))
* **main:** release 0.0.4 ([#19](https://github.com/validator-labs/validator/issues/19)) ([98fd554](https://github.com/validator-labs/validator/commit/98fd554209f11f93f001eec9ac32ed7c55c13166))
* **main:** release 0.0.5 ([#23](https://github.com/validator-labs/validator/issues/23)) ([ae58d10](https://github.com/validator-labs/validator/commit/ae58d10a0a931543aee6157c190b9cb14f37d2a2))
* **main:** release 0.0.6 ([#38](https://github.com/validator-labs/validator/issues/38)) ([849b4ca](https://github.com/validator-labs/validator/commit/849b4ca0d60154c33bbbcccef5cdc9d3b70479e1))
* **main:** release 0.0.7 ([#42](https://github.com/validator-labs/validator/issues/42)) ([85f2933](https://github.com/validator-labs/validator/commit/85f2933a4a1d6cab927f225bd46f0664d1102232))
* **main:** release 0.0.8 ([#44](https://github.com/validator-labs/validator/issues/44)) ([b452df4](https://github.com/validator-labs/validator/commit/b452df441271f814e9cc37326d8a6d432241a388))
* **main:** release 0.0.9 ([#58](https://github.com/validator-labs/validator/issues/58)) ([6d31932](https://github.com/validator-labs/validator/commit/6d319325bef13ab0ab9b77707971b60fcbac3d76))
* Merge branch 'main' of https://github.com/validator-labs/validator ([42cf0ef](https://github.com/validator-labs/validator/commit/42cf0efc8ee8f8187b260ac0f4d1f8b1b9db6c16))
* Merge branch 'main' of https://github.com/validator-labs/validator ([a2700a2](https://github.com/validator-labs/validator/commit/a2700a26661d6c7668965cc5b5a377adfd49d109))
* release 0.0.1 ([a23551a](https://github.com/validator-labs/validator/commit/a23551a1984d43d9acbc7de3cacad6ee928cc517))
* release 0.0.10 ([65ce207](https://github.com/validator-labs/validator/commit/65ce2076727fd3d1da49afd884329c55a6394c91))
* release 0.0.10 ([9a324e9](https://github.com/validator-labs/validator/commit/9a324e9e666b0da68a1e0c3be225ab19bfd04a6b))
* release 0.0.11 ([f694577](https://github.com/validator-labs/validator/commit/f694577bb2b0fae8127935c3f9fd9e90f14fe328))
* release 0.0.12 ([4903786](https://github.com/validator-labs/validator/commit/49037866402a7f16ef2c15cf172a11626392b9ff))
* release 0.0.13 ([1233488](https://github.com/validator-labs/validator/commit/1233488756f02cd3de9ee2a0d800cb29522545c8))
* release 0.0.14 ([2471411](https://github.com/validator-labs/validator/commit/247141147841953b69fe69df297244d3119b8d40))
* release 0.0.15 ([d574854](https://github.com/validator-labs/validator/commit/d57485496dc90f4b6e421d978bdf0451edfcc59c))
* release 0.0.16 ([f031904](https://github.com/validator-labs/validator/commit/f031904a18fb5854586d58145842689d12028360))
* release 0.0.2 ([40cdd88](https://github.com/validator-labs/validator/commit/40cdd88ebb8b75f9908c5dab6aa29337f5d778d8))
* release 0.0.25 ([7af7723](https://github.com/validator-labs/validator/commit/7af772348e3562d22a1afb58957d82959456443e))
* release 0.0.3 ([5b2473d](https://github.com/validator-labs/validator/commit/5b2473dce974a160b14640e86e88028f66c94f5e))
* release 0.0.36 ([f004e5b](https://github.com/validator-labs/validator/commit/f004e5b35b8c622f3e2081ce9a8751d8bd94d9a3))
* release 0.0.4 ([28f8418](https://github.com/validator-labs/validator/commit/28f8418356694ad5c86b97d2b9df9d51c2f6d279))
* release 0.0.5 ([24e9712](https://github.com/validator-labs/validator/commit/24e9712ffbf6fb8d33333d5d8c063f3935c7ceae))
* release 0.0.6 ([10421d5](https://github.com/validator-labs/validator/commit/10421d59a5fbc4dd840a6aac9f24657ee8d1be7c))
* release 0.0.9 ([3cbcfa9](https://github.com/validator-labs/validator/commit/3cbcfa9e3f1e5c75dcfc083df1494aab0c41c31f))
* update AWS auth config ([#134](https://github.com/validator-labs/validator/issues/134)) ([9513e43](https://github.com/validator-labs/validator/commit/9513e43a9860cc1ac2f1fcea01e8d7727d81b11c))
* update AWS plugin version; fix default values.yaml ([516213b](https://github.com/validator-labs/validator/commit/516213ba35d80c4e8692b1448b22f6e1b9756c5d))
* Update default values.yaml ([f9af283](https://github.com/validator-labs/validator/commit/f9af2832bd652c9ea404d08936a6879930d29753))
* update network plugin values ([#135](https://github.com/validator-labs/validator/issues/135)) ([1049435](https://github.com/validator-labs/validator/commit/104943556a73e3aa4e5acd80c99542222035a867))
* update README ([f8254d4](https://github.com/validator-labs/validator/commit/f8254d4a0bca6db90523ace43503cf6d80d3af30))


### Docs

* add valid8or-plugin-network to default values.yaml ([1aada24](https://github.com/validator-labs/validator/commit/1aada2492d518e7ed50934d2cb6e184da5d1e031))
* issue template addition ([#109](https://github.com/validator-labs/validator/issues/109)) ([36ce4a1](https://github.com/validator-labs/validator/commit/36ce4a1d5630c22b39d481bc45641c5c06e6db04))
* refer to secret templates in values.yaml ([eeae1e7](https://github.com/validator-labs/validator/commit/eeae1e710a2dd584ea74e8017ddb359a165c9079))
* update chart description ([0a5635f](https://github.com/validator-labs/validator/commit/0a5635f94f949c39592dbe7f20e5301f4836f291))
* update README, fix release-please annotations ([c2c96e8](https://github.com/validator-labs/validator/commit/c2c96e8e3e91820826242b36d6760ab1d2530baf))


### Refactoring

* accept VR in HandleNewValidationResult for flexibility in plugins ([#113](https://github.com/validator-labs/validator/issues/113)) ([1ead151](https://github.com/validator-labs/validator/commit/1ead15146156ac278aedb2a77cab0604488fda4f))
* change devspace port; always use exponential backoff; tidy validation result helpers ([c866429](https://github.com/validator-labs/validator/commit/c866429772e8d301916eed7ea8908b811cad3e7c))
* expose helm package ([2a34e0a](https://github.com/validator-labs/validator/commit/2a34e0ae780b287a8a1a48ee00f5016e667c8304))
* expose sink types ([#117](https://github.com/validator-labs/validator/issues/117)) ([f28d8af](https://github.com/validator-labs/validator/commit/f28d8afc5092998189da4112e8a41febdadf1c96))
* standardize get CR in Reconcile ([9fbfff0](https://github.com/validator-labs/validator/commit/9fbfff0c059debab0c9c6044c360af07c8cd4382))
* use header block for msg titles ([a6ce7af](https://github.com/validator-labs/validator/commit/a6ce7af748dcb6cbd50c5d13e1a2c9a724e9a328))
* valid8or -&gt; validator ([#67](https://github.com/validator-labs/validator/issues/67)) ([ff88026](https://github.com/validator-labs/validator/commit/ff8802656d8115dd6afbbfbaf12613c4f205feb5))

## [0.0.35](https://github.com/validator-labs/validator/compare/v0.0.34...v0.0.35) (2024-02-12)


### Features

* OCI repository support for Helm charts ([#207](https://github.com/validator-labs/validator/issues/207)) ([4b25d79](https://github.com/validator-labs/validator/commit/4b25d7980c9aac07a39294a27ca3540e83dce1fd))


### Bug Fixes

* ensure default helm cache is writable ([85b3286](https://github.com/validator-labs/validator/commit/85b3286fec22fd02c3f7ebee0d3c6fa74895bfdd))

## [0.0.34](https://github.com/validator-labs/validator/compare/v0.0.33...v0.0.34) (2024-02-06)


### Bug Fixes

* **deps:** update kubernetes packages to v0.29.1 ([#153](https://github.com/validator-labs/validator/issues/153)) ([60bc244](https://github.com/validator-labs/validator/commit/60bc244537838596be433c3e2a59ce66fcb80350))
* SafeUpdateValidationResult: update VR spec and status ([#205](https://github.com/validator-labs/validator/issues/205)) ([972eb55](https://github.com/validator-labs/validator/commit/972eb550968c32447aeba5ab7acfcbc566f0d929))

## [0.0.33](https://github.com/validator-labs/validator/compare/v0.0.32...v0.0.33) (2024-02-05)


### Bug Fixes

* **deps:** update golang.org/x/exp digest to 0dcbfd6 ([#175](https://github.com/validator-labs/validator/issues/175)) ([31469da](https://github.com/validator-labs/validator/commit/31469da899767e2d8a6766f609a9e280b6e49609))
* **deps:** update golang.org/x/exp digest to 1b97071 ([#187](https://github.com/validator-labs/validator/issues/187)) ([27581fa](https://github.com/validator-labs/validator/commit/27581fa0282aee29a032bfff15a29f23f9ff6390))
* **deps:** update golang.org/x/exp digest to be819d1 ([#171](https://github.com/validator-labs/validator/issues/171)) ([6d5d2a2](https://github.com/validator-labs/validator/commit/6d5d2a2c5fe210647db4d347b83b81fe7486e31b))
* **deps:** update golang.org/x/exp digest to db7319d ([#179](https://github.com/validator-labs/validator/issues/179)) ([04e3a3d](https://github.com/validator-labs/validator/commit/04e3a3d87324910cb8f32dc46865386e20f43378))
* **deps:** update module buf.build/gen/go/spectrocloud/spectro-cleanup/connectrpc/go to v1.14.0-20231213011348-5645e27c876a.1 ([#192](https://github.com/validator-labs/validator/issues/192)) ([31123a2](https://github.com/validator-labs/validator/commit/31123a25e9c4a1b3f99645113fc78dc178f3814e))
* **deps:** update module buf.build/gen/go/spectrocloud/spectro-cleanup/protocolbuffers/go to v1.32.0-20231213011348-5645e27c876a.1 ([#193](https://github.com/validator-labs/validator/issues/193)) ([2abdb84](https://github.com/validator-labs/validator/commit/2abdb848f6de7b4b3d8206cd3c5edf8391cea574))
* **deps:** update module github.com/onsi/ginkgo/v2 to v2.14.0 ([#178](https://github.com/validator-labs/validator/issues/178)) ([b25c95c](https://github.com/validator-labs/validator/commit/b25c95c8b8b8cd2dd7f5e492fe33e39e3f6a1fa6))
* **deps:** update module github.com/onsi/ginkgo/v2 to v2.15.0 ([#182](https://github.com/validator-labs/validator/issues/182)) ([d421bd6](https://github.com/validator-labs/validator/commit/d421bd61e0e427cc95aa3a271f8773cfc91cc4b9))
* **deps:** update module github.com/onsi/gomega to v1.31.0 ([#183](https://github.com/validator-labs/validator/issues/183)) ([6a4ee30](https://github.com/validator-labs/validator/commit/6a4ee30cc58348fd117c1b6384ae97534c606187))
* **deps:** update module github.com/onsi/gomega to v1.31.1 ([#188](https://github.com/validator-labs/validator/issues/188)) ([4a86ea9](https://github.com/validator-labs/validator/commit/4a86ea968446181b33acb3a45d3f17b9917acb8e))
* **deps:** update module k8s.io/klog/v2 to v2.120.0 ([#176](https://github.com/validator-labs/validator/issues/176)) ([62bdb0a](https://github.com/validator-labs/validator/commit/62bdb0a802ea622796c90188c45f6712274a6d2c))
* **deps:** update module k8s.io/klog/v2 to v2.120.1 ([#185](https://github.com/validator-labs/validator/issues/185)) ([fad66e0](https://github.com/validator-labs/validator/commit/fad66e04a584456bfaf9991fd661d070f716afac))
* update VRs to support rule addition ([#198](https://github.com/validator-labs/validator/issues/198)) ([c8de386](https://github.com/validator-labs/validator/commit/c8de3861fd4fe639c2cd27aa76a7883a0f1ae6d2))


### Other

* bump go version ([#199](https://github.com/validator-labs/validator/issues/199)) ([06d0a9a](https://github.com/validator-labs/validator/commit/06d0a9a9a1b0f60ad97dbb282c3d3a58bce41e52))
* **deps:** update actions/upload-artifact digest to 1eb3cb2 ([#180](https://github.com/validator-labs/validator/issues/180)) ([e1d9cb7](https://github.com/validator-labs/validator/commit/e1d9cb7a3cfa552c84c657cd55845e13f87ddbea))
* **deps:** update actions/upload-artifact digest to 26f96df ([#190](https://github.com/validator-labs/validator/issues/190)) ([43897d9](https://github.com/validator-labs/validator/commit/43897d987f95e9078c88ee42f0badb1e3bc2453e))
* **deps:** update actions/upload-artifact digest to 694cdab ([#186](https://github.com/validator-labs/validator/issues/186)) ([d33add9](https://github.com/validator-labs/validator/commit/d33add9c8a0bde65eae799a0ed4c37af817a8ecc))
* **deps:** update anchore/sbom-action action to v0.15.2 ([#170](https://github.com/validator-labs/validator/issues/170)) ([a37185a](https://github.com/validator-labs/validator/commit/a37185abcff714d1eb00d5eee9fef7ef82f18bbf))
* **deps:** update anchore/sbom-action action to v0.15.3 ([#173](https://github.com/validator-labs/validator/issues/173)) ([3a5be4e](https://github.com/validator-labs/validator/commit/3a5be4eb483c051b26a6fa0143d0380708e2c7e5))
* **deps:** update anchore/sbom-action action to v0.15.4 ([#184](https://github.com/validator-labs/validator/issues/184)) ([a58d980](https://github.com/validator-labs/validator/commit/a58d980b2ba3285e8efd59d9b389cc8d66859b78))
* **deps:** update anchore/sbom-action action to v0.15.8 ([#189](https://github.com/validator-labs/validator/issues/189)) ([56d8a75](https://github.com/validator-labs/validator/commit/56d8a754bdfa950ab2e727780b0234dc9d4c6d6d))
* **deps:** update codecov/codecov-action digest to 4fe8c5f ([#191](https://github.com/validator-labs/validator/issues/191)) ([7f55aad](https://github.com/validator-labs/validator/commit/7f55aad95486e87774e9c8bffe4453ccce14fa5f))
* **deps:** update codecov/codecov-action digest to ab904c4 ([#194](https://github.com/validator-labs/validator/issues/194)) ([053afa3](https://github.com/validator-labs/validator/commit/053afa31e65fe700a4973f502aea2ad63b1e51d9))
* **deps:** update codecov/codecov-action digest to e0b68c6 ([#197](https://github.com/validator-labs/validator/issues/197)) ([69fe200](https://github.com/validator-labs/validator/commit/69fe200a73f51c7f913d605f16c22091a029d91e))
* **deps:** update gcr.io/spectro-images-public/golang docker tag to v1.22 ([#174](https://github.com/validator-labs/validator/issues/174)) ([d9beecf](https://github.com/validator-labs/validator/commit/d9beecfff64c26ddce8206de808667b1ba3e0f24))

## [0.0.32](https://github.com/validator-labs/validator/compare/v0.0.31...v0.0.32) (2023-12-27)


### Bug Fixes

* **deps:** update golang.org/x/exp digest to 02704c9 ([#161](https://github.com/validator-labs/validator/issues/161)) ([72990b7](https://github.com/validator-labs/validator/commit/72990b7d01a917d0d2bbdedcfa41e019af993b34))
* **deps:** update module connectrpc.com/connect to v1.14.0 ([#165](https://github.com/validator-labs/validator/issues/165)) ([4a7f94a](https://github.com/validator-labs/validator/commit/4a7f94abc464d77d0f162ae9b71e73bd40ff7ee8))
* **deps:** update module github.com/go-logr/logr to v1.4.1 ([#164](https://github.com/validator-labs/validator/issues/164)) ([bfd0488](https://github.com/validator-labs/validator/commit/bfd04887ba3c430e14cdc7964eff21e64cd3e924))
* write Helm CA cert to disk ([#169](https://github.com/validator-labs/validator/issues/169)) ([51c7e6d](https://github.com/validator-labs/validator/commit/51c7e6df1d6d88621dada3e8fa2fa22a73d3361e))


### Other

* **deps:** update actions/upload-artifact action to v4 ([#154](https://github.com/validator-labs/validator/issues/154)) ([7792867](https://github.com/validator-labs/validator/commit/77928676b59c162411491144f42e75a28a07fec6))

## [0.0.31](https://github.com/validator-labs/validator/compare/v0.0.30...v0.0.31) (2023-12-26)


### Features

* add caFile Helm option ([#166](https://github.com/validator-labs/validator/issues/166)) ([2425599](https://github.com/validator-labs/validator/commit/24255997119b75edde4e8f94568a906c7da6b37f))

## [0.0.30](https://github.com/validator-labs/validator/compare/v0.0.29...v0.0.30) (2023-12-19)


### Features

* update helm chart to enable fast cleanup in spectro-cleanup via gRPC endpoint ([#156](https://github.com/validator-labs/validator/issues/156)) ([6f91745](https://github.com/validator-labs/validator/commit/6f91745c9db6e92071e691a74942c0bb53692384))


### Bug Fixes

* **deps:** update golang.org/x/exp digest to 73b9e39 ([#151](https://github.com/validator-labs/validator/issues/151)) ([bfaf2de](https://github.com/validator-labs/validator/commit/bfaf2de249f6edfb583ff5870be7572ddd2a55ee))


### Other

* **deps:** update google-github-actions/release-please-action digest to cc61a07 ([#152](https://github.com/validator-labs/validator/issues/152)) ([908de87](https://github.com/validator-labs/validator/commit/908de87b359b6e22cb16e69f981d793aec18aa71))

## [0.0.29](https://github.com/validator-labs/validator/compare/v0.0.28...v0.0.29) (2023-12-19)


### Features

* implement client to send FinalizeCleanup requests to spectro-cleanup ([#155](https://github.com/validator-labs/validator/issues/155)) ([2ae0348](https://github.com/validator-labs/validator/commit/2ae03480b48d7b1bffb20227889c7f1c2a12f60e))


### Bug Fixes

* resolve all gosec issues ([#158](https://github.com/validator-labs/validator/issues/158)) ([dbca19b](https://github.com/validator-labs/validator/commit/dbca19bc47ae73a1589ba2b561002a431b881d12))
* use http to make request to gRPC server ([#157](https://github.com/validator-labs/validator/issues/157)) ([1c960f8](https://github.com/validator-labs/validator/commit/1c960f80e2014d87696f993a598ecd49d36fc84b))


### Other

* **deps:** update actions/setup-go action to v5 ([#147](https://github.com/validator-labs/validator/issues/147)) ([335d452](https://github.com/validator-labs/validator/commit/335d452d73f4ec5d86ca9cb00d2d5cb9598c0c0b))
* **deps:** update actions/setup-python action to v5 ([#146](https://github.com/validator-labs/validator/issues/146)) ([d8fec57](https://github.com/validator-labs/validator/commit/d8fec5758fb948ee758c5a1deb1f4c4d9fe86d63))

## [0.0.28](https://github.com/validator-labs/validator/compare/v0.0.27...v0.0.28) (2023-12-04)


### Features

* add OCI plugin ([#139](https://github.com/validator-labs/validator/issues/139)) ([3ae7b70](https://github.com/validator-labs/validator/commit/3ae7b70920ed27f8ad1325fdbe0461c71ae48a71))


### Bug Fixes

* delete plugins in parallel to avoid hitting timeouts ([#141](https://github.com/validator-labs/validator/issues/141)) ([a3fc0dc](https://github.com/validator-labs/validator/commit/a3fc0dc56c7c1964b2ee38b666a866c78711010e))
* dynamically set cleanup wait time ([#143](https://github.com/validator-labs/validator/issues/143)) ([13d0399](https://github.com/validator-labs/validator/commit/13d039915a4f5d9c66dae3c6938f55fc2ff210d6))


### Other

* Bump vsphere plugin version to 0.0.15 ([#142](https://github.com/validator-labs/validator/issues/142)) ([8b69b33](https://github.com/validator-labs/validator/commit/8b69b33463280ab3f44330f037f2958c41367d9a))
* **deps:** update anchore/sbom-action action to v0.15.1 ([#145](https://github.com/validator-labs/validator/issues/145)) ([7cbb94c](https://github.com/validator-labs/validator/commit/7cbb94c17599865ea68e7fe1019cb93acad37524))
* **deps:** update google-github-actions/release-please-action action to v4 ([#144](https://github.com/validator-labs/validator/issues/144)) ([c4d008c](https://github.com/validator-labs/validator/commit/c4d008c391fbf4c0bcd932668ed0684b571fa7fe))

## [0.0.27](https://github.com/validator-labs/validator/compare/v0.0.26...v0.0.27) (2023-11-29)


### Features

* support proxy configuration (env vars & CA certificate) ([#137](https://github.com/validator-labs/validator/issues/137)) ([63c3bc8](https://github.com/validator-labs/validator/commit/63c3bc8497766564d7e76a3da2dabc53ba3b7d54))


### Bug Fixes

* add yaml tags to ValidatorConfig types ([63afc70](https://github.com/validator-labs/validator/commit/63afc70819a4fffc6d90ba40cea12bca2577e743))
* **deps:** update module github.com/onsi/ginkgo/v2 to v2.13.2 ([#138](https://github.com/validator-labs/validator/issues/138)) ([31746d6](https://github.com/validator-labs/validator/commit/31746d68c3f43ac3be1c0abacb62a35c57a7f1ce))


### Other

* bump AWS chart version ([c569524](https://github.com/validator-labs/validator/commit/c569524ad2cb0e7ca2995c1fde327abe82c61425))
* update network plugin values ([#135](https://github.com/validator-labs/validator/issues/135)) ([1049435](https://github.com/validator-labs/validator/commit/104943556a73e3aa4e5acd80c99542222035a867))

## [0.0.26](https://github.com/validator-labs/validator/compare/v0.0.25...v0.0.26) (2023-11-28)


### Features

* add Azure plugin ([#131](https://github.com/validator-labs/validator/issues/131)) ([25073e3](https://github.com/validator-labs/validator/commit/25073e3c6fdc6b2556d3f0d59192b1b6ec65281b))
* support private helm repos ([#132](https://github.com/validator-labs/validator/issues/132)) ([cb0cf32](https://github.com/validator-labs/validator/commit/cb0cf32e1c8b09bdaa791c1933f36321f1687046))


### Bug Fixes

* **deps:** update golang.org/x/exp digest to 6522937 ([#133](https://github.com/validator-labs/validator/issues/133)) ([531c8ca](https://github.com/validator-labs/validator/commit/531c8ca2647164ac99777601c95ae469cb5c67c3))


### Other

* bump azure chart version ([9230953](https://github.com/validator-labs/validator/commit/9230953434f15ef1c8bc19658d8bc8e9156df74a))
* **deps:** update anchore/sbom-action action to v0.15.0 ([#129](https://github.com/validator-labs/validator/issues/129)) ([961377b](https://github.com/validator-labs/validator/commit/961377b591c7c071f85046a8c5211ee6e161a38d))
* update AWS auth config ([#134](https://github.com/validator-labs/validator/issues/134)) ([9513e43](https://github.com/validator-labs/validator/commit/9513e43a9860cc1ac2f1fcea01e8d7727d81b11c))


### Docs

* refer to secret templates in values.yaml ([eeae1e7](https://github.com/validator-labs/validator/commit/eeae1e710a2dd584ea74e8017ddb359a165c9079))

## [0.0.25](https://github.com/validator-labs/validator/compare/v0.0.24...v0.0.25) (2023-11-17)


### Bug Fixes

* remove redundant b64 in vsphere secret ([f7f0555](https://github.com/validator-labs/validator/commit/f7f0555f9ad87a6726406c0b39c172e545ad9067))

## [0.0.24](https://github.com/validator-labs/validator/compare/v0.0.23...v0.0.24) (2023-11-17)


### Bug Fixes

* omit secret data from ValidatorConfig ([#125](https://github.com/validator-labs/validator/issues/125)) ([e96d4fe](https://github.com/validator-labs/validator/commit/e96d4fe3cee5bb8791ea95dcdae471f111255798))

## [0.0.23](https://github.com/validator-labs/validator/compare/v0.0.22...v0.0.23) (2023-11-17)


### Bug Fixes

* quote all optional fields in sink secret ([e0a1365](https://github.com/validator-labs/validator/commit/e0a1365d577f329c8dc4d0bca8f6d3eb25e0c9c5))


### Other

* **deps:** update docker/build-push-action digest to 4a13e50 ([#119](https://github.com/validator-labs/validator/issues/119)) ([24b3bfc](https://github.com/validator-labs/validator/commit/24b3bfc927dc5d04fa77b58157bda2f2f18fcb12))

## [0.0.22](https://github.com/validator-labs/validator/compare/v0.0.21...v0.0.22) (2023-11-17)


### Bug Fixes

* format alertmanager cert properly ([#120](https://github.com/validator-labs/validator/issues/120)) ([f38635f](https://github.com/validator-labs/validator/commit/f38635f4de63a66f645bca3c9d6a239695a1ab2a))

## [0.0.21](https://github.com/validator-labs/validator/compare/v0.0.20...v0.0.21) (2023-11-17)


### Features

* expose insecureSkipVerify for Alertmanager sink ([#118](https://github.com/validator-labs/validator/issues/118)) ([8dc7548](https://github.com/validator-labs/validator/commit/8dc7548ed9511abbd382ec2881b473f0b42fd607))


### Refactoring

* expose sink types ([#117](https://github.com/validator-labs/validator/issues/117)) ([f28d8af](https://github.com/validator-labs/validator/commit/f28d8afc5092998189da4112e8a41febdadf1c96))
* standardize get CR in Reconcile ([9fbfff0](https://github.com/validator-labs/validator/commit/9fbfff0c059debab0c9c6044c360af07c8cd4382))

## [0.0.20](https://github.com/validator-labs/validator/compare/v0.0.19...v0.0.20) (2023-11-16)


### Bug Fixes

* retry all status updates due to controller contention ([#114](https://github.com/validator-labs/validator/issues/114)) ([35f03a4](https://github.com/validator-labs/validator/commit/35f03a407a3d0bbcfd76c749908e4b1c9581afac))

## [0.0.19](https://github.com/validator-labs/validator/compare/v0.0.18...v0.0.19) (2023-11-16)


### Features

* add alertmanager sink ([#107](https://github.com/validator-labs/validator/issues/107)) ([855e70e](https://github.com/validator-labs/validator/commit/855e70e69c67cd338f83add9b0b18026e3395184))


### Bug Fixes

* **deps:** update kubernetes packages to v0.28.4 ([#112](https://github.com/validator-labs/validator/issues/112)) ([fc10444](https://github.com/validator-labs/validator/commit/fc104445fab89a663ff0e3fee8ea500b1d0a0904))
* ensure plugin removal during Helm uninstall ([#111](https://github.com/validator-labs/validator/issues/111)) ([0917418](https://github.com/validator-labs/validator/commit/0917418b6ae3f2940bf8048c0cb09ca4056f21da))


### Docs

* issue template addition ([#109](https://github.com/validator-labs/validator/issues/109)) ([36ce4a1](https://github.com/validator-labs/validator/commit/36ce4a1d5630c22b39d481bc45641c5c06e6db04))


### Refactoring

* accept VR in HandleNewValidationResult for flexibility in plugins ([#113](https://github.com/validator-labs/validator/issues/113)) ([1ead151](https://github.com/validator-labs/validator/commit/1ead15146156ac278aedb2a77cab0604488fda4f))

## [0.0.18](https://github.com/validator-labs/validator/compare/v0.0.17...v0.0.18) (2023-11-12)


### Bug Fixes

* **deps:** update module github.com/onsi/ginkgo/v2 to v2.13.1 ([#95](https://github.com/validator-labs/validator/issues/95)) ([496ecad](https://github.com/validator-labs/validator/commit/496ecada5655f5760e46f7d647ce381f616ad56f))
* **deps:** update module sigs.k8s.io/yaml to v1.4.0 ([#98](https://github.com/validator-labs/validator/issues/98)) ([5f35bba](https://github.com/validator-labs/validator/commit/5f35bbac77502a944d6d5641e1e2f88f98cf7c79))
* retry VR status updates ([21b3808](https://github.com/validator-labs/validator/commit/21b3808f36a621f89ddc22aa5362d4d7b47265b5))
* SafeUpdateValidationResult not handling all edge cases ([#104](https://github.com/validator-labs/validator/issues/104)) ([8f34e2f](https://github.com/validator-labs/validator/commit/8f34e2f677a2b70c3c931491ce8b5cd6ac7abd0b))


### Other

* **deps:** pin codecov/codecov-action action to eaaf4be ([#105](https://github.com/validator-labs/validator/issues/105)) ([70c3a0d](https://github.com/validator-labs/validator/commit/70c3a0d834cccc0bc6400c12395b7952c3fb25bb))
* **deps:** pin ncruces/go-coverage-report action to a12281c ([#101](https://github.com/validator-labs/validator/issues/101)) ([bb5d6ac](https://github.com/validator-labs/validator/commit/bb5d6ac4b235013eeba198e2ba4fd86c5eeb93a2))

## [0.0.17](https://github.com/validator-labs/validator/compare/v0.0.16...v0.0.17) (2023-11-10)


### Bug Fixes

* ensure ValidationResult Status resets to successful if all checks pass ([10237c5](https://github.com/validator-labs/validator/commit/10237c5a17adbc61288c6c6b9b981b072ea0e46f))


### Other

* add coverage report ([#100](https://github.com/validator-labs/validator/issues/100)) ([5d24704](https://github.com/validator-labs/validator/commit/5d2470488e1a09ca0d0f623762ae926dea5ba3ef))

## [0.0.16](https://github.com/validator-labs/validator/compare/v0.0.15...v0.0.16) (2023-11-10)


### Bug Fixes

* **deps:** update golang.org/x/exp digest to 9a3e603 ([#97](https://github.com/validator-labs/validator/issues/97)) ([17c83e2](https://github.com/validator-labs/validator/commit/17c83e2f9bc4580eb8b2d9b1f9b4767a42a32d2f))
* ensure State always remains failed if any conditions fail ([4f55823](https://github.com/validator-labs/validator/commit/4f558234af6c190186f8d82fb4a6e135d83030e8))
* ensure State always remains failed if any conditions fail ([711485f](https://github.com/validator-labs/validator/commit/711485f130804eab34de871f86c138adcb3bf257))


### Other

* fix Helm chart lint error ([c938c89](https://github.com/validator-labs/validator/commit/c938c89947371fa3e3e3e2d43044749a413ec443))
* release 0.0.16 ([f031904](https://github.com/validator-labs/validator/commit/f031904a18fb5854586d58145842689d12028360))

## [0.0.15](https://github.com/validator-labs/validator/compare/v0.0.14...v0.0.15) (2023-11-10)


### Bug Fixes

* ensure State always remains failed if any conditions fail ([67e8462](https://github.com/validator-labs/validator/commit/67e846293ce26e8f416fbb24d4b247d38f2a15aa))


### Other

* Merge branch 'main' of https://github.com/validator-labs/validator ([42cf0ef](https://github.com/validator-labs/validator/commit/42cf0efc8ee8f8187b260ac0f4d1f8b1b9db6c16))
* release 0.0.15 ([d574854](https://github.com/validator-labs/validator/commit/d57485496dc90f4b6e421d978bdf0451edfcc59c))

## [0.0.14](https://github.com/validator-labs/validator/compare/v0.0.13...v0.0.14) (2023-11-10)


### Bug Fixes

* prevent extraneous sink emissions ([22de119](https://github.com/validator-labs/validator/commit/22de119a3ce93c8fb751473f5250d55446012d83))


### Other

* Merge branch 'main' of https://github.com/validator-labs/validator ([a2700a2](https://github.com/validator-labs/validator/commit/a2700a26661d6c7668965cc5b5a377adfd49d109))


### Refactoring

* use header block for msg titles ([a6ce7af](https://github.com/validator-labs/validator/commit/a6ce7af748dcb6cbd50c5d13e1a2c9a724e9a328))

## [0.0.13](https://github.com/validator-labs/validator/compare/v0.0.11...v0.0.13) (2023-11-10)


### Bug Fixes

* handle VRs w/ multiple conditions ([8a3a243](https://github.com/validator-labs/validator/commit/8a3a2431dbd317dce80376c64d336317c688c2a3))
* omit sink state from hash to avoid double-emitting on VR creation ([071b960](https://github.com/validator-labs/validator/commit/071b9602fee35262a66a9f403ceb878ac4a422b4))


### Other

* release 0.0.12 ([4903786](https://github.com/validator-labs/validator/commit/49037866402a7f16ef2c15cf172a11626392b9ff))
* release 0.0.13 ([1233488](https://github.com/validator-labs/validator/commit/1233488756f02cd3de9ee2a0d800cb29522545c8))
* Update default values.yaml ([f9af283](https://github.com/validator-labs/validator/commit/f9af2832bd652c9ea404d08936a6879930d29753))


### Refactoring

* change devspace port; always use exponential backoff; tidy validation result helpers ([c866429](https://github.com/validator-labs/validator/commit/c866429772e8d301916eed7ea8908b811cad3e7c))

## [0.0.11](https://github.com/validator-labs/validator/compare/v0.0.10...v0.0.11) (2023-11-09)


### Features

* add sink interface, Slack sink ([#84](https://github.com/validator-labs/validator/issues/84)) ([dac2c3a](https://github.com/validator-labs/validator/commit/dac2c3a83eebbe433790aa439cab1433eadfb0ec))


### Bug Fixes

* **deps:** update golang.org/x/exp digest to 2478ac8 ([#83](https://github.com/validator-labs/validator/issues/83)) ([0d5200f](https://github.com/validator-labs/validator/commit/0d5200faf789343c10149d5f1333894a51d13ff8))
* **deps:** update module github.com/go-logr/logr to v1.3.0 ([#77](https://github.com/validator-labs/validator/issues/77)) ([532fd6f](https://github.com/validator-labs/validator/commit/532fd6f82cf6a9f8322a74428e0834bb023ff67e))
* **deps:** update module github.com/onsi/gomega to v1.28.1 ([#74](https://github.com/validator-labs/validator/issues/74)) ([eb7d990](https://github.com/validator-labs/validator/commit/eb7d990a840d14c5700dffc549de8f27feb0b614))
* **deps:** update module github.com/onsi/gomega to v1.29.0 ([#76](https://github.com/validator-labs/validator/issues/76)) ([ca58e8c](https://github.com/validator-labs/validator/commit/ca58e8c622c75ce345550e2115f8311428bfceb5))
* **deps:** update module github.com/onsi/gomega to v1.30.0 ([#82](https://github.com/validator-labs/validator/issues/82)) ([7bfa8f7](https://github.com/validator-labs/validator/commit/7bfa8f71aa77db2953aaba698f6acf4f2700a03e))
* **deps:** update module k8s.io/klog/v2 to v2.110.1 ([#78](https://github.com/validator-labs/validator/issues/78)) ([8a79219](https://github.com/validator-labs/validator/commit/8a79219a40a4beb47182da4f5aea0d8045dd119f))
* update plugin versions in default values.yaml ([6f7f90d](https://github.com/validator-labs/validator/commit/6f7f90d15fbeb105df6c5b6c568c35fa4f12121f))


### Other

* add license ([065ef97](https://github.com/validator-labs/validator/commit/065ef97d16dadf35a54b84fe3cd1026e14f172d5))
* **deps:** update gcr.io/kubebuilder/kube-rbac-proxy docker tag to v0.15.0 ([#72](https://github.com/validator-labs/validator/issues/72)) ([4f0fc63](https://github.com/validator-labs/validator/commit/4f0fc630bc3ab969d6411fb1d31d968e313a20be))
* **deps:** update google-github-actions/release-please-action digest to db8f2c6 ([#81](https://github.com/validator-labs/validator/issues/81)) ([20956a3](https://github.com/validator-labs/validator/commit/20956a3fa864d5cc0e37349ba0632da61879b8b8))
* **deps:** update helm/chart-testing-action action to v2.6.0 ([#79](https://github.com/validator-labs/validator/issues/79)) ([3272b22](https://github.com/validator-labs/validator/commit/3272b226df2586344efd1dcf3f077483ca3f45a7))
* **deps:** update helm/chart-testing-action action to v2.6.1 ([#80](https://github.com/validator-labs/validator/issues/80)) ([cbb430e](https://github.com/validator-labs/validator/commit/cbb430e70a8aad6558816d3ce2c5c42cacefcab8))
* **main:** release 0.0.10 ([#70](https://github.com/validator-labs/validator/issues/70)) ([6c5e4fc](https://github.com/validator-labs/validator/commit/6c5e4fcc1182862e3902044e543309800e96b31e))
* **main:** release 0.0.10 ([#75](https://github.com/validator-labs/validator/issues/75)) ([32e4025](https://github.com/validator-labs/validator/commit/32e4025ba21223c7214e5378bb060769b931e685))
* release 0.0.10 ([65ce207](https://github.com/validator-labs/validator/commit/65ce2076727fd3d1da49afd884329c55a6394c91))
* release 0.0.11 ([f694577](https://github.com/validator-labs/validator/commit/f694577bb2b0fae8127935c3f9fd9e90f14fe328))

## [0.0.10](https://github.com/validator-labs/validator/compare/v0.0.10...v0.0.10) (2023-11-09)


### Features

* add sink interface, Slack sink ([#84](https://github.com/validator-labs/validator/issues/84)) ([dac2c3a](https://github.com/validator-labs/validator/commit/dac2c3a83eebbe433790aa439cab1433eadfb0ec))


### Bug Fixes

* **deps:** update golang.org/x/exp digest to 2478ac8 ([#83](https://github.com/validator-labs/validator/issues/83)) ([0d5200f](https://github.com/validator-labs/validator/commit/0d5200faf789343c10149d5f1333894a51d13ff8))
* **deps:** update module github.com/go-logr/logr to v1.3.0 ([#77](https://github.com/validator-labs/validator/issues/77)) ([532fd6f](https://github.com/validator-labs/validator/commit/532fd6f82cf6a9f8322a74428e0834bb023ff67e))
* **deps:** update module github.com/onsi/gomega to v1.28.1 ([#74](https://github.com/validator-labs/validator/issues/74)) ([eb7d990](https://github.com/validator-labs/validator/commit/eb7d990a840d14c5700dffc549de8f27feb0b614))
* **deps:** update module github.com/onsi/gomega to v1.29.0 ([#76](https://github.com/validator-labs/validator/issues/76)) ([ca58e8c](https://github.com/validator-labs/validator/commit/ca58e8c622c75ce345550e2115f8311428bfceb5))
* **deps:** update module github.com/onsi/gomega to v1.30.0 ([#82](https://github.com/validator-labs/validator/issues/82)) ([7bfa8f7](https://github.com/validator-labs/validator/commit/7bfa8f71aa77db2953aaba698f6acf4f2700a03e))
* **deps:** update module k8s.io/klog/v2 to v2.110.1 ([#78](https://github.com/validator-labs/validator/issues/78)) ([8a79219](https://github.com/validator-labs/validator/commit/8a79219a40a4beb47182da4f5aea0d8045dd119f))
* update plugin versions in default values.yaml ([6f7f90d](https://github.com/validator-labs/validator/commit/6f7f90d15fbeb105df6c5b6c568c35fa4f12121f))


### Other

* add license ([065ef97](https://github.com/validator-labs/validator/commit/065ef97d16dadf35a54b84fe3cd1026e14f172d5))
* **deps:** update gcr.io/kubebuilder/kube-rbac-proxy docker tag to v0.15.0 ([#72](https://github.com/validator-labs/validator/issues/72)) ([4f0fc63](https://github.com/validator-labs/validator/commit/4f0fc630bc3ab969d6411fb1d31d968e313a20be))
* **deps:** update google-github-actions/release-please-action digest to db8f2c6 ([#81](https://github.com/validator-labs/validator/issues/81)) ([20956a3](https://github.com/validator-labs/validator/commit/20956a3fa864d5cc0e37349ba0632da61879b8b8))
* **deps:** update helm/chart-testing-action action to v2.6.0 ([#79](https://github.com/validator-labs/validator/issues/79)) ([3272b22](https://github.com/validator-labs/validator/commit/3272b226df2586344efd1dcf3f077483ca3f45a7))
* **deps:** update helm/chart-testing-action action to v2.6.1 ([#80](https://github.com/validator-labs/validator/issues/80)) ([cbb430e](https://github.com/validator-labs/validator/commit/cbb430e70a8aad6558816d3ce2c5c42cacefcab8))
* **main:** release 0.0.10 ([#70](https://github.com/validator-labs/validator/issues/70)) ([6c5e4fc](https://github.com/validator-labs/validator/commit/6c5e4fcc1182862e3902044e543309800e96b31e))
* release 0.0.10 ([65ce207](https://github.com/validator-labs/validator/commit/65ce2076727fd3d1da49afd884329c55a6394c91))

## [0.0.10](https://github.com/validator-labs/validator/compare/v0.0.10...v0.0.10) (2023-10-20)


### Bug Fixes

* update plugin versions in default values.yaml ([6f7f90d](https://github.com/validator-labs/validator/commit/6f7f90d15fbeb105df6c5b6c568c35fa4f12121f))


### Other

* release 0.0.10 ([65ce207](https://github.com/validator-labs/validator/commit/65ce2076727fd3d1da49afd884329c55a6394c91))

## [0.0.10](https://github.com/validator-labs/validator/compare/v0.0.9...v0.0.10) (2023-10-20)


### Bug Fixes

* **deps:** update module sigs.k8s.io/controller-runtime to v0.16.3 ([#63](https://github.com/validator-labs/validator/issues/63)) ([6f79f8a](https://github.com/validator-labs/validator/commit/6f79f8af8f7a10c70ce403fadeb16d6eb9e13020))


### Other

* **deps:** bump golang.org/x/net from 0.16.0 to 0.17.0 ([#61](https://github.com/validator-labs/validator/issues/61)) ([eca7abd](https://github.com/validator-labs/validator/commit/eca7abd8da204cabb28d7fb6ee8c381d1cb60e7b))
* **deps:** update actions/checkout digest to b4ffde6 ([#64](https://github.com/validator-labs/validator/issues/64)) ([d9bbd21](https://github.com/validator-labs/validator/commit/d9bbd21fe962de4f14e0e734d697ebd2ceb7e144))
* **deps:** update actions/setup-python digest to 65d7f2d ([#65](https://github.com/validator-labs/validator/issues/65)) ([bdb95d0](https://github.com/validator-labs/validator/commit/bdb95d038149ed9eab6c8372018bb771b848157c))
* **deps:** update gcr.io/kubebuilder/kube-rbac-proxy docker tag to v0.14.4 ([#59](https://github.com/validator-labs/validator/issues/59)) ([78847f5](https://github.com/validator-labs/validator/commit/78847f54c35f6fc684d333a286e1315cc558e6e2))
* **deps:** update google-github-actions/release-please-action digest to 4c5670f ([#66](https://github.com/validator-labs/validator/issues/66)) ([2c24d48](https://github.com/validator-labs/validator/commit/2c24d48035b7ebddbbd20ca182e3352fa6c2f09e))
* enable renovate automerges ([84ad3cd](https://github.com/validator-labs/validator/commit/84ad3cdee59ed634e5f9577d801dc13701958e46))
* release 0.0.10 ([9a324e9](https://github.com/validator-labs/validator/commit/9a324e9e666b0da68a1e0c3be225ab19bfd04a6b))


### Refactoring

* valid8or -&gt; validator ([#67](https://github.com/validator-labs/validator/issues/67)) ([ff88026](https://github.com/validator-labs/validator/commit/ff8802656d8115dd6afbbfbaf12613c4f205feb5))

## [0.0.9](https://github.com/validator-labs/validator/compare/v0.0.8...v0.0.9) (2023-10-10)


### Bug Fixes

* **deps:** update golang.org/x/exp digest to 7918f67 ([#55](https://github.com/validator-labs/validator/issues/55)) ([3f173d4](https://github.com/validator-labs/validator/commit/3f173d4dc256415d9f447133afc70024d0115021))
* **deps:** update kubernetes packages to v0.28.2 ([#51](https://github.com/validator-labs/validator/issues/51)) ([f43d5a0](https://github.com/validator-labs/validator/commit/f43d5a098d6d5923fa540564279defe31701f3c7))
* **deps:** update module github.com/onsi/ginkgo/v2 to v2.12.1 ([#52](https://github.com/validator-labs/validator/issues/52)) ([516693e](https://github.com/validator-labs/validator/commit/516693e3d150dfceb0a6be3a5f00bfbe260a1cb6))
* **deps:** update module github.com/onsi/ginkgo/v2 to v2.13.0 ([#57](https://github.com/validator-labs/validator/issues/57)) ([be32cb8](https://github.com/validator-labs/validator/commit/be32cb85ea38e8ddfd03e9d6837dddda1967c4b7))
* **deps:** update module github.com/onsi/gomega to v1.28.0 ([#54](https://github.com/validator-labs/validator/issues/54)) ([e89431e](https://github.com/validator-labs/validator/commit/e89431e5685fc9ab83dd7dd2a87864f57b835bcb))
* **deps:** update module sigs.k8s.io/controller-runtime to v0.16.2 ([#50](https://github.com/validator-labs/validator/issues/50)) ([1a07c2a](https://github.com/validator-labs/validator/commit/1a07c2a134bc26d86a17c04fcf71bae2f601c3e5))


### Other

* **deps:** update actions/checkout digest to 8ade135 ([#53](https://github.com/validator-labs/validator/issues/53)) ([d42a8a6](https://github.com/validator-labs/validator/commit/d42a8a6832c15b76eac882a7f35fc3a2f8ba842a))
* **deps:** update docker/build-push-action action to v5 ([#47](https://github.com/validator-labs/validator/issues/47)) ([40d0053](https://github.com/validator-labs/validator/commit/40d0053011c9d05582d96ab07c7e743867fbed8d))
* **deps:** update docker/build-push-action digest to 0a97817 ([#45](https://github.com/validator-labs/validator/issues/45)) ([91a951b](https://github.com/validator-labs/validator/commit/91a951b03d40db29020af68cb7b1a22db8cad19c))
* **deps:** update docker/login-action action to v3 ([#48](https://github.com/validator-labs/validator/issues/48)) ([9c940d7](https://github.com/validator-labs/validator/commit/9c940d7acebf0e237ed5129cc3b25aca596a257f))
* **deps:** update docker/setup-buildx-action action to v3 ([#49](https://github.com/validator-labs/validator/issues/49)) ([a3fe730](https://github.com/validator-labs/validator/commit/a3fe730ae38b9655583be333e255e48767c2cf0c))
* release 0.0.9 ([3cbcfa9](https://github.com/validator-labs/validator/commit/3cbcfa9e3f1e5c75dcfc083df1494aab0c41c31f))

## [0.0.8](https://github.com/validator-labs/validator/compare/v0.0.7...v0.0.8) (2023-09-06)


### Other

* bump aws plugin version ([bac567e](https://github.com/validator-labs/validator/commit/bac567e69ac6a957b78f29b5bcec81d2f3b0c6b8))
* **deps:** update actions/upload-artifact digest to a8a3f3a ([#43](https://github.com/validator-labs/validator/issues/43)) ([fc33646](https://github.com/validator-labs/validator/commit/fc3364674e45c9dd3fdbf15773a7fbe2f04e3395))

## [0.0.7](https://github.com/validator-labs/validator/compare/v0.0.6...v0.0.7) (2023-09-06)


### Bug Fixes

* **deps:** update golang.org/x/exp digest to 9212866 ([#41](https://github.com/validator-labs/validator/issues/41)) ([50ad9cb](https://github.com/validator-labs/validator/commit/50ad9cbd72f531ab3e29eb43d59f75026efc96b0))
* include pkg/ in dockerfiles ([b45bb22](https://github.com/validator-labs/validator/commit/b45bb22532adc3e860aea56edf7ab3f3a95519fd))


### Other

* bump plugin versions ([6658d19](https://github.com/validator-labs/validator/commit/6658d190022815cf56d1a345ae66d46fd260c10c))
* **deps:** update actions/checkout action to v4 ([#39](https://github.com/validator-labs/validator/issues/39)) ([3c848b6](https://github.com/validator-labs/validator/commit/3c848b61c3294cdd1ceed376c3e4a48521221b6f))
* update AWS plugin version; fix default values.yaml ([516213b](https://github.com/validator-labs/validator/commit/516213ba35d80c4e8692b1448b22f6e1b9756c5d))


### Refactoring

* expose helm package ([2a34e0a](https://github.com/validator-labs/validator/commit/2a34e0ae780b287a8a1a48ee00f5016e667c8304))

## [0.0.6](https://github.com/validator-labs/validator/compare/v0.0.5...v0.0.6) (2023-09-01)


### Bug Fixes

* **deps:** update kubernetes packages to v0.28.1 ([#30](https://github.com/validator-labs/validator/issues/30)) ([f94b40d](https://github.com/validator-labs/validator/commit/f94b40d7d9b0be097cad185b0426727380cef822))
* **deps:** update module github.com/onsi/ginkgo/v2 to v2.12.0 ([#31](https://github.com/validator-labs/validator/issues/31)) ([98a7aa7](https://github.com/validator-labs/validator/commit/98a7aa785946534db076a093c0715ac63782d72f))
* **deps:** update module github.com/onsi/gomega to v1.27.10 ([#29](https://github.com/validator-labs/validator/issues/29)) ([8697124](https://github.com/validator-labs/validator/commit/8697124f495004f71ed9ba694ab8116880a4ae7f))
* **deps:** update module sigs.k8s.io/controller-runtime to v0.16.1 ([#33](https://github.com/validator-labs/validator/issues/33)) ([94bf0ad](https://github.com/validator-labs/validator/commit/94bf0ad33b8524a82427b12320ba70493af9ac21))
* NET_ADMIN -&gt; NET_RAW ([122cc80](https://github.com/validator-labs/validator/commit/122cc808ed6b83eb6a33dc38e031d86805440784))


### Other

* **deps:** update actions/checkout digest to f43a0e5 ([#25](https://github.com/validator-labs/validator/issues/25)) ([fa0b3d9](https://github.com/validator-labs/validator/commit/fa0b3d95f74e70226f07114bffcde2e1b270ad33))
* **deps:** update actions/setup-go digest to 93397be ([#26](https://github.com/validator-labs/validator/issues/26)) ([e32d52b](https://github.com/validator-labs/validator/commit/e32d52bc1243451b6a0f6a27f228869da9497761))
* **deps:** update docker/setup-buildx-action digest to 885d146 ([#28](https://github.com/validator-labs/validator/issues/28)) ([f4b1dd1](https://github.com/validator-labs/validator/commit/f4b1dd10a4170f2596bf2cebf64a363d21efcc44))
* release 0.0.6 ([10421d5](https://github.com/validator-labs/validator/commit/10421d59a5fbc4dd840a6aac9f24657ee8d1be7c))

## [0.0.5](https://github.com/validator-labs/validator/compare/v0.0.4...v0.0.5) (2023-08-31)


### Bug Fixes

* omit conditions for uninstalled plugins ([c9f430d](https://github.com/validator-labs/validator/commit/c9f430d81bfdbb077edfa5d3cc48f314bf831c45))
* preserve VC annotations when updating plugin hashes ([19c9463](https://github.com/validator-labs/validator/commit/19c9463a4ed7b516731fdaf76cf487e682a6a2c4))
* securityContext blocking MTU check w/ ping ([131e5d9](https://github.com/validator-labs/validator/commit/131e5d91015b54b470b61708bfe8675f7eb26a0e))
* update 2+ plugin conditions properly ([a12488f](https://github.com/validator-labs/validator/commit/a12488f9dc376d1b0a8b413791b3aa2e25b185a5))


### Other

* release 0.0.5 ([24e9712](https://github.com/validator-labs/validator/commit/24e9712ffbf6fb8d33333d5d8c063f3935c7ceae))
* update README ([f8254d4](https://github.com/validator-labs/validator/commit/f8254d4a0bca6db90523ace43503cf6d80d3af30))


### Docs

* add validator-plugin-network to default values.yaml ([1aada24](https://github.com/validator-labs/validator/commit/1aada2492d518e7ed50934d2cb6e184da5d1e031))

## [0.0.4](https://github.com/validator-labs/validator/compare/v0.0.3...v0.0.4) (2023-08-29)


### Other

* add public validation result utils ([528be5f](https://github.com/validator-labs/validator/commit/528be5f91e8bfb7d6f1530002dd99971e4983a7e))
* release 0.0.4 ([28f8418](https://github.com/validator-labs/validator/commit/28f8418356694ad5c86b97d2b9df9d51c2f6d279))


### Docs

* update chart description ([0a5635f](https://github.com/validator-labs/validator/commit/0a5635f94f949c39592dbe7f20e5301f4836f291))

## [0.0.3](https://github.com/validator-labs/validator/compare/v0.0.2...v0.0.3) (2023-08-29)


### Features

* handle plugin updates via values hashes ([7f485b4](https://github.com/validator-labs/validator/commit/7f485b41f5dfba40e8c08a5da79410dfc0c97e0c))
* log ValidationResult metadata on completion ([0cc38e5](https://github.com/validator-labs/validator/commit/0cc38e5cf464d6f9342865f3a41787dfe9bc3c5c))
* plugin management w/ helm ([537faac](https://github.com/validator-labs/validator/commit/537faac4c3f1c6695f1db34114401a14ad292906))
* update status and handle plugin removal ([bae7e9d](https://github.com/validator-labs/validator/commit/bae7e9dc36a1a22e8f08828421d0cc7e73deb54f))


### Bug Fixes

* increase memory limit for helm upgrade ([660a80d](https://github.com/validator-labs/validator/commit/660a80d57fcae2dc3a16e610699b60a5749e0786))
* update RBAC in helm templates ([6ff735c](https://github.com/validator-labs/validator/commit/6ff735c891e40328bba0524f4b8b240d3e85a6c9))


### Other

* add pull_request test trigger ([3e75bdb](https://github.com/validator-labs/validator/commit/3e75bdbff56bab925ca98b81c57fe9f4b1a60471))
* release 0.0.3 ([5b2473d](https://github.com/validator-labs/validator/commit/5b2473dce974a160b14640e86e88028f66c94f5e))


### Docs

* update README, fix release-please annotations ([c2c96e8](https://github.com/validator-labs/validator/commit/c2c96e8e3e91820826242b36d6760ab1d2530baf))

## [0.0.2](https://github.com/validator-labs/validator/compare/v0.0.1...v0.0.2) (2023-08-25)


### Other

* release 0.0.2 ([40cdd88](https://github.com/validator-labs/validator/commit/40cdd88ebb8b75f9908c5dab6aa29337f5d778d8))

## [0.0.1](https://github.com/validator-labs/validator/compare/v0.0.1...v0.0.1) (2023-08-25)


### Bug Fixes

* helm chart CI ([46f37f0](https://github.com/validator-labs/validator/commit/46f37f0cea87e90e6effb85cb15128ab5970a621))
* release image push repo ([4a2aca6](https://github.com/validator-labs/validator/commit/4a2aca6ecbfeca48ed4dd7566441923815281432))


### Other

* release 0.0.1 ([a23551a](https://github.com/validator-labs/validator/commit/a23551a1984d43d9acbc7de3cacad6ee928cc517))
