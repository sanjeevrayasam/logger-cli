/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package list

import (
	"fmt"

	"github.com/spf13/cobra"
)

// list/listCmd represents the list/list command
var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List command helps to get the list of logs",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list/list called")
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// list/listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// list/listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
