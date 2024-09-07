package fileops

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCreateDir(t *testing.T) {
	testDir := filepath.Join(os.TempDir(), "test_create_dir")
	defer os.RemoveAll(testDir)

	err := CreateDir(testDir)
	if err != nil {
		t.Fatalf("CreateDir failed: %v", err)
	}

	if !PathExists(testDir) {
		t.Errorf("Directory was not created")
	}
}

func TestCreateFile(t *testing.T) {
	testFile := filepath.Join(os.TempDir(), "test_create_file.txt")
	defer os.Remove(testFile)

	content := "Test content"
	err := CreateFile(testFile, content)
	if err != nil {
		t.Fatalf("CreateFile failed: %v", err)
	}

	if !PathExists(testFile) {
		t.Errorf("File was not created")
	}

	data, err := os.ReadFile(testFile)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	if string(data) != content {
		t.Errorf("File content does not match. Expected: %s, Got: %s", content, string(data))
	}
}

func TestCreateFileInNonExistentDir(t *testing.T) {
	nonExistentDir := filepath.Join(os.TempDir(), "non_existent_dir")
	nonExistentFile := filepath.Join(nonExistentDir, "test.txt")
	defer os.RemoveAll(nonExistentDir)

	err := CreateFile(nonExistentFile, "content")
	if err == nil {
		t.Errorf("CreateFile should fail when creating file in non-existent directory")
	}
}

func TestPathExists(t *testing.T) {
	existingFile := filepath.Join(os.TempDir(), "existing_file.txt")
	err := os.WriteFile(existingFile, []byte("test"), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove(existingFile)

	if !PathExists(existingFile) {
		t.Errorf("PathExists returned false for an existing file")
	}
}

func TestPathDoesNotExist(t *testing.T) {
	nonExistingFile := filepath.Join(os.TempDir(), "non_existing_file.txt")

	if PathExists(nonExistingFile) {
		t.Errorf("PathExists returned true for a non-existing file")
	}
}

func TestAppendToFile(t *testing.T) {
	testFile := filepath.Join(os.TempDir(), "test_append_file.txt")
	defer os.Remove(testFile)

	initialContent := "Initial content\n"
	appendedContent := "Appended content"

	err := CreateFile(testFile, initialContent)
	if err != nil {
		t.Fatalf("Failed to create initial file: %v", err)
	}

	err = AppendToFile(testFile, appendedContent)
	if err != nil {
		t.Fatalf("AppendToFile failed: %v", err)
	}

	data, err := os.ReadFile(testFile)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	expectedContent := initialContent + appendedContent
	if string(data) != expectedContent {
		t.Errorf("File content does not match. Expected: %s, Got: %s", expectedContent, string(data))
	}
}

func TestAppendToReadOnlyFile(t *testing.T) {
	readOnlyFile := filepath.Join(os.TempDir(), "readonly_file.txt")
	err := os.WriteFile(readOnlyFile, []byte("Read-only content"), 0444)
	if err != nil {
		t.Fatalf("Failed to create read-only file: %v", err)
	}
	defer os.Remove(readOnlyFile)

	err = AppendToFile(readOnlyFile, "Appended content")
	if err == nil {
		t.Errorf("AppendToFile should fail when appending to a read-only file")
	}
}

func TestCreateDirectoryIfNotFound(t *testing.T) {
	testDir := filepath.Join(os.TempDir(), "test_create_dir_if_not_found")
	defer os.RemoveAll(testDir)

	created, err := CreateDirectoryIfNotFound(testDir)
	if err != nil {
		t.Fatalf("CreateDirectoryIfNotFound failed: %v", err)
	}
	if !created {
		t.Errorf("CreateDirectoryIfNotFound returned false for a new directory")
	}
	if !PathExists(testDir) {
		t.Errorf("Directory was not created")
	}
}

func TestCreateExistingDirectory(t *testing.T) {
	testDir := filepath.Join(os.TempDir(), "test_create_existing_dir")
	err := os.Mkdir(testDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create test directory: %v", err)
	}
	defer os.RemoveAll(testDir)

	created, err := CreateDirectoryIfNotFound(testDir)
	if err != nil {
		t.Fatalf("CreateDirectoryIfNotFound failed for existing directory: %v", err)
	}
	if created {
		t.Errorf("CreateDirectoryIfNotFound returned true for an existing directory")
	}
}
