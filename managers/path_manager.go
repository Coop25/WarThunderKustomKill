package manager

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var defaultPaths = []string{
	`C:\Program Files (x86)\Steam\steamapps\common\War Thunder`,
	`C:\Games\War Thunder`,
	`D:\Games\War Thunder`,
}

func FindWarThunderPath() (string, error) {
	for _, path := range defaultPaths {
		if _, err := os.Stat(path); err == nil {
			return path, nil
		}
	}

	// Ask user
	fmt.Println("Could not automatically find War Thunder folder.")
	fmt.Print("Please enter the full path to your War Thunder folder: ")

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
