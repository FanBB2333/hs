package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"runtime"
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
	openBrowser("https://developer.huawei.com/consumer/cn/service/josp/agc/index.html")
}

func generateRandomFileName(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	fileName := make([]byte, length)

	for i := range fileName {
		fileName[i] = charset[r.Intn(len(charset))]
	}

	return string(fileName)
}

func promptPath(prompt string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)

	path, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	// drop the newline character
	return path[:len(path)-1], nil
}

func downloadFile(url string, filepath string) error {
	// Create HTTP request
	response, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("unable to fetch file: %v", err)
	}
	defer response.Body.Close()

	// Check HTTP response status
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("download failed: %s", response.Status)
	}

	// Create file
	outFile, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("unable to create file: %v", err)
	}
	defer outFile.Close()

	// Write response content to file
	_, err = io.Copy(outFile, response.Body)
	if err != nil {
		return fmt.Errorf("error writing to file: %v", err)
	}

	return nil
}
