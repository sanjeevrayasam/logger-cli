/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package tag

import (
	"fmt"

	"github.com/spf13/cobra"
)

// tagCmd represents the tag command
var TagCmd = &cobra.Command{
	Use:   "tag",
	Short: "Use tag command to perform operation on tags",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("tag called")
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tagCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tagCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
