package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var activeCmd = &cobra.Command{
	Use:   "active",
	Short: "Show the active dbt Cloud project",
	RunE: func(cmd *cobra.Command, args []string) error {
		projectName := ""
		for _, project := range dbtConfig.Projects {
			if project.ProjectId == dbtConfig.Context.ActiveProject {
				projectName = project.ProjectName
			}
		}
		fmt.Printf("Active Project: %s\nActive Project ID: %d\n",
			projectName,
			dbtConfig.Context.ActiveProject,
		)
		return nil
	},
}