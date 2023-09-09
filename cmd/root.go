/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/sanjeevrayasam/logger-cli/cmd/add"
	"github.com/sanjeevrayasam/logger-cli/cmd/board"
	"github.com/sanjeevrayasam/logger-cli/cmd/list"
	"github.com/sanjeevrayasam/logger-cli/cmd/tag"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "log",
	Short: "Log the day routine from your terminal",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
func addCommands() {
	rootCmd.AddCommand(tag.TagCmd)
	rootCmd.AddCommand(add.AddCmd)
	rootCmd.AddCommand(board.BoardCmd)
	rootCmd.AddCommand(list.ListCmd)

}

func init() {
	addCommands()
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.logger-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
