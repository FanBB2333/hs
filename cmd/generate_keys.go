package main

import (
	"fmt"
	"bytes"
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
	password := "12345678"
	p12Path := "hs.p12"
	csrPath := "hs.csr"
	alias := "hs"
	generateP12File(p12Path, alias, password)
	generateCSRFile(p12Path, alias, csrPath, password)
}

func prepareCert() {
	// 1. Request a release cert

	// 2. Request a release profile
}

func main() {
	// 1. Generate a keys and CSR file
	prepareSign()
	// 2. Create a HarmonyOS app/service in the developer console
	// 3. Request a release cert and profile
	prepareCert()
}
