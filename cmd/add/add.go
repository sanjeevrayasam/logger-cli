/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package add

import (
	"fmt"

	"github.com/spf13/cobra"
)

// add/addCmd represents the add/add command
var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new entry of log into the system",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add/add called")
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// add/addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// add/addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
