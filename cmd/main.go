package main

import (
	"fmt"
)

func main() {
	// 1. Generate a keys and CSR file
	prepareSign()
	// 2. Create a HarmonyOS app/service in the developer console
	// prompt the user to input the path of Hap file
	localHap, err := promptPath("Input the path of the Hap file (Ends with '.hap'): ")
	if err != nil {
		fmt.Println("Error getting Hap path: ", err)
		return
	}
	// 3. Request a release cert and profile
	cert, profile := prepareCert()
	sign(localHap, signedPath, csrPath, alias, password, profile, cert)
	// install the hap file using installHap
	err = installHap(localHap)
	if err != nil {
		fmt.Println("Error installing Hap: ", err)
		return
	}
}
