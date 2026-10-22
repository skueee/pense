package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/skueee/pense/pkg/configfile"
)

var addKeywordsFlag []string
var addNoteFlag string

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

		configPath, err := configfile.GetConfigPath()
		if err != nil {
			log.Fatal(err)
		}

		err = configfile.AddCommand(configPath, name, command, addKeywordsFlag, addNoteFlag)
		if err != nil {
			log.Fatalf("Error: %v", err)
		} else {
			fmt.Printf("Saved command !")
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	addCmd.Flags().StringSliceVarP(&addKeywordsFlag, "keywords", "k", []string{}, "add keyword (repeatable)")
	addCmd.Flags().StringVarP(&addNoteFlag, "note", "n", "", "add note")
}
