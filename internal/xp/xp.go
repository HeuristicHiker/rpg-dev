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
func CalculateXPFromCommand(command string) int {
	commandValue := make(map[string]int)

	if command == "" {
		return 0
	}

	parts := strings.Fields(command)
	if len(parts) == 0 {
		return 0
	}

	baseCommand := parts[0]
	xpCommands := GetXPCommands()

	for _, xpCmd := range xpCommands {
		if baseCommand == xpCmd.Prefix {
			// Special case for git a-c (add and commit)
			fmt.Println("Additional command", baseCommand)
			commandValue[baseCommand] = commandValue[baseCommand] + 1

			if baseCommand == "git" && len(parts) > 1 && parts[1] == "a-c" {
				return 3 // Higher XP for git commits
			}
			return xpCmd.XP
		}
	}

	// fmt.Println(commandValue)

	return 0
}

// CalculateTotalXP calculates total XP from a list of commands
func CalculateTotalXP(commands []string) int {
	totalXP := 0
	for _, cmd := range commands {
		totalXP += CalculateXPFromCommand(cmd)
	}
	return totalXP
}
