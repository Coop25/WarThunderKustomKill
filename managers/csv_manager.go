package manager

import (
	"bufio"
	"encoding/csv"
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
	Rows     [][]string
}

func NewCSVManager(filePath string) (*CSVManager, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return &CSVManager{
		FilePath: filePath,
		Rows:     rows,
	}, nil
}

func (cm *CSVManager) UpdateSecondColumnInteractive(targets []CSVRowTarget) error {
	reader := bufio.NewReader(os.Stdin)

	for i, row := range cm.Rows {
		if i == 0 {
			continue // skip header
		}

		if len(row) < 2 {
			continue
		}

		for _, target := range targets {
			if row[0] == target.Key {
				fmt.Printf("Enter new value for \"%s\" (or leave empty to skip): ", target.HumanName)
				input, _ := reader.ReadString('\n')
				input = strings.TrimSpace(input)
				if input != "" {
					cm.Rows[i][1] = input
				}
			}
		}
	}
	return nil
}

func (cm *CSVManager) ResetSecondColumnForMatches(targets []CSVRowTarget) error {
	for i, row := range cm.Rows {
		if i == 0 {
			continue // skip header
		}
		if len(row) < 2 {
			continue
		}

		for _, target := range targets {
			if row[0] == target.Key {
				cm.Rows[i][1] = target.OriginalVal
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

	writer := csv.NewWriter(file)
	defer writer.Flush()

	return writer.WriteAll(cm.Rows)
}
