package manager

import (
	"fmt"
	"io"
	"os"
	"time"
)

func NewLogManager() (*os.File, error) {
	// Generate the log file name based on the current date
	logFileName := fmt.Sprintf("app-%s.log", time.Now().Format("2006-01-02"))

	// Open or create the log file in append mode
	logFile, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file: %w", err)
	}

	// Create multi-writers for stdout and stderr
	multiOut := io.MultiWriter(os.Stdout, logFile)
	multiErr := io.MultiWriter(os.Stderr, logFile)

	// Set up pipes for capturing output
	stdoutReader, stdoutWriter, _ := os.Pipe()
	stderrReader, stderrWriter, _ := os.Pipe()

	// Redirect stdout and stderr
	os.Stdout = stdoutWriter
	os.Stderr = stderrWriter

	// Read from the pipes and write to the original outputs
	go func() {
		_, _ = io.Copy(multiOut, stdoutReader)
	}()
	go func() {
		_, _ = io.Copy(multiErr, stderrReader)
	}()

	return logFile, nil
}
