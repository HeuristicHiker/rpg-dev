package xp_test

import (
	"testing"

	"github.com/heuristichiker/rpg-dev/internal/xp"
)

func TestXPCalculationIntegration(t *testing.T) {
	// Test realistic command sequences
	commands := []string{
		"docker-compose up -d",
		"go mod tidy",
		"git add .",
		"git commit -m 'feature: add new functionality'",
		"git push origin main",
		"npm run build",
		"make deploy",
	}

	totalXP := xp.CalculateTotalXP(commands)
	expectedXP := 5 + 3 + 1 + 1 + 1 + 2 + 2 // 15 total

	if totalXP != expectedXP {
		t.Errorf("Expected total XP %d, got %d", expectedXP, totalXP)
	}
}
