package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all configured dbt Cloud projects",
	RunE: func(cmd *cobra.Command, args []string) error {
		for _, p := range dbtConfig.Projects {
			fmt.Printf("- %s (Project ID: %d, Account: %s)\n", p.ProjectName, p.ProjectId, p.AccountName)
		}
		return nil
	},
}