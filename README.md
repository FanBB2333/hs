# hs
HarmonyOS Magic


## Acknowledgements
- [Huawei Developers - Signing Your App/Service](https://developer.huawei.com/consumer/en/doc/harmonyos-guides-V5/ide-signing-V5#section297715173233)
- [Huawei Developers - Releasing Your App/Service](https://developer.huawei.com/consumer/en/doc/harmonyos-guides-V5/ide-publish-app-V5)
- Generate the `.p12` keys using cmd: https://developer.huawei.com/consumer/cn/doc/harmonyos-guides-V2/publish_app-0000001053223745-V2
```shell
keytool -genkeypair -alias "ide_demo_app" -keyalg EC -sigalg SHA256withECDSA -dname "C=CN,O=HUAWEI,OU=HUAWEI IDE,CN=ide_demo_app"  -keystore d:\\idedemokey.p12 -storetype pkcs12 -validity 9125 -storepass 123456Abc -keypass 123456Abc
```