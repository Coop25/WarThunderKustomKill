package manager

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var relativePaths  = []string{
	`Program Files (x86)\Steam\steamapps\common\War Thunder`,
	`SteamLibrary\steamapps\common\War Thunder`,
	`Steam\steamapps\common\War Thunder`,
	`Games\War Thunder`,
}

func FindWarThunderPath() (string, error) {
	drives := GetAvailableDrives()

	for _, drive := range drives {
		for _, relPath := range relativePaths {
			fullPath := filepath.Join(drive + `\`, relPath)
			if _, err := os.Stat(fullPath); err == nil {
				fmt.Println("Found War Thunder at:", fullPath)
				return fullPath, nil
			}
		}
	}

	// If no automatic match, ask the user
	fmt.Println("Could not automatically find War Thunder install folder.")
	fmt.Print("Please enter the full path to your War Thunder install folder: ")

	reader := bufio.NewReader(os.Stdin)
	userPath, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	userPath = strings.TrimSpace(userPath)
	userPath = filepath.Clean(userPath)

	if _, err := os.Stat(userPath); os.IsNotExist(err) {
		return "", fmt.Errorf("provided path does not exist")
	}

	return userPath, nil
}

func GetAvailableDrives() []string {
	drives := []string{}
	for c := 'C'; c <= 'Z'; c++ { // Start from C:, skip A: and B:
		drive := fmt.Sprintf("%c:", c)
		if _, err := os.Stat(drive + `\`); err == nil {
			drives = append(drives, drive)
		}
	}
	return drives
}
