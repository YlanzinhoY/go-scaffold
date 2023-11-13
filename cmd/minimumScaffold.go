/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/ylanzinhoy/go-scaffold/internal"

	"github.com/spf13/cobra"
)

// minimumScaffoldCmd represents the minimumScaffold command
var minimumScaffoldCmd = &cobra.Command{
	Use:   "minimumScaffold",
	Short: "Generate a minimal folder scaffold for your application Go",
	Run: func(cmd *cobra.Command, args []string) {
		internal.MinimumFiles()
	},
}

func init() {
	rootCmd.AddCommand(minimumScaffoldCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// minimalScaffoldCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// minimalScaffoldCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
