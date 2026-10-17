/*
 Operations related to config files
*/

package configfile

import (
	"os"
	"path/filepath"
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
