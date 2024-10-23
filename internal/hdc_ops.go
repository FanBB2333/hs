package internal

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"
)

type Device struct {
	Name     string
	Type     string
	Status   string
	Location string
	Driver   string
}

func executeHDC(commands ...string) ([]string, error) {
	cmd := exec.Command("hdc", commands...)
	output, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	var lines []string
	scanner := bufio.NewScanner(output)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	if err := cmd.Wait(); err != nil {
		return nil, err
	}

	return lines, nil
}

func checkConnection() ([]Device, error) {
	// check the connection using hdc list targets -v
	lines, err := executeHDC("list", "targets", "-v")
	if err != nil {
		return nil, err
	}

	var devices []Device
	for _, line := range lines {
		// read in the output from return
		fields := strings.Fields(line)
		// split the output into fields
		if len(fields) == 5 && fields[2] == "Connected" { //  filter the connected devices
			device := Device{
				Name:     fields[0],
				Type:     fields[1],
				Status:   fields[2],
				Location: fields[3],
				Driver:   fields[4],
			}
			devices = append(devices, device)
		}
	}
	// print the number of connected devices
	fmt.Println("Detected connected devices: ", len(devices))

	return devices, nil
}

func installHap(localHap string) error {
	_, err := executeHDC("install", localHap)
	if err != nil {
		return err
	}
	// return success
	return nil
}

func installHapOld(localHap string) error {
	// prepare the target dir
	tmpDir := generateRandomFileName(8)
	remoteHap := "data/local/tmp/" + tmpDir
	// hdc shell mkdir data/local/tmp/tmpDir
	_, err := executeHDC("shell", "mkdir", remoteHap)
	if err != nil {
		return err
	}
	// send the hap file to the device using hdc install
	// hdc file send localHap "data/local/tmp/tmpDir"
	_, err = executeHDC("file", "send", localHap, remoteHap)
	if err != nil {
		return err
	}
	// install the hap file using hdc install
	// hdc shell bm install -p data/local/tmp/tmpDir
	_, err = executeHDC("shell", "bm", "install", "-p", remoteHap)
	if err != nil {
		return err
	}
	// hdc shell rm -rf data/local/tmp/
	_, err = executeHDC("shell", "rm", "-rf", remoteHap)
	if err != nil {
		return err
	}
	// return success
	return nil
}
