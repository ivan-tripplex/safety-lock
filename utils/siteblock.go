package utils

import (
	"bufio"
	"os"
	"runtime"
	"strings"
)

const redirectIP = "127.0.0.1"

var websites = []string{
	"youtube.com",
	"www.youtube.com",
	"m.youtube.com",
	"facebook.com",
	"www.facebook.com",
	"m.facebook.com",
	"twitter.com",
	"vk.com",
}

// Get the path to the hosts file
func getHostPath() string {
	if runtime.GOOS == "windows" {
		return "C:\\Windows\\System32\\drivers\\etc\\hosts"
	} else {
		return "/etc/hosts"
	}
}

// Block Websites
func BlockWebsites() {
	hostPath := getHostPath()

	// Open the hosts file
	file, err := os.OpenFile(hostPath, os.O_RDWR, 0644)
	if err != nil {
		return
	}
	defer file.Close()

	// Read the file and copy in content
	scanner := bufio.NewScanner(file)
	var content []string

	for scanner.Scan() {
		content = append(content, scanner.Text())
	}

	// Block the websites
	for _, site := range websites {
		found := false
		for _, line := range content {
			if strings.Contains(line, site) {
				found = true
				break
			}
		}
		if !found {
			content = append(content, redirectIP+" "+site)
		}
	}

	// Write the updated content back to the hosts file
	file.Truncate(0)
	file.Seek(0, 0)

	writer := bufio.NewWriter(file)
	for _, line := range content {
		writer.WriteString(line + "\n")
	}
	writer.Flush()
}

// Unblock Websites
func UnblockWebsites() {
	hostsPath := getHostPath()

	// Open the hosts file
	file, err := os.OpenFile(hostsPath, os.O_RDWR, 0644)
	if err != nil {
		return
	}
	defer file.Close()

	// Read the file
	scanner := bufio.NewScanner(file)
	var result []string

	//Unblock the websites
	for scanner.Scan() {
		line := scanner.Text()
		blocked := false

		for _, site := range websites {
			if strings.Contains(line, site) {
				blocked = true
				break
			}
		}

		if !blocked {
			result = append(result, line)
		}
	}

	// Write the update result back to the hosts file
	file.Truncate(0)
	file.Seek(0, 0)

	writer := bufio.NewWriter(file)
	for _, line := range result {
		writer.WriteString(line + "\n")
	}
	writer.Flush()
}
