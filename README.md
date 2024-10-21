# hs
HarmonyOS Magic


## Acknowledgements
- [Huawei Developers - Signing Your App/Service](https://developer.huawei.com/consumer/en/doc/harmonyos-guides-V5/ide-signing-V5#section297715173233)
- [Huawei Developers - Releasing Your App/Service](https://developer.huawei.com/consumer/en/doc/harmonyos-guides-V5/ide-publish-app-V5)
- [Huawei Developers - Application Development Overview (ArkTS)](https://developer.huawei.com/consumer/en/doc/harmonyos-guides-V3/ide_debug_device-0000001053822404-V3#section677893315228)
- [Huawei Developers - Debugging Commands](https://developer.huawei.com/consumer/en/doc/harmonyos-guides/debugging-commands-V5)
  - [Debugging Commands - bm tool](https://developer.huawei.com/consumer/en/doc/harmonyos-guides/bm-tool-V5)
  - [Debugging Commands - aa tool](https://developer.huawei.com/consumer/en/doc/harmonyos-guides/aa-tool-V5)
  - [Debugging Commands - hdc tool](https://developer.huawei.com/consumer/en/doc/harmonyos-guides/bm-tool-V5)
- Generate the `.p12` keys using cmd: [Huawei Developers - Releasing Your App/Service](https://developer.huawei.com/consumer/en/doc/harmonyos-guides-V2/publish_app-0000001053223745-V2)


```shell
keytool -genkeypair -alias "ide_demo_app" -keyalg EC -sigalg SHA256withECDSA -dname "C=CN,O=HUAWEI,OU=HUAWEI IDE,CN=ide_demo_app"  -keystore d:\\idedemokey.p12 -storetype pkcs12 -validity 9125 -storepass 123456Abc -keypass 123456Abc
```