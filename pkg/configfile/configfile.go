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

type Command struct {
	Name     string `json:"name"`
	Command  string `json:"command"`
	Keywords []string `json:"keywords"`
	Note     string `json:"note"`
}

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

func AddCommand(configPath string, cmdName string, cmdExec string, cmdKeywords []string, cmdNote string) {
	var configs []Command
	fileData, err := os.ReadFile(filepath.Join(configPath, "commands.json"))
	if err == nil {
		_ = json.Unmarshal(fileData, &configs)
	}

	configs = append(configs, Command{
		Name:     cmdName,
		Command:  cmdExec,
		Keywords: cmdKeywords,
		Note:     cmdNote,
	})

	jsonData, err := json.MarshalIndent(configs, "", "   ")
	if err != nil {
		panic(fmt.Errorf("failed to marshal json: %w", err))
	}

	// 0644 gives read/write permissions to the owner, and read-only to others
	err = os.WriteFile(filepath.Join(configPath, "commands.json"), jsonData, 0644)
	if err != nil {
		panic(fmt.Errorf("failed to write file: %w", err))
	}
}
