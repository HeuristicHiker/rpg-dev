package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/heuristichiker/rpg-dev/internal/history"
)

func TestXPCommand(t *testing.T) {
	// Create temporary directory for test files
	tempDir, err := ioutil.TempDir("", "rpd_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create mock .zsh_history file
	mockHistory := `1733590800 go run main.go
1733590801 git a-c -m "test commit"
1733590802 docker ps
1733590803 npm install
1733590804 make build
1733590805 python script.py
1733590806 ls
1733590807 go test
1733590808 git push`

	historyPath := filepath.Join(tempDir, ".zsh_history")
	err = ioutil.WriteFile(historyPath, []byte(mockHistory), 0644)
	if err != nil {
		t.Fatalf("Failed to create mock history file: %v", err)
	}

	// Create mock .rpd_state.json file
	mockState := history.State{
		LastChecked: "2025-01-01T00:00:00Z",
		LastSize:    0,
		LastXP:      0,
	}
	stateData, _ := json.MarshalIndent(mockState, "", "  ")
	statePath := filepath.Join(tempDir, ".rpd_state.json")
	err = ioutil.WriteFile(statePath, stateData, 0644)
	if err != nil {
		t.Fatalf("Failed to create mock state file: %v", err)
	}

	// Test LoadHistory with mock paths
	historyResult, err := history.LoadHistoryWithPaths(historyPath, statePath)
	if err != nil {
		t.Fatalf("LoadHistory failed: %v", err)
	}

	// Assertions
	if historyResult == nil {
		t.Fatal("Expected history result, got nil")
	}

	if historyResult.TotalXP <= 0 {
		t.Errorf("Expected positive XP, got %d", historyResult.TotalXP)
	}

	// Expected XP calculation:
	// go run main.go: 3 XP
	// git a-c: 3 XP (special case)
	// docker ps: 5 XP
	// npm install: 2 XP
	// make build: 2 XP
	// python script.py: 2 XP
	// go test: 3 XP
	// Total: 20 XP
	expectedXP := 20
	if historyResult.TotalXP != expectedXP {
		t.Errorf("Expected XP %d, got %d", expectedXP, historyResult.TotalXP)
	}

	t.Logf("Test passed! Total XP: %d", historyResult.TotalXP)
}
