/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/ylanzinhoy/go-scaffold/internal"
)

// maximumScaffoldCmd represents the maximumScaffold command
var maximumScaffoldCmd = &cobra.Command{
	Use:   "maximumScaffold",
	Short: "Generate a Maximum files for your Go Application",
	Run: func(cmd *cobra.Command, args []string) {
		internal.MaximumFiles()
	},
}

func init() {
	rootCmd.AddCommand(maximumScaffoldCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// maximumScaffoldCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// maximumScaffoldCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
