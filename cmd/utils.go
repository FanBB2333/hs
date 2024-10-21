package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"math/rand"
	"time"
)

func openBrowser(url string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	case "darwin": // macOS
		cmd = exec.Command("open", url)
	case "linux":
		cmd = exec.Command("xdg-open", url)
	default:
		return fmt.Errorf("unsupported platform")
	}

	return cmd.Start() // start the command
}

func login() {
	openBrowser("https://developer.huawei.com/consumer/en/console")
}

func generateRandomFileName(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano()) // time seed
	fileName := make([]byte, length)

	for i := range fileName {
		fileName[i] = charset[rand.Intn(len(charset))]
	}

	return string(fileName)
}

func main() {
	login()
}
