package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/skueee/pense/pkg/configfile"
)

var keywords []string
var note string

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [name] [command]",
	Short: "Add a command to the tool",
	Long: `Adds a command to the database.
The command will then be available for later use.`,

	Args:  cobra.ExactArgs(2),

	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		command := args[1]
		fmt.Printf("add called with %s for the name and %s for the command\n", name, command)
		configPath, err := configfile.GetConfigPath()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Config path : %s\n", configPath)
		configfile.AddCommand(configPath, name, command, keywords, note)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	addCmd.Flags().StringSliceVarP(&keywords, "keywords", "k", []string{}, "add keyword (repeatable)")
	addCmd.Flags().StringVarP(&note, "note", "n", "", "add note")
}
