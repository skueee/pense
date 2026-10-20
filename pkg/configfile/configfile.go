/*
 Operations related to config files
*/

package configfile

import (
	"fmt"
	"os"
	"path/filepath"
	"encoding/json"
)

// Gets the config path from the debug environment variable or the common path on the user's system
func GetConfigPath() (string, error) {
	configEnv := os.Getenv("PENSE_PATH")
	if configEnv != "" {
		return configEnv, nil
	} else {
		dir, err := os.UserConfigDir()
		if err != nil {
			return "", err
		}
		return filepath.Join(dir, "pense"), nil
	}
}

// A command in the JSON commands file
type Command struct {
	Name     string `json:"name"`
	Command  string `json:"command"`
	Keywords []string `json:"keywords"`
	Note     string `json:"note"`
}

// Gets the commands present in the json
func GetCommands(configPath string) []Command {
	file, err := os.ReadFile(filepath.Join(configPath, "commands.json"))
	if err != nil {
		panic(err)
	}

	var commands []Command
		err = json.Unmarshal(file, &commands)
		if err != nil {
			panic(err)
		}

	return commands
}

// Adds a command, after checking if the command can be added
func AddCommand(configPath string, cmdName string, cmdExec string, cmdKeywords []string, cmdNote string) error {
	var configs []Command
	fileData, err := os.ReadFile(filepath.Join(configPath, "commands.json"))
	if err == nil {
		_ = json.Unmarshal(fileData, &configs)
	}

	// Iterates over every command to check if a command with a similar name already exists
	for i:= 0; i < len(configs); i++ {
		if configs[i].Name == cmdName {
			return fmt.Errorf("a command can't have the same name as an already existing command")
		}
	}

	configs = append(configs, Command{
		Name:     cmdName,
		Command:  cmdExec,
		Keywords: cmdKeywords,
		Note:     cmdNote,
	})

	jsonData, err := json.MarshalIndent(configs, "", "   ")
	if err != nil {
		return fmt.Errorf("failed to marshal json: %w", err)
	}

	// 0644 gives read/write permissions to the owner, and read-only to others
	err = os.WriteFile(filepath.Join(configPath, "commands.json"), jsonData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	// No error, nice :)
	return nil
}
