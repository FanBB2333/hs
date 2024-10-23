package main

import (
	"hs/internal"
)

func main() {
	internal.InitBrowser()
}

// func main() {
// 	// 1. Generate a keys and CSR file
// 	internal.PrepareSign()
// 	// 2. Create a HarmonyOS app/service in the developer console
// 	// prompt the user to input the path of Hap file
// 	localHap, err := internal.PromptPath("Input the path of the Hap file (Ends with '.hap'): ")
// 	if err != nil {
// 		fmt.Println("Error getting Hap path: ", err)
// 		return
// 	}
// 	// 3. Request a release cert and profile
// 	certInput, profileInput := internal.PrepareCert()
// 	certInput = internal.DefaultIfEmpty(certInput, internal.Cert)
// 	profileInput = internal.DefaultIfEmpty(profileInput, internal.Profile)
// 	err = internal.SignApp(
// 		localHap,
// 		internal.SignedPath,
// 		internal.CsrPath,
// 		internal.Alias,
// 		internal.Password,
// 		profileInput,
// 		certInput)

// 	if err != nil {
// 		fmt.Println("Error signing Hap: ", err)
// 		return
// 	}
// 	// install the hap file using InstallHap
// 	err = internal.InstallHap(localHap)
// 	if err != nil {
// 		fmt.Println("Error installing Hap: ", err)
// 		return
// 	}
// }
