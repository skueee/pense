package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

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
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
