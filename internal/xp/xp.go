package xp

import (
	"fmt"
	"strings"
)

// XPCommand represents a command that awards XP
type XPCommand struct {
	Prefix string
	XP     int
}

// GetXPCommands returns the list of commands that award XP
func GetXPCommands() []XPCommand {
	return []XPCommand{
		{"docker", 5},
		{"go", 3},
		{"pnpm", 2},
		{"npm", 2},
		{"git", 1},
		{"make", 2},
		{"cargo", 3},
		{"python", 2},
		{"node", 2},
	}
}

// CalculateXPFromCommand calculates XP for a given command
func CalculateXPFromCommand(command string) (int, string) {
	if command == "" {
		return 0, ""
	}

	parts := strings.Fields(command)
	if len(parts) == 0 {
		return 0, ""
	}

	baseCommand := parts[0]
	xpCommands := GetXPCommands()

	for _, xpCmd := range xpCommands {
		if baseCommand == xpCmd.Prefix {
			// Special case for git a-c (add and commit)
			if baseCommand == "git" && len(parts) > 1 && parts[1] == "a-c" {
				return 3, baseCommand + " a-c"
			}
			return xpCmd.XP, baseCommand
		}
	}

	return 0, baseCommand
}

// CalculateTotalXP calculates total XP from a list of commands and provides summary
func CalculateTotalXP(commands []string) int {
	totalXP := 0
	commandSummary := make(map[string]int) // command -> total XP
	commandCounts := make(map[string]int)  // command -> count

	for _, cmd := range commands {
		xpEarned, commandName := CalculateXPFromCommand(cmd)
		if xpEarned > 0 {
			totalXP += xpEarned
			commandSummary[commandName] += xpEarned
			commandCounts[commandName]++
		}
	}

	// Print summary
	fmt.Println("\n📊 XP Summary by Command:")
	for command, totalCommandXP := range commandSummary {
		count := commandCounts[command]
		fmt.Printf("  %s: %d XP (%d uses, avg %.1f XP per use)\n",
			command, totalCommandXP, count, float64(totalCommandXP)/float64(count))
	}
	fmt.Printf("\n💎 Total XP: %d\n", totalXP)

	return totalXP
}
