- [v4.1.4-alpha3](#v414-alpha3)
  - [Downloads for v4.1.4-alpha3](#downloads-for-v414-alpha3)
    - [Source Code](#source-code)
    - [Client Binaries](#client-binaries)
  - [Usage](#usage)
  - [Changelog since v4.1.2-alpha11](#changelog-since-v412-alpha11)

# [v4.1.4-alpha3](https://github.com/labring/sealos/releases/tag/v4.1.4-alpha3)

## Downloads for v4.1.4-alpha3

### Source Code

filename |
-------- |
[v4.1.4-alpha3.tar.gz](https://github.com/labring/sealos/archive/refs/tags/v4.1.4-alpha3.tar.gz) |

### Client Binaries

filename |
-------- |
[sealos_4.1.4-alpha3_linux_amd64.tar.gz](https://github.com/labring/sealos/releases/download/v4.1.4-alpha3/sealos_4.1.4-alpha3_linux_amd64.tar.gz) |
[sealos_4.1.4-alpha3_linux_arm64.tar.gz](https://github.com/labring/sealos/releases/download/v4.1.4-alpha3/sealos_4.1.4-alpha3_linux_arm64.tar.gz) |

## Usage

amd64:

```shell
wget  https://github.com/labring/sealos/releases/download/v4.1.4-alpha3/sealos_4.1.4-alpha3_linux_amd64.tar.gz  &&     tar -zxvf sealos_4.1.4-alpha3_linux_amd64.tar.gz sealos &&  chmod +x sealos && mv sealos /usr/bin
## create a cluster for sealos
sealos run labring/kubernetes:v1.24.0 labring/calico:v3.22.1 --masters 192.168.64.2,192.168.64.22,192.168.64.20 --nodes 192.168.64.21,192.168.64.19 --passwd your-own-ssh-passwd
```

arm64:

```shell
wget  https://github.com/labring/sealos/releases/download/v4.1.4-alpha3/sealos_4.1.4-alpha3_linux_arm64.tar.gz  &&     tar -zxvf sealos_4.1.4-alpha3_linux_arm64.tar.gz sealos &&  chmod +x sealos && mv sealos /usr/bin
## create a cluster for sealos
sealos run labring/kubernetes:v1.24.0 labring/calico:v3.22.1 --masters 192.168.64.2,192.168.64.22,192.168.64.20 --nodes 192.168.64.21,192.168.64.19 --passwd your-own-ssh-passwd
```

## Changelog since v4.1.2-alpha11

### What's Changed
* 30426036 - Mercurio - ci: publish docker image for the release workflow with goreleaser (https://github.com/labring/sealos/pull/#1964)
* 21a6607e - Mercurio - fix: goreleaser actions failure (https://github.com/labring/sealos/pull/#1962)
* 5b44b8f3 - 中弈 - add sealos cloud example: hello world demo (https://github.com/labring/sealos/pull/#1960)
* 74d303ea - Harry - docs: update workspace preparation steps (https://github.com/labring/sealos/pull/#1955)
* 7bddc8ca - cuisongliu - feature(main): fix cri socket to multiple CRI (https://github.com/labring/sealos/pull/#1945)
* 58333443 - zzjin - Try fix goreleaser for arm package. (https://github.com/labring/sealos/pull/#1952)
* 159771d1 - cuisongliu - feature(main): add sealos action for docs (https://github.com/labring/sealos/pull/#1949)
* 10a488fc - ining7 - del: (1) (https://github.com/labring/sealos/pull/#1948)
* 38de1a82 - ining7 - fix: error link (https://github.com/labring/sealos/pull/#1947)
* f06b3489 - cuisongliu - feature(main): add GetContextDir for build (https://github.com/labring/sealos/pull/#1935)
* 5117036f - cuisongliu - feature(main): fix nodes for single (https://github.com/labring/sealos/pull/#1938)
* 0601da33 - Bingchang Chen - fix: casdoor non root security (https://github.com/labring/sealos/pull/#1940)
* dfdd8bde - cuisongliu - feature(main): sync ttl for token and kubeadm-certs (https://github.com/labring/sealos/pull/#1936)
* 78fd97f1 - 晓杰 - shell:creat a sealos cluster based on multipass (https://github.com/labring/sealos/pull/#1934)
* 6954674b - Mercurio - fix: get host arch using SSH for apply/run, apply/reset and apply/scale (https://github.com/labring/sealos/pull/#1927)
* 6d990827 - zzjin - Update doc typo. (https://github.com/labring/sealos/pull/#1933)
* f1a1cbdb - cuisongliu - feature(main): add user-controller rbac for sa (https://github.com/labring/sealos/pull/#1926) (https://github.com/labring/sealos/pull/#1928)
* 97067cfe - YAO - [fix]add pull policy (https://github.com/labring/sealos/pull/#1921)
* cfbf051d - cuisongliu - feature(main): add user-controller rbac for sa (https://github.com/labring/sealos/pull/#1926)
* 50bda39e - zzjin - temp comments out L (https://github.com/labring/sealos/pull/#1925)
* 90341162 - YanJun-Zhao - fix readme and intro (https://github.com/labring/sealos/pull/#1918)
* b24684f6 - 中弈 - add contribute guide, add golang install and FAQ (https://github.com/labring/sealos/pull/#1919)
* 4c5dae8e - Mercurio - fix: use unknown user instead of running useradd (https://github.com/labring/sealos/pull/#1913)
* 58009184 - xiongxinwei - desktop/frontend/README.md (https://github.com/labring/sealos/pull/#1912)
* f7f6afeb - TomatoQt - docs:A few typos(Kubernetes first letter, grammar mistake:you->your, Chinese punctuation in English document) (https://github.com/labring/sealos/pull/#1911)
* 9c6d2102 - Mercurio - ci: fix docs site build error (https://github.com/labring/sealos/pull/#1910)
* ce3e2a34 - leezq - fix: alert cause copy fail (https://github.com/labring/sealos/pull/#1908)
* f8a7bbfb - leezq - add copy token (https://github.com/labring/sealos/pull/#1906)
* 20ffd239 - 晓杰 - doc: update k8s image name (https://github.com/labring/sealos/pull/#1907)
* 073d8d47 - zzjin - Fix app import typo (https://github.com/labring/sealos/pull/#1904)
* 9eaf95c8 - leezq - Fix frontend (https://github.com/labring/sealos/pull/#1895)
* 449d465e - cuisongliu - feature(main): add registry sdk (https://github.com/labring/sealos/pull/#1880)
* 86f82bbf - zzjin - Dev front (https://github.com/labring/sealos/pull/#1901)
* 52879453 - Mercurio - docs: improve existing docs (https://github.com/labring/sealos/pull/#1896)
* 24b57b23 - 中弈 - add mock dapps (https://github.com/labring/sealos/pull/#1898)
* f751220c - NTH19 - make SetClusterRunArgs conform to 'return fast' (https://github.com/labring/sealos/pull/#1889)
* e5891757 - zzjin - Add apps (https://github.com/labring/sealos/pull/#1893)
* 954bf3c8 - Bingchang Chen - fix: watch+select to get kubeconfig (https://github.com/labring/sealos/pull/#1892)
* 4c35d397 - cuisongliu - feature(main): add log for user controller (https://github.com/labring/sealos/pull/#1888)
* 89bf5774 - gitccl - feat: deploy terminal in terminal-app ns (https://github.com/labring/sealos/pull/#1835)
* c577eb98 - Mercurio - fix: the rest of semgrep scan issues (https://github.com/labring/sealos/pull/#1886)
* b0e5e206 - Mercurio - fix: semgrep scan issues (https://github.com/labring/sealos/pull/#1881)
* 9a850136 - zzjin - Fix terminal rbac usage (https://github.com/labring/sealos/pull/#1885)
* cfd7a3d0 - Mercurio - docs: update sealos API docs (https://github.com/labring/sealos/pull/#1879)
* 2cdc0be7 - HURUIZHE - Frontend and ssh add pkdata (https://github.com/labring/sealos/pull/#1831)
* b27be7e7 - cuisongliu - feature(main): fix miss rbac (https://github.com/labring/sealos/pull/#1874)
* 08d992d5 - 中弈 - add support kubernetes versions in readme (https://github.com/labring/sealos/pull/#1868)
* 15f098fe - 晓杰 - feat:add metering module (https://github.com/labring/sealos/pull/#1824)
* 0d3a91ff - 晓杰 - feat:payment success and account delete payment delay 5 minutes (https://github.com/labring/sealos/pull/#1863)
* 1c1d86a3 - 晓杰 - fix:sealos creat xxxx image not known (https://github.com/labring/sealos/pull/#1864)
* af773ada - fengxsong - optimize the scale process (https://github.com/labring/sealos/pull/#1855)
* bd99f547 - Mercurio - feat: support pulling in sealos save and loading in sealos run (https://github.com/labring/sealos/pull/#1861)
* 39fd93e1 - 晓杰 - fix:sealos creat xxxx image not known (https://github.com/labring/sealos/pull/#1860)
* 395dc475 - zzjin - Fix appstore typo (https://github.com/labring/sealos/pull/#1856)
* 021f117d - cuisongliu - feature(main): add auto build cluster-image (https://github.com/labring/sealos/pull/#1859)
* e8f92d74 - cuisongliu - feature(main): fix image pull for user-controller (https://github.com/labring/sealos/pull/#1857)
* 4bcfd096 - cuisongliu - feature(main): add sa kubeconfig (https://github.com/labring/sealos/pull/#1836)
* 762e078b - BambooSword - feat: appstore (https://github.com/labring/sealos/pull/#1837)
* 2b0fe153 - zzjin - Update Frontend apps (https://github.com/labring/sealos/pull/#1847)
* 33dc275a - Raving - docs: fix some typo in CONTRIBUTING.md. (https://github.com/labring/sealos/pull/#1852)
* 3d43c7de - 晓杰 - feat:add payment secvet set env (https://github.com/labring/sealos/pull/#1848)
* bc02dce3 - leezq - fix: Hydration failed (https://github.com/labring/sealos/pull/#1840)
* 2fcd9121 - fengxsong - fix: return error if any error occured (https://github.com/labring/sealos/pull/#1844)
* 343e7be2 - fengxsong - fix: copy binary with execute permission (https://github.com/labring/sealos/pull/#1842)
* 01f2aa24 - 晓杰 - fix:change log level (https://github.com/labring/sealos/pull/#1838)
* 3a7941b1 - zzjin - Fix terminal rbac permission. (https://github.com/labring/sealos/pull/#1834)
* 7b2ebcf0 - zzjin - Update accounts.user.sealos.io sdk usage. (https://github.com/labring/sealos/pull/#1832)
* 26f9b501 - 晓杰 - feat:update role of user access ownerreference account (https://github.com/labring/sealos/pull/#1833)
* 66b1b2ef - leezq - feat: wechat pay (https://github.com/labring/sealos/pull/#1828)
* d949fe32 - leezq - Feat client sdk and wechat pay (https://github.com/labring/sealos/pull/#1827)
* 2103f964 - 小朝 - fix:update doc calico version to v3.24.1 (https://github.com/labring/sealos/pull/#1775) (https://github.com/labring/sealos/pull/#1826)
* 36852d24 - Mercurio - fix: ineffective env and cmd override (https://github.com/labring/sealos/pull/#1825)
* 816b6dfe - 晓杰 - fix:add create status when payment create (https://github.com/labring/sealos/pull/#1821)
* 9eaaa916 - zzjin - Dev front (https://github.com/labring/sealos/pull/#1818)
* 24fa9d97 - 晓杰 - feat: add user access account permissions (https://github.com/labring/sealos/pull/#1815)
* 3792dccc - HURUIZHE - CREATE SSH pair key when creating instance (https://github.com/labring/sealos/pull/#1808)
* 72ba1f51 - zzjin - Update user's kubeconfig using new user-controller's resp. (https://github.com/labring/sealos/pull/#1814)
* 275cfac1 - cuisongliu - feature(main): add disable_webhook_env (https://github.com/labring/sealos/pull/#1809)
* 0d9cd171 - cuisongliu - feature(main): add user namespace Pod Security Admission (https://github.com/labring/sealos/pull/#1811)
* 9c373176 - Bingchang Chen - fix: empty status error of quick query (https://github.com/labring/sealos/pull/#1810)
* 84441451 - muicoder - Generate certificate suitable for use with any Kubernetes Webhook. (https://github.com/labring/sealos/pull/#1805)
* 3d8261b1 - zzjin - Allow semgrep fails (https://github.com/labring/sealos/pull/#1806)
* 6b59e6e8 - zzjin - Update seg to standalone version. (https://github.com/labring/sealos/pull/#1803)
* 26e56065 - Bingchang Chen - feat: generate kubeconfig by user controller (https://github.com/labring/sealos/pull/#1799)
* 02663e81 - 中弈 - fixed docs, calico new version needs helm (https://github.com/labring/sealos/pull/#1802)
* 3bbcf069 - 中弈 - update readme (https://github.com/labring/sealos/pull/#1800)
* 879164f3 - HURUIZHE - Cluster Debug (https://github.com/labring/sealos/pull/#1792)
* 6e71ac0d - leezq - feat: add pwa support (https://github.com/labring/sealos/pull/#1785)
* 71f325f6 - cuisongliu - feature(main): add Platform for sealos pull/create (https://github.com/labring/sealos/pull/#1797)
* 5ed02d27 - cuisongliu - feature(main): fix user-controller-rbac for csr (https://github.com/labring/sealos/pull/#1796)
* 0687fd26 - zzjin - Update ci dep. (https://github.com/labring/sealos/pull/#1794)
* 5f6e81bd - BambooSword - bugfix: display the correct icon in app store (https://github.com/labring/sealos/pull/#1789)
* 3abbe6bf - Xia - update readme (https://github.com/labring/sealos/pull/#1722) (https://github.com/labring/sealos/pull/#1786)
* 10745a88 - cuisongliu - feature(main): update base image is ubuntu (https://github.com/labring/sealos/pull/#1784)
* 875bccce - 中弈 - update readme (https://github.com/labring/sealos/pull/#1782)
* 5acb7f5e - 中弈 - update docs (https://github.com/labring/sealos/pull/#1781)
* 1fda4fcc - zzjin - Dev front (https://github.com/labring/sealos/pull/#1778)
* 2c058a75 - leezq - add start menu  &  fix iframe url bug (https://github.com/labring/sealos/pull/#1780)
* 9c54aa27 - zzjin - Dev front (https://github.com/labring/sealos/pull/#1774)
* d58fa977 - cuisongliu - add user-controller webhook (https://github.com/labring/sealos/pull/#1766)
* 5b781cc7 - leezq - update: window manage (https://github.com/labring/sealos/pull/#1767)
* 35afa809 - gitccl - remove certificate from terminal-controller (https://github.com/labring/sealos/pull/#1765)
* 1669bc64 - zzjin - Temp Add `terminal`'s namespace (https://github.com/labring/sealos/pull/#1762)
* 5472f55f - zzjin - Update favicon& typo (https://github.com/labring/sealos/pull/#1763)
* da26e9f6 - HURUIZHE - Add Controller Cluster (https://github.com/labring/sealos/pull/#1697)
* 0c9a8047 - zzjin - Update cert manager to use dns01 for wildcard. (https://github.com/labring/sealos/pull/#1761)
* 2320a6ee - zzjin - tmp fix terminal waiting (https://github.com/labring/sealos/pull/#1757)
* 26469b16 - cuisongliu - feature(main):  fix user group role (https://github.com/labring/sealos/pull/#1754)
* bbd304c6 - zzjin - Init support of terminal and kubernetes-dashboard. (https://github.com/labring/sealos/pull/#1738)
* 9f5c0ba6 - cuisongliu - feature(main):  add registry sdk (https://github.com/labring/sealos/pull/#1746)
* ce95080b - cuisongliu - feature(main):  add sync user role (https://github.com/labring/sealos/pull/#1750)
* 80978460 - cuisongliu - feature(main):  add sync user role (https://github.com/labring/sealos/pull/#1749)
* 71d1ef74 - gitccl - fix ci terminal controller build error (https://github.com/labring/sealos/pull/#1748)
* 7a9346ad - gitccl - update terminal controller image, disable CGO (https://github.com/labring/sealos/pull/#1745)
* e940d903 - fengxsong - feat: generate config of components (https://github.com/labring/sealos/pull/#1743)
* 8847db03 - zzjin - update go work modules (https://github.com/labring/sealos/pull/#1742)
* c35e0613 - zzjin - Fix user controller module (https://github.com/labring/sealos/pull/#1739)
* 5585754a - fengxsong - refactor: applier (https://github.com/labring/sealos/pull/#1736)
* f0813767 - fengxsong - fix: issue 1724 (https://github.com/labring/sealos/pull/#1733)
* 8700ffb4 - zzjin - Update Frontend. (https://github.com/labring/sealos/pull/#1726)
* f12035c2 - cuisongliu - feature(main):  add sync user role (https://github.com/labring/sealos/pull/#1718)
* 4646ab41 - zzjin - Update and fix login method. (https://github.com/labring/sealos/pull/#1719)
* 3c848075 - pangqyy - update readme (https://github.com/labring/sealos/pull/#1725)
* b7a9ce43 - cuisongliu - feature(main): sealos merge fearure (https://github.com/labring/sealos/pull/#1706)
* cace461b - PangQYY - update readme (https://github.com/labring/sealos/pull/#1722)
* a3f15bd5 - PangQYY - update readme (https://github.com/labring/sealos/pull/#1721)
* 97a1b547 - YAO - Fix Optimized to use buildah sdk with push and pull (https://github.com/labring/sealos/pull/#1714)
* 2fc34c90 - zzjin - Update go module to match workspace (https://github.com/labring/sealos/pull/#1716)
* 2ba813bc - leezq - fix: add index page (https://github.com/labring/sealos/pull/#1715)
* 61da0870 - muicoder - feature(main): optimize Dockerfile(https://github.com/labring/sealos/pull/#1711) (https://github.com/labring/sealos/pull/#1713)
* b2ba9705 - cuisongliu - feature(main):  add upx for sealos (https://github.com/labring/sealos/pull/#1710)
* 868104f0 - cuisongliu - feature(main):  add binary for buildah (https://github.com/labring/sealos/pull/#1709)
* af1390bf - zzjin - Frontend updates (https://github.com/labring/sealos/pull/#1707)
* 1c6a9683 - zzjin - remove upx compress on sealos binary (https://github.com/labring/sealos/pull/#1705)
* 617e8773 - YAO - fix sealos will not use local images if there is a newer one remote (https://github.com/labring/sealos/pull/#1689)
* 6f222ca0 - zzjin - update setuo go with cache to speedup ci (https://github.com/labring/sealos/pull/#1699)
* a1c969ef - cuisongliu - feature(main): fix upx for release (https://github.com/labring/sealos/pull/#1702)
* 7933121f - Bingchang Chen - feat(auth): casdoor  static file server (https://github.com/labring/sealos/pull/#1698)

**Full Changelog**: https://github.com/labring/sealos/compare/v4.1.2-alpha11...v4.1.4-alpha3
