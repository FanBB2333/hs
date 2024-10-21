# [WIP] hs - HarmonyOS Dev-App Install
HarmonyOS Dev-App Install

## Steps
**本项目仍在开发中，以下步骤节选自华为开发者文档，仅供参考。**

参考[通过命令行方式构建应用或服务](https://developer.huawei.com/consumer/cn/doc/harmonyos-guides-V3/command-line-building-app-hap-0000001193655754-V3)
### 1. 准备工作
按需安装`java`, `HarmonyOS SDK`, `DevEco Studio` 等工具

### 2.构建HAP
获取未签名的HAP文件或自行编译

### 3.[配置应用签名信息](https://developer.huawei.com/consumer/cn/doc/HMSCore-Guides/harmonyos-java-config-app-signing-0000001199536987)
其中3.2-3.4需要在[AppGallery Connect](https://developer.huawei.com/consumer/cn/service/josp/agc/index.html)完成

#### 3.1 生成密钥和证书请求文件
密钥：格式为.p12，包含非对称加密中使用的公钥和私钥，存储在密钥库文件中，公钥和私钥对用于数字签名和验证。
证书请求文件：格式为.csr，全称为Certificate Signing Request，包含密钥对中的公钥和公共名称、组织名称、组织单位等信息，用于向AppGallery Connect申请数字证书。

#### 3.2 申请应用调试证书
数字证书：格式为.cer，由华为AppGallery Connect颁发，分为应用调试证书和应用发布证书。

#### 3.3 注册调试设备

#### 3.4 申请调试Profile
Profile文件：格式为.p7b，包含HarmonyOS应用的包名、数字证书信息、描述应用允许申请的证书权限列表，以及允许应用调试的设备列表（如果应用类型为Release类型，则设备列表为空）等内容，每个应用包中均必须包含一个Profile文件，分为调试Profile和发布Profile。

### 4.HAP签名
参考[使用调试证书对HAP进行签名](https://developer.huawei.com/consumer/cn/doc/harmonyos-guides-V3/command-line-building-app-hap-0000001193655754-V3#section18188942105313)或[为应用/服务进行签名](https://developer.huawei.com/consumer/cn/doc/harmonyos-guides-V3/ide_debug_device-0000001053822404-V3#section677893315228)

### 5.使用`hdc`工具安装HAP(仅手动签名需要)
参考[`hdc install`相关文档](https://developer.huawei.com/consumer/cn/doc/harmonyos-guides/hdc-V5#应用相关命令)

## Acknowledgements
- [Huawei Developers - Signing Your App/Service](https://developer.huawei.com/consumer/cn/doc/harmonyos-guides-V5/ide-signing-V5#section297715173233)
- [Huawei Developers - Releasing Your App/Service](https://developer.huawei.com/consumer/cn/doc/harmonyos-guides-V5/ide-publish-app-V5)
- [Huawei Developers - Application Development Overview (ArkTS)](https://developer.huawei.com/consumer/cn/doc/harmonyos-guides-V3/ide_debug_device-0000001053822404-V3#section677893315228)
- [Huawei Developers - Debugging Commands](https://developer.huawei.com/consumer/cn/doc/harmonyos-guides/debugging-commands-V5)
  - [Debugging Commands - bm tool](https://developer.huawei.com/consumer/cn/doc/harmonyos-guides/bm-tool-V5)
  - [Debugging Commands - aa tool](https://developer.huawei.com/consumer/cn/doc/harmonyos-guides/aa-tool-V5)
  - [Debugging Commands - hdc tool](https://developer.huawei.com/consumer/cn/doc/harmonyos-guides/bm-tool-V5)
- Generate the `.p12` keys using cmd: [Huawei Developers - Releasing Your App/Service](https://developer.huawei.com/consumer/cn/doc/harmonyos-guides-V2/publish_app-0000001053223745-V2)
- [Huawei Developers - 通过命令行方式构建应用或服务](https://developer.huawei.com/consumer/cn/doc/harmonyos-guides-V3/command-line-building-app-hap-0000001193655754-V3)
- [Huawei Developers - 调试应用（HarmonyOS）](https://developer.huawei.com/consumer/cn/doc/app/agc-help-debug-app-0000001914423098)
