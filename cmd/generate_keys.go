package main

import (
	"fmt"
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
	if err != nil {
		return fmt.Errorf("error generating p12 file: %v, output: %s", err, output)
	}

	return nil
}

func generateCSRFile(keystorePath, alias, outputPath string) error {
	cmd := exec.Command("keytool",
		"-certreq",
		"-alias", alias,
		"-keystore", keystorePath,
		"-storetype", "pkcs12",
		"-file", outputPath,
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error generating CSR file: %v, output: %s", err, output)
	}

	return nil
}

func prepareSign() {
	generateP12File("hs.p12", "hs", "123456")
	generateCSRFile("hs.p12", "hs", "hs.csr")
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
