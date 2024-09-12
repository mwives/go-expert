/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/mwives/go-expert/cli/internal/database"
	"github.com/spf13/cobra"
)

func newCreateCmd(categoryDb database.Category) *cobra.Command {
	return &cobra.Command{
		Use:   "create",
		Short: "Create a category",
		Long:  "Create a category is a tool to create a category.",
		RunE:  runCreate(categoryDb),
	}
}

func runCreate(categoryDb database.Category) RunEFunc {
	return func(cmd *cobra.Command, args []string) error {
		name, _ := cmd.Flags().GetString("name")
		description, _ := cmd.Flags().GetString("description")

		_, err := categoryDb.Create(name, description)
		return err
	}
}

func init() {
	createCmd := newCreateCmd(GetCategoryDB(OpenDatabase()))

	categoryCmd.AddCommand(createCmd)

	createCmd.Flags().StringP("name", "n", "", "Name of the category")
	createCmd.Flags().StringP("description", "d", "", "Description of the category")

	createCmd.MarkFlagRequired("name")
	createCmd.MarkFlagsRequiredTogether("name", "description")
}
