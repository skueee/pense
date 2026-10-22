package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/skueee/pense/pkg/configfile"
	"github.com/skueee/pense/pkg/search"
)

var searchKeywordsFlag []string
var searchProgramFlag string
var searchNameFlag string

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Short: "Search for a command in the commands file",
	Long: `Search for a command in the commands file
Any flag will turn it into "advanced" mode, which is case-sensitive`,
	Use: "search [request (optional)]",

	Args: cobra.MaximumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		path, err := configfile.GetConfigPath()
		if err != nil {
			panic(err)
		}
		commands := configfile.GetCommands(path)
		searchResult := search.Search(args[0], commands)

		for _, result := range searchResult {
			fmt.Printf("%s   %s\n", result.Name, result.Command)
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	searchCmd.Flags().StringSliceVarP(&searchKeywordsFlag, "keywords", "k", []string{}, "search for specific keywords")
	searchCmd.Flags().StringVarP(&searchProgramFlag, "program", "p", "", "search for a command related to a specific program")
	searchCmd.Flags().StringVarP(&searchNameFlag, "name", "n", "", "search for a command with a specific name")
}
