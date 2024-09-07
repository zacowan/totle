package fileops

import (
	"os"
)

// CreateDir creates a directory and all necessary parent directories.
// It returns an error if the operation fails.
func CreateDir(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

// CreateFile creates a new file with the given content.
// It returns an error if the file cannot be created or written to.
func CreateFile(path, content string) error {
	return os.WriteFile(path, []byte(content), 0644)
}

// PathExists checks if a file or directory exists.
// It returns true if the path exists, false otherwise.
func PathExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// AppendToFile appends content to an existing file or creates a new one if it doesn't exist.
// It returns an error if the file cannot be opened, created, or written to.
func AppendToFile(path, content string) error {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(content)
	return err
}

// CreateDirectoryIfNotFound creates a directory if it doesn't exist.
// It returns true if the directory was created, false if it already existed.
// If an error occurs during creation, it returns false and the error.
func CreateDirectoryIfNotFound(path string) (bool, error) {
	if !PathExists(path) {
		err := CreateDir(path)
		if err != nil {
			return false, err
		}
		return true, nil
	}
	return false, nil
}
