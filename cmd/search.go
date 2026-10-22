/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var keywords []string
var program string
var name string

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Short: "Search for a command in the commands file",
	Long: `Search for a command in the commands file
Any flag will turn it into "advanced" mode, which is case-sensitive`,
	Use: "pense search [request (optional)]",

	Args: cobra.MaximumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("search called")
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	addCmd.Flags().StringSliceVarP(&keywords, "keywords", "k", []string{}, "search for specific keywords")
	addCmd.Flags().StringVarP(&program, "program", "p", "", "search for a command related to a specific program")
	addCmd.Flags().StringVarP(&name, "name", "n", "", "search for a command with a specific name")
}
