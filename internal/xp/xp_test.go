package xp_test

import (
	"testing"

	"github.com/heuristichiker/rpg-dev/internal/xp"
)

func TestGetXPCommands(t *testing.T) {
	commands := xp.GetXPCommands()

	if len(commands) == 0 {
		t.Error("Expected non-empty list of XP commands")
	}

	// Check for expected commands
	expectedCommands := map[string]int{
		"docker": 5,
		"go":     3,
		"git":    1,
		"pnpm":   2,
		"npm":    2,
		"make":   2,
		"cargo":  3,
		"python": 2,
		"node":   2,
	}

	commandMap := make(map[string]int)
	for _, cmd := range commands {
		commandMap[cmd.Prefix] = cmd.XP
	}

	for prefix, expectedXP := range expectedCommands {
		if actualXP, exists := commandMap[prefix]; !exists {
			t.Errorf("Expected command %s not found", prefix)
		} else if actualXP != expectedXP {
			t.Errorf("Command %s: expected %d XP, got %d XP", prefix, expectedXP, actualXP)
		}
	}
}

func TestCalculateXPFromCommand(t *testing.T) {
	tests := []struct {
		name            string
		command         string
		expectedXP      int
		expectedCommand string
	}{
		{"docker command", "docker run -p 8080:80 nginx", 5, "docker"},
		{"go command", "go build main.go", 3, "go"},
		{"git command", "git commit -m 'test'", 1, "git"},
		{"git a-c special case", "git a-c", 3, "git a-c"},
		{"pnpm command", "pnpm install", 2, "pnpm"},
		{"npm command", "npm run dev", 2, "npm"},
		{"make command", "make build", 2, "make"},
		{"cargo command", "cargo build", 3, "cargo"},
		{"python command", "python app.py", 2, "python"},
		{"node command", "node server.js", 2, "node"},
		{"unknown command", "ls -la", 0, "ls"},
		{"empty command", "", 0, ""},
		{"whitespace only", "   ", 0, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			xp, cmd := xp.CalculateXPFromCommand(tt.command)
			if xp != tt.expectedXP {
				t.Errorf("Expected XP %d, got %d", tt.expectedXP, xp)
			}
			if cmd != tt.expectedCommand {
				t.Errorf("Expected command %q, got %q", tt.expectedCommand, cmd)
			}
		})
	}
}

func TestCalculateTotalXP(t *testing.T) {
	commands := []string{
		"docker run nginx",
		"go build",
		"git commit -m 'test'",
		"git a-c",
		"ls -la",
		"npm install",
		"unknown command",
	}

	// Expected: docker(5) + go(3) + git(1) + git a-c(3) + npm(2) = 14
	expectedTotal := 14

	total := xp.CalculateTotalXP(commands)
	if total != expectedTotal {
		t.Errorf("Expected total XP %d, got %d", expectedTotal, total)
	}
}
