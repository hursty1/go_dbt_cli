package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var activeCmd = &cobra.Command{
	Use:   "active",
	Short: "Show the active dbt Cloud project",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("Active Host: %s\nActive Project ID: %d\n",
			dbtConfig.Context.ActiveHost,
			dbtConfig.Context.ActiveProject,
		)
		return nil
	},
}