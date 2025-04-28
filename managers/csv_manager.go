package manager

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type CSVRowTarget struct {
	Key         string
	HumanName   string
	OriginalVal string
}

type CSVManager struct {
	FilePath string
	Lines    []string
}

func NewCSVManager(filePath string) (*CSVManager, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return &CSVManager{
		FilePath: filePath,
		Lines:    lines,
	}, nil
}

func (cm *CSVManager) UpdateSecondColumnInteractive(targets []CSVRowTarget) error {
	reader := bufio.NewReader(os.Stdin)

	for i, line := range cm.Lines {
		fields := strings.Split(line, ";")
		if len(fields) < 2 {
			continue
		}

		key := strings.Trim(fields[0], `"`) // Remove surrounding quotes

		for _, target := range targets {
			if key == target.Key {
				currentValue := strings.Trim(fields[1], `"`)
				fmt.Printf(`Enter new value for "%s" (currently "%s", leave empty to skip): `, target.HumanName, currentValue)
				input, _ := reader.ReadString('\n')
				input = strings.TrimSpace(input)
				if input != "" {
					fields[1] = `"` + input + `"` // Re-wrap updated field
					cm.Lines[i] = strings.Join(fields, ";")
				}
			}
		}
	}
	return nil
}

func (cm *CSVManager) ResetSecondColumnForMatches(targets []CSVRowTarget) error {
	for i, line := range cm.Lines {
		fields := strings.Split(line, ";")
		if len(fields) < 2 {
			continue
		}

		key := strings.Trim(fields[0], `"`) // Remove quotes

		for _, target := range targets {
			if key == target.Key {
				fields[1] = `"` + target.OriginalVal + `"` // Rewrap with quotes
				cm.Lines[i] = strings.Join(fields, ";")
			}
		}
	}
	return nil
}

func (cm *CSVManager) Save() error {
	file, err := os.Create(cm.FilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for _, line := range cm.Lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}
