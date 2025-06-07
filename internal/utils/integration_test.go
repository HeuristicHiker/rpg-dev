package integration_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/heuristichiker/rpg-dev/internal/history"
	"github.com/heuristichiker/rpg-dev/internal/utils"
)

func TestHistoryIntegration(t *testing.T) {
	tmpDir := t.TempDir()

	// Create a mock zsh history file
	historyPath := filepath.Join(tmpDir, ".zsh_history")
	historyContent := `: 1640995200:0;docker run nginx
: 1640995300:0;go build
: 1640995400:0;git commit -m "test"
: 1640995500:0;npm install
: 1640995600:0;unknown command
`

	err := os.WriteFile(historyPath, []byte(historyContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create history file: %v", err)
	}

	// Test file utilities
	modTime, size, err := utils.GetFileMetadata(historyPath)
	if err != nil {
		t.Fatalf("GetFileMetadata failed: %v", err)
	}

	if size != int64(len(historyContent)) {
		t.Errorf("Expected size %d, got %d", len(historyContent), size)
	}

	// Read file content
	data, err := utils.ReadLastNBytes(historyPath, 1024)
	if err != nil {
		t.Fatalf("ReadLastNBytes failed: %v", err)
	}

	if string(data) != historyContent {
		t.Error("File content doesn't match expected content")
	}

	t.Logf("Integration test completed successfully. File size: %d, mod time: %v", size, modTime)
}

func TestStateManagement(t *testing.T) {
	tmpDir := t.TempDir()

	// Temporarily change home directory
	originalHome := os.Getenv("HOME")
	os.Setenv("HOME", tmpDir)
	defer os.Setenv("HOME", originalHome)

	// First call should return 0 (no state file)
	xp1, err := history.GetCurrentXP()
	if err != nil {
		t.Fatalf("First GetCurrentXP failed: %v", err)
	}

	if xp1 != 0 {
		t.Errorf("Expected initial XP 0, got %d", xp1)
	}

	// Create mock history file
	historyPath := filepath.Join(tmpDir, ".zsh_history")
	historyContent := `: 1640995200:0;docker run nginx
: 1640995300:0;go build
`

	err = os.WriteFile(historyPath, []byte(historyContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create history file: %v", err)
	}

	// Load history should create state file
	_, err = history.LoadHistory()
	if err != nil {
		t.Fatalf("LoadHistory failed: %v", err)
	}

	// Second call should return non-zero XP
	xp2, err := history.GetCurrentXP()
	if err != nil {
		t.Fatalf("Second GetCurrentXP failed: %v", err)
	}

	// docker(5) + go(3) = 8
	expectedXP := 8
	if xp2 != expectedXP {
		t.Errorf("Expected XP %d, got %d", expectedXP, xp2)
	}
}
