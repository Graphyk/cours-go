package main

import (
	"log/slog"

	"github.com/spf13/cobra"
)

var getBookCmd = &cobra.Command{
	Use: "get-book",
	RunE: func(cmd *cobra.Command, args []string) error {
		db, cfg, err := config.ReadConfig(cmd)
		if err != nil {
			return err
		}
		bookName, err := cmd.Flags().GetString("book-name")
		if err != nil {
			return err
		}

		slog.Info(bookName)
		return nil
	},
}

func init() {
	getBookCmd.Flags().String("book-name", "", "name of the required book")
	getBookCmd.Flags().String("config", "", "path to the config file")
	getBookCmd.MarkFlagRequired("book-name")
	getBookCmd.MarkFlagRequired("config")

	rootCmd.AddCommand(getBookCmd)
}
