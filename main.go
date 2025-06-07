package main

import (
	"fmt"
	"os"

	"github.com/heuristichiker/rpg-dev/internal/history"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: rpd <command>")
		return
	}

	switch os.Args[1] {
	case "hi":
		fmt.Println("Hello from RPG Dev!!!")
	case "help":
		fmt.Println("Available commands:")
		fmt.Println("  hi     - Say hello")
		fmt.Println("  xp  - Calculate and display current XP")
		fmt.Println("  xpGet  - Display current XP")
		fmt.Println("  help   - Show this help message")
	case "xp":
		_, err := history.LoadHistory()
		if err != nil {
			fmt.Printf("Error loading history: %v\n", err)
			return
		}
	case "get":
		// Quick XP check without full scan
		currentXP, err := history.GetCurrentXP()
		if err != nil {
			fmt.Printf("Error getting XP: %v\n", err)
			return
		}
		fmt.Printf("💎 Current XP: %d\n", currentXP)
	default:
		fmt.Println("Unknown command:", os.Args[1])
		fmt.Println("Use 'rpd help' for available commands")
	}
}
