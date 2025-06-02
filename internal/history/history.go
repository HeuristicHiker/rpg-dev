package history

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/heuristichiker/rpg-dev/internal/utils"
)

type History struct {
	Commands []string
}

type State struct {
	LastChecked time.Time `json:"last_checked"`
	LastSize    int64     `json:"last_size"`
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

	fmt.Printf("History file last modified: %v\n", modTime)
	fmt.Printf("Bytes added since last check: %d\n", size-previousSize)

	const lastBytes = 10 * 1024 * 1024
	data, err := utils.ReadLastNBytes(historyPath, lastBytes)
	if err != nil {
		return nil, err
	}

	commands := []string{}
	scanner := bufio.NewScanner(strings.NewReader(string(data)))
	cutoff := time.Now().Add(-12 * time.Hour)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 || line[0] != ':' {
			continue
		}
		parts := strings.SplitN(line, ";", 2)
		if len(parts) != 2 {
			continue
		}

		meta := parts[0]
		cmd := parts[1]

		metaParts := strings.Split(meta, ":")
		if len(metaParts) < 2 {
			continue
		}
		tsStr := strings.TrimSpace(metaParts[1])
		tsInt, err := strconv.ParseInt(tsStr, 10, 64)
		if err != nil {
			continue
		}
		timestamp := time.Unix(tsInt, 0)
		if timestamp.After(cutoff) {
			commands = append(commands, cmd)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	state.LastChecked = time.Now()
	state.LastSize = size
	if err := saveState(statePath, state); err != nil {
		return nil, err
	}

	return &History{Commands: commands}, nil
}
