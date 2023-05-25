# Sealos v4.2.0-rc4 🎉🎉

We are excited to announce the official release of Sealos v4.2.0-rc4 🎉🎉!

## Sealos Cloud: Powerful Cloud Operating System Distribution

Sealos Cloud is a cloud operating system distribution with Kubernetes at its core. Users can directly use Sealos Cloud or run Sealos in their private environment to have the same capabilities as Sealos Cloud. Sealos Cloud offers a range of advantages, including a sleek product experience, fully open-source architecture, consistent public and private cloud experiences, cross-platform compatibility, and highly competitive pricing.

### Sealos Cloud Usage Guide

Sealos Cloud offers you exceptional public cloud services for cloud-native applications, making it easy to manage cloud-native applications. Sealos Cloud provides two ways to use: cloud access and private access. The online mode is now officially launched, and offline mode will be introduced in future releases.

#### Cloud Access

Cloud access is provided by directly accessing the cloud services provided by Sealos Cloud. Just enter the following link in your browser to start using the powerful features of Sealos Cloud immediately:

```
https://cloud.sealos.io
```

Cloud access allows you to access and manage your cloud-native applications anytime, anywhere, without any additional configuration and deployment. This makes the online mode an ideal choice for quickly getting started with Sealos Cloud.

### Component Introduction

Sealos Cloud's main features are divided into the frontend interface, backend API services, and Kubernetes Operator, working together to provide a complete cloud-native application management experience.

#### Main Components

- **Auth-Service**：Provides authentication services using casdoor, ensuring the security and accuracy of user identities.
- **Image Hub**：Sealos image repository-related services, making it easy for users to manage and use images.
- **Desktop**：Public cloud frontend service, providing a friendly user interface and operation experience.
- **User**：User and user group management services, making it easy for administrators to assign and manage user permissions.
- **Account** & **Metering**：Provides billing and account capabilities, helping users to reasonably control and plan costs.
- **App**：Provides the Sealos Cloud desktop application, making it easier for users to use and manage cloud services.
- **Infra**：Provides basic settings, currently supports AWS and Alibaba Cloud, and may support more cloud service providers in the future.
- **Cluster**：One-click cluster startup on Sealos Cloud, simplifying the cluster deployment and management process.
- **Terminal**：Terminal services on Sealos Cloud, providing convenient access and management of cloud services.

For more detailed documentation about Sealos Cloud, please visit `https://sealos.io/docs/cloud/Intro`.


## Sealos Boot: Professional Cloud-Native Application Management Tool

Sealos Boot is the core component of Sealos, mainly responsible for the lifecycle management of clusters, downloading and deploying OCI-compatible distributed applications, and customizing distributed applications.


### How to Install

#### Binary Installation

```shell
    curl -sfL  https://raw.githubusercontent.com/cuisongliu/sealos/v4.2.0-rc4/scripts/install.sh \
    | sh -s v4.2.0-rc4 cuisongliu/sealos
```



### Component Introduction

Sealos provides two Docker containers: sealos and lvscare, as well as two binary files: sealctl and image-cri-shim. Below is a brief introduction to these components:

- Sealos Container: This container is the core component of the Sealos project, responsible for deploying and managing Kubernetes clusters and distributed applications. It offers a range of command-line tools to help users quickly build and maintain clusters.
- lvscare Container: This container is used to support load balancing management within Kubernetes clusters. It can monitor node status in real-time, ensuring that the load balancer always routes traffic to available nodes.
- sealctl Binary: This is a command-line tool provided by the Sealos project, used to simplify cluster management tasks such as certificate management, IPVS, hosts, and cluster certificates.
- image-cri-shim Binary: This component is a CRI adapter for the Sealos project, supporting different container runtimes (such as Docker and containerd). It allows Sealos to seamlessly integrate with multiple runtimes, enhancing the project's flexibility and scalability.

### Quick Start

```shell
# Create a cluster
sealos run labring/kubernetes:v1.25.0 labring/helm:v3.8.2 labring/calico:v3.24.1 \
    --masters 192.168.64.2,192.168.64.22,192.168.64.20 \
    --nodes 192.168.64.21,192.168.64.19 \
    --passwd your-own-ssh-passwd
```


## Changelog
### New Features
* 4131325b32ce4ba7445b1ca2202d02961406612a: feat(debt): Debt-controller should delete user resources over 7 days in arrears  (#2899) (@xiao-jay)
* 90b9a163eafe28c185a3e752cacda171872a9d57: feat: Add support for single-node mode without hosts in Clusterfile (#3024) (@LZiHaN)
* ef9d393890ffad3eacddf03b736324bb5ac0a4d0: feat: add infra metering implementation with e2e test (#2797) (@xiaohan1202)
* 6c0a2163da5c3b73c04c1e3656ef23090d5f7910: feat: add resource group (#2905) (@xiaohan1202)
* 014bc8c698bdc177bc63fbd2569aad7cc4446b25: feat: cmd and params input way (#3074) (@c121914yu)
* 34d3f8a3a56bcb3ac8c01185cebc3d11f5faf956: feat: desktop pwa (#3077) (@zjy365)
* bce2fc384a13175a58002918e93f0cdeba09e134: feat: env set secretName (#3095) (@c121914yu)
* 6ef4df2f9048cee482235840f0df077552f7286f: feat: grouping subcommands (#2920) (@fengxsong)
* d4a9076ead5f9a24e38d7cce97bee74571976df9: feat: logs analyses (#3041) (@c121914yu)
* 1c6dd12c36754ec4cc75c9889fd852e45fdf0336: feat: open other app by sdk (#3011) (@c121914yu)
* 5207f953494c6104a8644e0d56d1ad175545bd09: feat: resources monitor, metering controller, billing controller, billingrecordquery, pricequery controller (#3008) (@bxy4543)
* 3c60b6a2a6bedf90204f43956be26394d5e2ac06: feat: terminal add loading animation (#3097) (@zjy365)
### Bug fixes
* 16ab2d2c81a3d1312937aa21538763f9080ea2b6: bugfix: fix duplicate print prompts (#3020) (@nowinkeyy)
* 5bb06f4e82b1facc0e4d58358f586ed1690686a4: bugfix: the failure to add master after updating the certificate (#2942) (@bxy4543)
* f7bb82c36605e341d9a74bd1734e2a5a162a4ffe: fix(ci): fix metering e2e install account error (#2901) (@xiao-jay)
* 96a404b0fd07c85fd5ebc3c210358d3297c4540b: fix: automatic completion of arch role and initialize default SSH in single-node mode (#3102) (@LZiHaN)
* 0c4a3cf71228a1a304019019126bdd8dfc154590: fix: check sg & vswitch existence (#2937) (@xiaohan1202)
* 369635f69e578bd4a52a186785431118041ee3b1: fix: compress/uncompress offline images (#2966) (@fengxsong)
* a968e1363eb15de8e3096cebeeb8bcbaaabb1106: fix: desktop image style (#3105) (@zjy365)
* 3de1dc66a4b6cfbd5a9f200ab071fabce5787eac: fix: fix always return the insecure registry (#2927) (@fengxsong)
* 57a04dac0b9b80bb0fde09b18be8784be9236109: fix: invalid reference format (#3070) (@fengxsong)
* 472d25c59ef7bab4bbd3acedbcae0d8895183d0a: fix: make pre-deploy error (#2904) (@xiao-jay)
* 2ecc16d047ec88140ac68aba313fc2b0f642b098: fix: modify the desktop height acquisition method (#3096) (@zjy365)
* 884f1d9d0ff23d96e17a43980dc37f107afc7fd4: fix: return image name if oci archive file has names instead of id (#2887) (@fengxsong)
* 09aa95f0c36e11ade2851c2b328e70e704fa501c: fix: rm master taint in single model (#2860) (@mixinkexue)
* 406ec806812e1fb344ebf9ff3cb7272abb687394: fix: running in container as user root (#3056) (@fengxsong)
* ddabe65678f86343b40687d1059d9b87136a3d1c: fix: use assertion, avoid patching external interface (#2891) (@fengxsong)
### Other work
* 8783311762ac938eb035a5c6ef15baf3aecffd72: Automated Changelog Update: Update directory for v4.2.0 release (#2977) (@sealos-ci-robot)
* aa72e131e7a0dc28e0ed3d096489709cc6c4b2ae: Automated Changelog Update: Update directory for v4.2.0-alpha3 release (#2964) (@sealos-ci-robot)
* 6e90dd63626a93852320806eb32de3a993f0df96: CI: Switch build dev to release (#3062) (@lingdie)
* fc7a761d8a501a12a77f70ca5863e0525fb2c1ca: CI: fix image build ci. (#3065) (@lingdie)
* 9aa43c8b4f15a14c60153d7752a1141363b86772: Cluster image desktop (#3028) (@lingdie)
* 3a66249c892341f44e2f7a7a637a9f505e0dcaa0: Dev adminer (#3120) (@zzjin)
* 89a519b55ecf9c266f55953d5a84ed0270d1fda3: E2e/apply test workflow  (#2952) (@bxy4543)
* 852b6a701896ffaa53b65869c7b157cda3a7f8c6: Fix bytebase Dockerfile to support outside binary. (#3116) (@zzjin)
* 96718247370bb04c9e5bd1d53b64df08f59843db: Fix bytebase build CGO&GOOS. (#3117) (@zzjin)
* 6b4e5618fefdf7656c97c0d191088242f2f167af: Fix bytebase nginx controller. (#3111) (@zzjin)
* 08f3246bf569ad63a9a227898a748c1a64102caf: Fix demo usage of side-login ingress. (#3072) (@zzjin)
* 31e40627dc8db35dfa2681804b2b895c653bcc1e: Fix login ingress deploy (#2909) (@zzjin)
* 139ed51e651141f92d9bae610eac2e55e109fcd7: Fix readme md usage. (#3046) (@zzjin)
* 441113ca103396d2c035181d44a2d46975cf3c87: Init adminer controller. (#3107) (@zzjin)
* 9b7fd80b7659f57e87b2b9f09e8e103e575a2e14: Move deploy doc. (#2951) (@zzjin)
* cdacf36cb0918462e6c517a342b00d40bd61295a: Reduce terminal limits a little. (#2915) (@zzjin)
* 421bed042d5dc07d15b359710b12b64678f175fa: Refactor preprocessor (#3101) (@fengxsong)
* 3cccfbaab4a99b023d28e2b684fec8b65172beb5: Replace all `math/rand` to safer `crypto/rand`. (#3031) (@zzjin)
* ff35c3e6ef8415067dace67c2256dd0750457a12: Stream log (#3018) (@c121914yu)
* 2d39bc86e7249a627ad01a7cdc799ea4a706bb19: Terminal controller support cors. (#2991) (@zzjin)
* 5efa4b1c2d6748b8542ff3c6f958f7527350496b: Title: Update auto-release logic to skip deb and rpm packages for non-production releases (#2882) (@cuisongliu)
* adc8cb7347f5b1cf1c5226b112b2445c046bdcdd: Try add db/bytebase controller ci. (#3115) (@zzjin)
* 015f9c54c4074ac18e8541fb59c2593fdb823b42: Update README, add the new version UI showcase. (#3040) (@fanux)
* bfee712a45f1e1590b75c69ae9fce3752241903e: Update image (#3081) (@lingdie)
* 0afd2edb04cceb57b270e25c472d3772ac0346db: Update pr,support auto `copilot:all`. (#3118) (@zzjin)
* b090af6800e20c6f8087988fef6e6a1f23f0680b: add app providers logo (#3054) (@zjy365)
* d528d6be713b9b9cf92169e5822d354d29fffb9d: add auth cluster image and update workflow (#2981) (@lingdie)
* 9725dedb5dde50e8bb6b04f975d17a1a97bce54d: add bytebase controller (#2841) (@dinoallo)
* b85eda0efc3a3986fd4d392eca3b17ecfed585e8: add cert prepare readme (#3100) (@lingdie)
* 42116471743184552a6c43b7803e4da71a954198: add client id and secret in auth svc (#3078) (@lingdie)
* db9c47d9d355599f3c62bdf5ce13a6f672a358c5: add create dev cluster image ci (#3058) (@lingdie)
* 1fa3c8dc16c161a684401d8591b67ea6cf212139: add doc/4.0/docs/cloud/examples/how-to-deploy-postgresql-with-kubeblocks-on-sealos-cloud.d doc/4.0/i18n/zh-Hans/cloud/examples/how-to-deploy-postgresql-with-kubeblocks-on-sealos-cloud.d (#3045) (@sakcer)
* 7ce51bb3074fbb8b2f4d9b8bd9c1ec20306c3f90: add e2e apply test (#2936) (@bxy4543)
* fba6b37728e840b15e41b5713d19d7438bcae479: add fetch Kubeadm Config: (#2943) (@bxy4543)
* b2ed4c7b9754091a4bf31311e8bda970101e8263: add image-cri-image test (#2849) (@bxy4543)
* 21cdb8bc7a3b9648e97a9ce006d47b3754af5541: add local deplopment docs. (#2945) (@lingdie)
* df615299d952dac6ad03d59cadf3775088ff86e8: add systemDB design (#2955) (@fanux)
* 6f1bbae0a14f0fe9beb5c0a8b14db6e4f0ab1131: change cert env (#3033) (@lingdie)
* 8a67f568a58287dbe2d0108fbb0c9d20598cb00d: ci，docs: add account e2e test and docs (#2935) (@xiao-jay)
* 2d8d793e7eaa338aa8261075d6b16596cb06f1cd: docs(main): Refactor Changelog generation script and update usage guide (#2884) (@cuisongliu)
* 6f0accc78c98843c171f1c2e6c725731fb8d0175: docs(main): add changelog action (#2990) (@cuisongliu)
* 72a58a94031eb8ceadcd35aca3c67df43bbf8cea: docs(main): add changelog full to shell (#2885) (@cuisongliu)
* 53974f5e7498593af7384c0a77062f551311711a: docs(main): add debug for merge logger (#2886) (@cuisongliu)
* 6406b475674ffc030579e61c17ea71d99fa2281d: docs(main): add default image config.yml to config patch image (#3006) (@cuisongliu)
* a100d1a5d4136369b49dfefb02e1dda58203ceba: docs(main): add docs for sealctl (#2922) (@cuisongliu)
* c781a29e50c45a57254c66eb7a09c25013d9ffd9: docs(main): add docs for user (#2929) (@cuisongliu)
* c8d475710ab305cc20a272fe2010cfec284dd846: docs(main): add e2e test images (#3013) (@cuisongliu)
* a6e33bc799c2b760cad3ff38ad55afc02212f84b: docs(main): add image guide docs (#2992) (@cuisongliu)
* e5a204154f236622aebcccc51ba10df5f32d84b0: docs(main): add image-cri-shim docs (#2880) (@cuisongliu)
* de11f2b790efa907a4c62c4e934ba16a60e594a0: docs(main): add robot for sealos (#2959) (@cuisongliu)
* d4bc9beba4a91fd074a7dca12e8bb79f9cfe36ae: docs(main): change issue token (#3106) (@cuisongliu)
* 8f9480316be65a401ad4cc510899dac9fb6ff8ea: docs(main): changelog body support (#3044) (@cuisongliu)
* e26332d3fedc8c684481515e64d7a3fce7b2be9f: docs(main): delete build sealos image (#2975) (@cuisongliu)
* c552e68b0097c256202ac0566f2a6ec9b3037e9f: docs(main): e2e for core sealos (#3030) (@cuisongliu)
* 9f8189aa7a3131c716ed6a8d2887d1fad347c863: docs(main): e2e for image shim for sealos (#3034) (@cuisongliu)
* f696a621e5a91ce771b725e8b465f7f84f2aaa13: docs(main): fix bot config (#2976) (@cuisongliu)
* bb7e82da0380547aae9c456f3fa0aa1cd2a305bb: docs(main): fix bot config (#3021) (@cuisongliu)
* 8a521506e2c8b335225af8f24c3e61c5aca9d0a2: docs(main): fix change owner (#2973) (@cuisongliu)
* 967baa6b081135e7579a74d9b45e82247401227c: docs(main): fix change owner (#2974) (@cuisongliu)
* 8679b13f08dce0c72cb1ae5bd66e05e0ca1474a0: docs(main): fix changelog for sealos (#2963) (@cuisongliu)
* 908039e028e6128a7ff62f5a1099fd251d97599f: docs(main): fix ci patch images (#2971) (@cuisongliu)
* fa1776069c2633e02f072eb814a225ff02c737d6: docs(main): fix image0 to scratch (#2969) (@cuisongliu)
* 453218412cac23039c9cbf74b61924d2dedbad54: docs(main): fix release comment (#2982) (@cuisongliu)
* 8e1d6afebc1e816412ad98382d8abc194bb2f9e5: docs(main): fix run workflow to test-run (#2970) (@cuisongliu)
* 19ece59be05bb572a76c36ccfdaf110fa877f1ec: docs(main): remove registry to sealos Experimental (#3047) (@cuisongliu)
* 0d2588faf32e51c4467c66b106c4018276d73af6: docs(main): replace lookup env to apply (#3039) (@cuisongliu)
* 884dfb0845d00f65f49782c50127e321d85eb1fb: docs(main): save sealos for controllers (#3026) (@cuisongliu)
* e758ed41264cc71d99e366185f7d2d572aa3653c: docs(main): save sealos for controllers,front,service (#3029) (@cuisongliu)
* 42e06f9ea34e7db1edee1277f2050601ffd5940b: docs(main): support v1.27 docs (#2939) (@cuisongliu)
* 60b038828da1ac2b0e915789b311104e05149d6c: docs(main): support v1.27 k8s (#2938) (@cuisongliu)
* 8c4a998bb45d48396878a28ff690c286a95720d2: docs(main): sync code user and email (#2900) (@cuisongliu)
* 18b344744a4df5c387c7323818b3c655ad805bd8: docs(main): update gorelease config for new release (#2953) (@cuisongliu)
* fbec0218a84ddd372b71b14bf30654f91bfd50e1: docs(main): update latest sealos version (#3104) (@cuisongliu)
* 4b5f9d794435f9f9835bf43ed74ccb4670ef2c55: docs(main): upgrade bot version (#3022) (@cuisongliu)
* a929e01b639a409d03c9db06694f9a01e0fcb936: docs(main): upgrade gomod (#2894) (@cuisongliu)
* 696c7415e2f196dd40722fb8ee4bfc53df5924e7: docs: how to write reousrce-controller (#2941) (@xiao-jay)
* e43c77cbf38e1d6fb4bd591621323a40e0e6a51e: docs: translate metering design from CN to EN (#2746) (@xiao-jay)
* fa063c1fa9c68c473a395d94f12b17d3270d8737: docs: upgrade bot config (#2968) (@cuisongliu)
* b49983dc993b0ffe0ac33b7f84a3589ec9c9afd4: feat add bytebase application (#2925) (@zjy365)
* 6e1db7b96a2d7631777677d0f728db3b03ad9b42: feat terminal add env ttyd image (#3012) (@zjy365)
* f8269fb07b215c76eee37ed6d08fd47ee4d14d34: feat terminal cors & decodeURIComponent url command (#2988) (@zjy365)
* 147ed6d372d3281f8a6e666db61572b07a46c5f3: feat web terminal frontend (#2972) (@zjy365)
* 22bdba549f669086556bf06e8080c8da1fa6d535: feat. build frontend cluster image  (#3016) (@lingdie)
* 5b8a1c33236195ca16055043728ee779ce599b89: feat. fix cloud cluster image (#3087) (#3089) (@cuisongliu)
* 926314a6e06d61bf1d484a9d6713794de242bf2a: feat. fix cloud cluster image (#3087) (@cuisongliu)
* f8224e0e473fa90e753be2ae11cb8d18205b921c: feat. init cloud cluster image (#3082) (@lingdie)
* c93c368967bf38061d3ff7bf8b5371554a75414e: feature: support checking cluster's no-odd number of master node (#2842) (@pixeldin)
* 9ca8b5e54d592c69aa33676f1540d575ee055c5a: feature: support pulling multiple images at the same time (#2931) (@dhanusaputra)
* 3a7970cc1b3512517020627bbaf4fdc2e5ab1af4: feature: using promptui replace confirm (#2993) (@nowinkeyy)
* 5feb91b71a25eed7ef534e87963a88bc99c2dddd: fix GetHostArch without use args port (#3042) (@ghostloda)
* e8fd11b8d356956a4890f7d898595c1d69ab5dfc: fix app cr name. (#3080) (@lingdie)
* ba49b269d884aa8dec6f0d379d85952fd4195564: fix app launchpad env (#3099) (@lingdie)
* 8438a3fd3eacc97cca798d310e9e657270448eb7: fix ci error (#3128) (@bxy4543)
* f64db58eada531c70575da32d012bdc83b320864: fix cluster image (#3064) (@lingdie)
* ee2693f9eabb7c5f672a32c3484dff056eaaadea: fix deploy file (#3093) (@lingdie)
* 1b842e244ffa500afa74ea0dd81c265bd6cf42f7: fix image name (#3066) (@lingdie)
* d5cf8f30c7f52795ff68ba4537c582558fb2c620: fix ingress's backend error (#3071) (@lingdie)
* 9256d27ef8f0d840f44f02526f64e46437b74695: fix len(deleteInstancesByOption.instanceIds) always gt 0 (#2934) (@bxy4543)
* 3018f694ebfdb918f1a577d9dffe59afd03cc578: fix multi infra apply (#2960) (@bxy4543)
* 7fd1373459b7e1172f1f6f6fe2281271b79c914f: fix notification (#2895) (@zjy365)
* 2fc130bec22217eb015b777ab48f971f97838834: fix pgsql status null (#2924) (@zjy365)
* 59723ed1820286507549987123027da9f2275085: fix role binding (#3088) (@lingdie)
* e71f980be1d586648b757d6938c8113e7ad1897e: fix sha len not same (#3025) (@lingdie)
* a26403fc794a015d97f89996e0431b3f7219388f: fix stop kube-apiserver pod cause `DeadlineExceeded` when update cert. (#2918) (@bxy4543)
* 86e8983b62d5296c3908ad3e4b8bc252cd620716: fix sync docs error (#2944) (@xiao-jay)
* d4f5eb26752a5cdc291bc36b992147715c6feffa: fix sync tag name (#3015) (@lingdie)
* e5ab5095ad4171c989dcde9d58a1a569ae222f24: fix terminal and applaunchpad (#3091) (@lingdie)
* b68437084ba4a3b8de520b258332469c2dfeaa17: fix. cluster image controller (#3017) (@lingdie)
* be282f8d12d25570ea8bb4171300889ee240cf80: format(account): format webhook name (#2928) (@xiao-jay)
* bf4ed6608ecf883ce9a207b39ec89e4ed831424a: move registry save command to sealctl (#3122) (@fengxsong)
* d57b729566c85aee12065a3d0f0388b4424c021c: new app: launchpad (#2987) (@c121914yu)
* 168319f56bfecffbbf819b32c97d5fa487eaf691: optimize reset node `ip link delete` command. (#2947) (@bxy4543)
* 7448fe0f5a92c9a1d64ab0ac6157909bdfc2a325: perf: default domain (#3092) (@c121914yu)
* 6e66a3bcfee01d79d6972f8634ca263ebec0382a: perf: export yaml name;focus restart app; (#3067) (@c121914yu)
* c9f4f58e1fafee9adddb1ab149faa4622281e62d: refactor(main): add bot for link check (#3083) (@cuisongliu)
* f26ad3d474cb5c044dbe5b619a0d09924004094c: refactor(main): add bot for link check (#3085) (@cuisongliu)
* 89342db1f3c9459146476c9769f303caa2378f7b: refactor(main): add image replace status interface (#3060) (@cuisongliu)
* a3dd0772f3260b4b4c12616ba6cbce630c7ba729: refactor(main): add operator interface (#3048) (@cuisongliu)
* a1c6a179d63ed78a0468897e55c8efe0b9922fd7: refactor(main): add registry docs and fix registry save bug (#3124) (@cuisongliu)
* 857a2386a8b420593bf8b25df6465849c68599b0: refactor(main): change config to envs (#3121) (@cuisongliu)
* bcd5a9bda17dbdc8912df417ddb86980275dd3db: refactor(main): enable build flags isolation (#3086) (@cuisongliu)
* b9705230281fdd3f975ae53a1cf9b7b0cccec30c: refactor(main): fix release error (@cuisongliu)
* 01eb8028b65eba664653651d352700a8062ea9eb: refactor(main): fix release error (@cuisongliu)
* 2a83f67da4b4622559ab5b364178b0bf58a77382: refactor(main): replace registry password config generator (#3119) (@cuisongliu)
* 266200f83cd166ec39883d7cf49a467c3e656c64: refactor(main): trigger release ci (#3063) (@cuisongliu)
* 7c66fd20cf6c424c4cc492bf119c38115d72c141: refactor: add back auto build and auto save abilities in `diff` command (#2906) (@fengxsong)
* 187be92fc0d62898683db58725a6a9a79cf7bb76: refactor: diff command (#2879) (@fengxsong)
* 6bdd9724cb2a042660f456a16cc089c107909e4b: refactor: rename interface and function names (#3112) (@fengxsong)
* d4a0c5153f4619e9adf10319d7afaf1df490bec1: refactor：smooth and simple new desktop (#3059) (@zjy365)
* 7af550569553a83067999650462a750d21aa894f: rename CopyR to Fetch (#3061) (@fengxsong)
* 189847f16cce2ee0e306e3946452ad6ab5a78fe9: temp commit (#3123) (@zzjin)
* f8774d06b944655a16fbcc1f65f31f1c5d6c4c5e: update controller and service cluster image build (#3001) (@lingdie)
* 4636666870934fec7e1914dbffe432fc2f77dab4: 🤖 add release changelog using rebot.<br/> (#3043) (@sealos-release-rebot)

**Full Changelog**: https://github.com/cuisongliu/sealos/compare/v4.2.0-alpha2...v4.2.0-rc4

See [the CHANGELOG](https://github.com/cuisongliu/sealos/blob/main/CHANGELOG/CHANGELOG.md) for more details.

## Roadmap

In the future development plan, Sealos Boot and Sealos Cloud will continue to expand their capabilities to meet the needs of more users. Our Roadmap includes the following key plans:

1. Sealos Cloud Private Deployment: We will introduce a private deployment solution for Sealos Cloud, allowing users to deploy and run Sealos Cloud in their own data centers or private cloud environments, achieving consistency between public and private cloud experiences.
2. Heterogeneous Deployment Support: Sealos Boot will support the deployment of Kubernetes clusters on various hardware platforms and operating systems, achieving broader compatibility and flexibility.
3. K3s and K0s Integration: We plan to integrate the lightweight Kubernetes distributions K3s and K0s into Sealos Boot, allowing users to choose different Kubernetes distributions for deployment according to their needs.

Please look forward to the launch of these new features. We will continue to work hard to provide users with better services and experiences.

We are very proud to introduce the two major functional modules of Sealos Cloud and Sealos Boot, which we believe will bring more convenience and efficiency to your cloud-native application management. We look forward to your feedback and suggestions so that we can continue to improve and provide better products and services.

Thank you for your support of Sealos🎉🎉.

If you encounter any problems during use, please submit an issue in the [GitHub repository](https://github.com/cuisongliu/sealos) , and we will solve your problem as soon as possible.
