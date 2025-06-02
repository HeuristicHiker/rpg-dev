package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: rpd <command>")
		return
	}

	switch os.Args[1] {
	case "hi":
		fmt.Println("Hello from RPG dev commands!!")
	case "update":
		buildAndInstall()
	default:
		fmt.Println("Unknown command:", os.Args[1])
	}
}

func buildAndInstall() {
	fmt.Println("🔨 Building rpd...")

	// Get user's home directory to locate ~/bin
	usr, err := user.Current()
	if err != nil {
		fmt.Println("Error getting user:", err)
		return
	}
	binPath := filepath.Join(usr.HomeDir, "bin", "rpd")

	// Build the binary
	cmd := exec.Command("go", "build", "-o", binPath)
	cmd.Dir = "." // current directory
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		fmt.Println("❌ Build failed:", err)
		return
	}

	fmt.Printf("✅ rpd installed to %s\n", binPath)
}
