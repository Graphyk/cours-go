package main

import (
	"github.com/spf13/cobra"
)

var bookCmd = &cobra.Command{
	Use:   "book",
	Short: "Manage book",
}

func init() {
	rootCmd.AddCommand(bookCmd)
}
