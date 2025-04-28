package file

import (
	"io/ioutil"
	"os"
)

func ReadFile(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func WriteFile(path, content string) error {
	return ioutil.WriteFile(path, []byte(content), 0644)
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
