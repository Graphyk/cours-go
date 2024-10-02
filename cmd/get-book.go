package main

import (
	"log/slog"

	"github.com/spf13/cobra"
)

var accountFetchIssuesCmd = &cobra.Command{
	Use: "get-book",
	RunE: func(cmd *cobra.Command, args []string) error {
		bookName, err := cmd.Flags().GetString("book-name")
		if err != nil {
			return err
		}

		slog.Info(bookName)
		return nil
	},
}

func init() {
	accountFetchIssuesCmd.Flags().String("book-name", "", "name of the required book")
	accountFetchIssuesCmd.MarkFlagRequired("book-name")
	rootCmd.AddCommand(accountFetchIssuesCmd)
}
