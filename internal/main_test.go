// main_test.go
package internal

import (
	"fmt"
	"os"
	"testing"
)

func TestGetEnvVars(t *testing.T) {
	envVars := os.Environ()

	if len(envVars) == 0 {
		t.Error("Expected non-empty environment variables")
	}

	for _, env := range envVars {
		t.Log(env)
	}
}

func TestConn(t *testing.T) {
	devices, err := checkConnection()
	if err != nil {
		fmt.Println("Error checking connection: ", err)
		return
	}

	// print the connected devices
	for _, device := range devices {
		fmt.Println(device)
	}
}

func TestMain(t *testing.T) {
	// 1. Generate a keys and CSR file
	PrepareSign()
	// 2. Create a HarmonyOS app/service in the developer console
	// prompt the user to input the path of Hap file
	localHap, err := PromptPath("Input the path of the Hap file (Ends with '.hap'): ")
	if err != nil {
		fmt.Println("Error getting Hap path: ", err)
		return
	}
	Login()
	// 3. Request a release cert and profile
	PrepareCert()

	// install the hap file using InstallHap
	err = InstallHap(localHap)
	if err != nil {
		fmt.Println("Error installing Hap: ", err)
		return
	}
}

func TestRequest(t *testing.T) {
	Login()
}
