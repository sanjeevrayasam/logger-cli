/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package board

import (
	"fmt"

	"github.com/spf13/cobra"
)

// board/boardCmd represents the board/board command
var BoardCmd = &cobra.Command{
	Use:   "board",
	Short: "To get the dashboard serving at localhost to visualize the logs",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("board/board called")
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// board/boardCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// board/boardCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
