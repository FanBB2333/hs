package internal

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"time"
)

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

func PromptPath(prompt string) (string, error) {
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

func DefaultIfEmpty(input, defaultValue string) string {
	if input == "" {
		return defaultValue
	}
	return input
}