package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func generateP12File(path, alias, password string) error {
	cmd := exec.Command("keytool",
		"-genkeypair",
		"-alias", alias,
		"-keyalg", "EC",
		"-sigalg", "SHA256withECDSA",
		"-dname", "C=CN,O=HUAWEI,OU=HUAWEI IDE,CN="+alias,
		"-keystore", path,
		"-storetype", "pkcs12",
		"-validity", "9125",
		"-storepass", password,
		"-keypass", password,
	)

	output, err := cmd.CombinedOutput()
	// print the output
	fmt.Println(string(output))
	if err != nil {
		return fmt.Errorf("error generating p12 file: %v, output: %s", err, output)
	}

	return nil
}

func generateCSRFile(keystorePath, alias, outputPath, password string) error {
	// keytool -certreq -alias "hs" -keystore "hs.p12" -storetype pkcs12 -file "hs.csr"
	cmd := exec.Command("keytool",
		"-certreq",
		"-alias", alias,
		"-keystore", keystorePath,
		"-storetype", "pkcs12",
		"-file", outputPath,
	)
	var inputs bytes.Buffer
	inputs.WriteString(password)
	cmd.Stdin = &inputs
	output, err := cmd.CombinedOutput()
	// print the output
	fmt.Println(string(output))
	if err != nil {
		return fmt.Errorf("error generating CSR file: %v, output: %s", err, output)
	}

	return nil
}

func prepareSign() {
	// assign the password to a var
	generateP12File(p12Path, alias, password)
	generateCSRFile(p12Path, alias, csrPath, password)
}

func prepareCert() (string, string) {
	// login()
	// Requests are online operations
	// 1. Request a release cert, prompt the user to input the path of cer file
	certPath, err := promptPath("Input the path of the cer file (Ends with '.cer'): ")
	if err != nil {
		fmt.Println("Error getting cer path: ", err)
		return "", ""
	}
	// https: //developer.huawei.com/consumer/cn/doc/harmonyos-guides-V3/command-line-building-app-hap-0000001193655754-V3
	// 2. Request a release profile
	profilePath, err := promptPath("Input the path of the profile file (Ends with '.p7b'): ")
	if err != nil {
		fmt.Println("Error getting profile path: ", err)
		return "", ""
	}
	return certPath, profilePath
}

func sign(unsignedPath, signedPath, keystorePath, alias, password, profile, cert string) error {
	// downloadFile if not exists
	// err := downloadFile("https://gitee.com/openharmony/signcenter_tool/raw/master/hapsigntool/hapsigntoolv2.jar", "hapsigntoolv2.jar")
	if _, err := os.Stat("hap-sign-tool.jar"); os.IsNotExist(err) {
		err := downloadFile("https://gitee.com/openharmony/developtools_hapsigner/raw/master/dist/hap-sign-tool.jar", "hap-sign-tool.jar")
		if err != nil {
			fmt.Println("Error downloading file: ", err)
			return err
		}
	}
	// https://gitee.com/openharmony/developtools_hapsigner
	// java -jar hap-sign-tool.jar sign-app -keyAlias "oh-app1-key-v1" -signAlg "SHA256withECDSA" -mode "localSign" -appCertFile "result\app1.pem" -profileFile "result\app1-profile.p7b" -inFile "app1-unsigned.zip" -keystoreFile "result\ohtest.jks" -outFile "result\app1-unsigned.hap" -keyPwd "123456" -keystorePwd "123456" -signCode "1"
	cmd := exec.Command("java",
		"-jar", "hap-sign-tool.jar",
		"sign-app", "-keyAlias", alias,
		"-signAlg", "SHA256withECDSA",
		"-mode", "localSign",
		"-appCertFile", cert,
		"-profileFile", profile,
		"-inFile", unsignedPath,
		"-keystoreFile", keystorePath,
		"-outFile", signedPath,
		"-keyPwd", password,
		"-keystorePwd", password,
		"-signCode", "1",
	)
	fmt.Println("Signing file...")
	output, err := cmd.CombinedOutput()
	// print the output
	fmt.Println(string(output))
	if err != nil {
		return fmt.Errorf("error signing file: %v, output: %s", err, output)
	}

	return nil
}
