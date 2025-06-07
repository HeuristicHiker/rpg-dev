package history

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/heuristichiker/rpg-dev/internal/utils"
	"github.com/heuristichiker/rpg-dev/internal/xp"
)

type History struct {
	Commands []string
	TotalXP  int
}

type State struct {
	LastChecked time.Time `json:"last_checked"`
	LastSize    int64     `json:"last_size"`
	LastXP      int       `json:"last_xp"`
}

func loadState(path string) (*State, error) {
	file, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return &State{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var state State
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&state); err != nil {
		return nil, err
	}
	return &state, nil
}

func saveState(path string, state *State) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(state)
}

func LoadHistory() (*History, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	historyPath := filepath.Join(homeDir, ".zsh_history")
	statePath := filepath.Join(homeDir, ".rpd_state.json")

	state, err := loadState(statePath)
	if err != nil {
		return nil, err
	}

	modTime, size, err := utils.GetFileMetadata(historyPath)
	if err != nil {
		return nil, err
	}

	previousSize := state.LastSize
	previousXP := state.LastXP

	fmt.Printf("History file last modified: %v\n", modTime)
	fmt.Printf("Bytes added since last check: %d\n", size-previousSize)

	const lastBytes = 10 * 1024 * 1024
	data, err := utils.ReadLastNBytes(historyPath, lastBytes)
	if err != nil {
		fmt.Printf("Error reading history file\n")
		return nil, err
	}

	commands := []string{}
	scanner := bufio.NewScanner(strings.NewReader(string(data)))
	currentXP := 0

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		parts := strings.SplitN(line, " ", 2)
		if len(parts) != 2 {
			continue
		}

		cmd := parts[1]

		commands = append(commands, cmd)
		cmdXP := xp.CalculateXPFromCommand(cmd)
		currentXP += cmdXP

		if cmdXP > 0 {
			fmt.Printf("🏅 %s -> +%d XP\n", cmd, cmdXP)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error with scanner")
		return nil, err
	}

	fmt.Printf("💎 Total XP: %d\n", currentXP)
	fmt.Printf("📈 XP gained since last check: %d\n", currentXP-previousXP)

	state.LastChecked = time.Now()
	state.LastSize = size
	state.LastXP = currentXP

	if err := saveState(statePath, state); err != nil {
		return nil, err
	}

	return &History{Commands: commands, TotalXP: currentXP}, nil
}

func GetCurrentXP() (int, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return 0, err
	}

	statePath := filepath.Join(homeDir, ".rpd_state.json")
	state, err := loadState(statePath)
	if err != nil {
		return 0, err
	}

	return state.LastXP, nil
}

func CategorizeCommands(args []string) {

}
