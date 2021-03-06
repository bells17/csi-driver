# Change Log

## [v0.3.0](https://github.com/kubevault/csi-driver/tree/v0.3.0) (2020-01-10)
[Full Changelog](https://github.com/kubevault/csi-driver/compare/0.1.0...v0.3.0)

**Merged pull requests:**

- Stop running e2e tests in CI, because we have none [\#74](https://github.com/kubevault/csi-driver/pull/74) ([tamalsaha](https://github.com/tamalsaha))
- Vendor kubevault/operator@v0.3.0 [\#73](https://github.com/kubevault/csi-driver/pull/73) ([tamalsaha](https://github.com/tamalsaha))
- Register CRDs from inside before starting the server [\#72](https://github.com/kubevault/csi-driver/pull/72) ([kamolhasan](https://github.com/kamolhasan))
- Fix ci workflow [\#71](https://github.com/kubevault/csi-driver/pull/71) ([tamalsaha](https://github.com/tamalsaha))
- Use pod's service account reference while performing Vault authN [\#70](https://github.com/kubevault/csi-driver/pull/70) ([kamolhasan](https://github.com/kamolhasan))
- Refactor health probe init function [\#69](https://github.com/kubevault/csi-driver/pull/69) ([tamalsaha](https://github.com/tamalsaha))
- Avoid error check for `r.Body.Close\(\)` [\#68](https://github.com/kubevault/csi-driver/pull/68) ([kamolhasan](https://github.com/kamolhasan))
- Fix make install and remove make ct [\#67](https://github.com/kubevault/csi-driver/pull/67) ([tamalsaha](https://github.com/tamalsaha))
- Use function literal while calling goroutine [\#66](https://github.com/kubevault/csi-driver/pull/66) ([kamolhasan](https://github.com/kamolhasan))
- Correct pass imagePullSecrets [\#65](https://github.com/kubevault/csi-driver/pull/65) ([tamalsaha](https://github.com/tamalsaha))
- Fix health check [\#64](https://github.com/kubevault/csi-driver/pull/64) ([tamalsaha](https://github.com/tamalsaha))
- Move installers to installer repo [\#63](https://github.com/kubevault/csi-driver/pull/63) ([tamalsaha](https://github.com/tamalsaha))
- Update client-go to kubernetes-1.16.3 [\#61](https://github.com/kubevault/csi-driver/pull/61) ([tamalsaha](https://github.com/tamalsaha))
- Update AppBinding crds [\#60](https://github.com/kubevault/csi-driver/pull/60) ([tamalsaha](https://github.com/tamalsaha))
- Properly handle empty image pull secret name in installer [\#59](https://github.com/kubevault/csi-driver/pull/59) ([tamalsaha](https://github.com/tamalsaha))
- Test installers [\#58](https://github.com/kubevault/csi-driver/pull/58) ([tamalsaha](https://github.com/tamalsaha))
- Add license header to files [\#57](https://github.com/kubevault/csi-driver/pull/57) ([tamalsaha](https://github.com/tamalsaha))
- Split imports into 3 blocks [\#56](https://github.com/kubevault/csi-driver/pull/56) ([tamalsaha](https://github.com/tamalsaha))
- Enable make ci [\#55](https://github.com/kubevault/csi-driver/pull/55) ([tamalsaha](https://github.com/tamalsaha))
- Update kubevault.dev/operator [\#54](https://github.com/kubevault/csi-driver/pull/54) ([tamalsaha](https://github.com/tamalsaha))
- Add release pipeline [\#53](https://github.com/kubevault/csi-driver/pull/53) ([tamalsaha](https://github.com/tamalsaha))
- Download onessl version v0.13.1 for Kubernetes 1.16 fix [\#52](https://github.com/kubevault/csi-driver/pull/52) ([tamalsaha](https://github.com/tamalsaha))
- Update dependencies [\#51](https://github.com/kubevault/csi-driver/pull/51) ([tamalsaha](https://github.com/tamalsaha))
- Change package path to kubevault.dev/csi-driver [\#50](https://github.com/kubevault/csi-driver/pull/50) ([tamalsaha](https://github.com/tamalsaha))
- Add license header to Makefiles [\#49](https://github.com/kubevault/csi-driver/pull/49) ([tamalsaha](https://github.com/tamalsaha))
- Update sidecar image version [\#48](https://github.com/kubevault/csi-driver/pull/48) ([kamolhasan](https://github.com/kamolhasan))
- Install mount binaries in Docker images [\#47](https://github.com/kubevault/csi-driver/pull/47) ([tamalsaha](https://github.com/tamalsaha))
- Always return IsFormatted = true from Mounter [\#46](https://github.com/kubevault/csi-driver/pull/46) ([tamalsaha](https://github.com/tamalsaha))
- Add make install, uninstall, purge commands [\#45](https://github.com/kubevault/csi-driver/pull/45) ([tamalsaha](https://github.com/tamalsaha))
- Add Makefile [\#44](https://github.com/kubevault/csi-driver/pull/44) ([tamalsaha](https://github.com/tamalsaha))
- Update to CSI 1.1.0 [\#43](https://github.com/kubevault/csi-driver/pull/43) ([sanjid133](https://github.com/sanjid133))
- Update to k8s 1.14.0 client libraries using go.mod [\#42](https://github.com/kubevault/csi-driver/pull/42) ([tamalsaha](https://github.com/tamalsaha))
- Update Kubernetes client libraries to 1.13.5 [\#41](https://github.com/kubevault/csi-driver/pull/41) ([tamalsaha](https://github.com/tamalsaha))

## [0.1.0](https://github.com/kubevault/csi-driver/tree/0.1.0) (2019-03-01)
[Full Changelog](https://github.com/kubevault/csi-driver/compare/0.2.0...0.1.0)

## [0.2.0](https://github.com/kubevault/csi-driver/tree/0.2.0) (2019-03-01)
**Merged pull requests:**

- Remove hostnetwork [\#40](https://github.com/kubevault/csi-driver/pull/40) ([sanjid133](https://github.com/sanjid133))
- Fix certificate [\#39](https://github.com/kubevault/csi-driver/pull/39) ([sanjid133](https://github.com/sanjid133))
- Update Kubernetes client libraries to 1.13.0 [\#38](https://github.com/kubevault/csi-driver/pull/38) ([tamalsaha](https://github.com/tamalsaha))
- Pass pod annotation to deployment [\#37](https://github.com/kubevault/csi-driver/pull/37) ([tamalsaha](https://github.com/tamalsaha))
- Don't use priority class when operator namespace is not kube-system [\#36](https://github.com/kubevault/csi-driver/pull/36) ([tamalsaha](https://github.com/tamalsaha))
- Fix the case for deploying using MINGW64 for windows [\#35](https://github.com/kubevault/csi-driver/pull/35) ([tamalsaha](https://github.com/tamalsaha))
- Use onessl 0.10.0 [\#34](https://github.com/kubevault/csi-driver/pull/34) ([tamalsaha](https://github.com/tamalsaha))
- Add Custom monitoring [\#33](https://github.com/kubevault/csi-driver/pull/33) ([sanjid133](https://github.com/sanjid133))
- Add appbinding user roles [\#32](https://github.com/kubevault/csi-driver/pull/32) ([tamalsaha](https://github.com/tamalsaha))
- Add CSI probe checker to healthz [\#31](https://github.com/kubevault/csi-driver/pull/31) ([tamalsaha](https://github.com/tamalsaha))
- Update csi driver sidecar image tags for 0.3.0 release [\#30](https://github.com/kubevault/csi-driver/pull/30) ([tamalsaha](https://github.com/tamalsaha))
- Fix apiserver cert mounting in installer script [\#29](https://github.com/kubevault/csi-driver/pull/29) ([tamalsaha](https://github.com/tamalsaha))
- Fix service monitor template in chart [\#28](https://github.com/kubevault/csi-driver/pull/28) ([tamalsaha](https://github.com/tamalsaha))
- Only support Kubernetes 1.12.x for release 0.1.0 [\#27](https://github.com/kubevault/csi-driver/pull/27) ([tamalsaha](https://github.com/tamalsaha))
- Add name of components to chart readme [\#26](https://github.com/kubevault/csi-driver/pull/26) ([tamalsaha](https://github.com/tamalsaha))
- Add /metrics and /healthz using k8s.io/apiserver [\#25](https://github.com/kubevault/csi-driver/pull/25) ([tamalsaha](https://github.com/tamalsaha))
- Organize installer yamls [\#24](https://github.com/kubevault/csi-driver/pull/24) ([tamalsaha](https://github.com/tamalsaha))
- Add DCO [\#23](https://github.com/kubevault/csi-driver/pull/23) ([tamalsaha](https://github.com/tamalsaha))
- Reduce log level to --v=5 [\#22](https://github.com/kubevault/csi-driver/pull/22) ([tamalsaha](https://github.com/tamalsaha))
- Update CSI Spec 1.0.0 [\#21](https://github.com/kubevault/csi-driver/pull/21) ([sanjid133](https://github.com/sanjid133))
- Generate reference in docs repository [\#20](https://github.com/kubevault/csi-driver/pull/20) ([tamalsaha](https://github.com/tamalsaha))
- Wait until crd is ready in install.sh [\#19](https://github.com/kubevault/csi-driver/pull/19) ([sanjid133](https://github.com/sanjid133))
- Add csi-vault Chart [\#18](https://github.com/kubevault/csi-driver/pull/18) ([sanjid133](https://github.com/sanjid133))
- Update readme and rename binary to csi-vault [\#17](https://github.com/kubevault/csi-driver/pull/17) ([tamalsaha](https://github.com/tamalsaha))
- Add travis.yml [\#16](https://github.com/kubevault/csi-driver/pull/16) ([tamalsaha](https://github.com/tamalsaha))
- E2E test added [\#15](https://github.com/kubevault/csi-driver/pull/15) ([sanjid133](https://github.com/sanjid133))
-  Enable DB secret support in csi-drive [\#14](https://github.com/kubevault/csi-driver/pull/14) ([sanjid133](https://github.com/sanjid133))
-  Use secret engine from operator repo in csi-driver [\#13](https://github.com/kubevault/csi-driver/pull/13) ([sanjid133](https://github.com/sanjid133))
- Set periodic analytics [\#11](https://github.com/kubevault/csi-driver/pull/11) ([tamalsaha](https://github.com/tamalsaha))
- Update Kubernetes client libraries to 1.12.0 [\#10](https://github.com/kubevault/csi-driver/pull/10) ([tamalsaha](https://github.com/tamalsaha))
- AppBinding cleanup [\#9](https://github.com/kubevault/csi-driver/pull/9) ([sanjid133](https://github.com/sanjid133))
- Introduce AppBinding [\#8](https://github.com/kubevault/csi-driver/pull/8) ([sanjid133](https://github.com/sanjid133))
- Update kubernetes client libraries to 1.12.0 [\#7](https://github.com/kubevault/csi-driver/pull/7) ([tamalsaha](https://github.com/tamalsaha))
- PKI secret engine added [\#6](https://github.com/kubevault/csi-driver/pull/6) ([sanjid133](https://github.com/sanjid133))
- Aws secret engine added [\#4](https://github.com/kubevault/csi-driver/pull/4) ([sanjid133](https://github.com/sanjid133))
- Pod info support  with kubernetes 1.12  [\#3](https://github.com/kubevault/csi-driver/pull/3) ([sanjid133](https://github.com/sanjid133))
- Initial implementation [\#2](https://github.com/kubevault/csi-driver/pull/2) ([sanjid133](https://github.com/sanjid133))
- Add build scripts [\#1](https://github.com/kubevault/csi-driver/pull/1) ([sanjid133](https://github.com/sanjid133))



\* *This Change Log was automatically generated by [github_changelog_generator](https://github.com/skywinder/Github-Changelog-Generator)*