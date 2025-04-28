package manager

import (
	"strings"

	"github.com/Coop25/WarThunderKustomKill/accessor"
)

type ConfigManager struct {
	FilePath string
	Content  string
}

func NewConfigManager(filePath string) (*ConfigManager, error) {
	content, err := accessor.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return &ConfigManager{
		FilePath: filePath,
		Content:  content,
	}, nil
}

func (cm *ConfigManager) InsertAfterSection(sectionMarker, newLine string) error {
	lines := strings.Split(cm.Content, "\n")
	var newLines []string
	inserted := false

	for _, line := range lines {
		newLines = append(newLines, line)
		if strings.Contains(line, sectionMarker) && !inserted {
			newLines = append(newLines, newLine)
			inserted = true
		}
	}

	cm.Content = strings.Join(newLines, "\n")
	return nil
}

func (cm *ConfigManager) RemoveInsertedLine(markerText string) error {
	lines := strings.Split(cm.Content, "\n")
	var newLines []string

	for _, line := range lines {
		if strings.TrimSpace(line) == markerText {
			continue
		}
		newLines = append(newLines, line)
	}

	cm.Content = strings.Join(newLines, "\n")
	return nil
}

func (cm *ConfigManager) Save() error {
	return accessor.WriteFile(cm.FilePath, cm.Content)
}
