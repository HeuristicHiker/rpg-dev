package history_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/heuristichiker/rpg-dev/internal/history"
)

func TestGetCurrentXP(t *testing.T) {
	tmpDir := t.TempDir()

	// Create a mock state file
	statePath := filepath.Join(tmpDir, ".rpd_state.json")
	state := struct {
		LastXP int `json:"last_xp"`
	}{LastXP: 42}

	file, err := os.Create(statePath)
	if err != nil {
		t.Fatalf("Failed to create state file: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(state); err != nil {
		t.Fatalf("Failed to encode state: %v", err)
	}

	// Temporarily change home directory for test
	originalHome := os.Getenv("HOME")
	os.Setenv("HOME", tmpDir)
	defer os.Setenv("HOME", originalHome)

	xp, err := history.GetCurrentXP()
	if err != nil {
		t.Fatalf("GetCurrentXP failed: %v", err)
	}

	if xp != 42 {
		t.Errorf("Expected XP 42, got %d", xp)
	}
}

func TestGetCurrentXPNoStateFile(t *testing.T) {
	tmpDir := t.TempDir()

	// Temporarily change home directory for test
	originalHome := os.Getenv("HOME")
	os.Setenv("HOME", tmpDir)
	defer os.Setenv("HOME", originalHome)

	xp, err := history.GetCurrentXP()
	if err != nil {
		t.Fatalf("GetCurrentXP failed: %v", err)
	}

	// Should return 0 for non-existent state file
	if xp != 0 {
		t.Errorf("Expected XP 0 for missing state file, got %d", xp)
	}
}
