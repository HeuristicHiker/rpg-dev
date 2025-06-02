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
		fmt.Println("We cannot help at all")
	case "getXp":
		history.LoadHistory()
	default:
		fmt.Println("Unknown command:", os.Args[1])
	}
}
