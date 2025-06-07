package utils_test

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/heuristichiker/rpg-dev/internal/utils"
)

func TestReadLastNBytes(t *testing.T) {
	// Create temporary test file
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test.txt")

	testContent := "Hello, World! This is a test file with some content."
	err := os.WriteFile(testFile, []byte(testContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	tests := []struct {
		name     string
		n        int64
		expected string
	}{
		{"read last 10 bytes", 10, "content."},
		{"read last 20 bytes", 20, "with some content."},
		{"read more than file size", 100, testContent},
		{"read 0 bytes", 0, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := utils.ReadLastNBytes(testFile, tt.n)
			if err != nil {
				t.Fatalf("ReadLastNBytes failed: %v", err)
			}

			if string(result) != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, string(result))
			}
		})
	}
}

func TestReadLastNBytesNonExistentFile(t *testing.T) {
	_, err := utils.ReadLastNBytes("/non/existent/file", 10)
	if err == nil {
		t.Error("Expected error for non-existent file, got nil")
	}
}

func TestGetFileMetadata(t *testing.T) {
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "metadata_test.txt")

	testContent := "Test content for metadata"
	err := os.WriteFile(testFile, []byte(testContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	modTime, size, err := utils.GetFileMetadata(testFile)
	if err != nil {
		t.Fatalf("GetFileMetadata failed: %v", err)
	}

	if size != int64(len(testContent)) {
		t.Errorf("Expected size %d, got %d", len(testContent), size)
	}

	if time.Since(modTime) > time.Minute {
		t.Error("Modification time seems too old for a just-created file")
	}
}

func TestGetFileMetadataNonExistentFile(t *testing.T) {
	_, _, err := utils.GetFileMetadata("/non/existent/file")
	if err == nil {
		t.Error("Expected error for non-existent file, got nil")
	}
}
