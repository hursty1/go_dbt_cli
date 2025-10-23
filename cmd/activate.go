package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var activateCmd = &cobra.Command{
	Use:   "activate [project-name]",
	Short: "Activate a dbt Cloud project by name",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		for _, p := range dbtConfig.Projects {
			if strings.EqualFold(p.ProjectName, name) {
				fmt.Printf("✅ Activated project: %s (ID: %d)\n", p.ProjectName, p.ProjectId)
				return nil
			}
		}
		return fmt.Errorf("project not found: %s", name)
	},
	// Tab completion
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		var matches []string
		for _, p := range dbtConfig.Projects {
			if strings.HasPrefix(strings.ToLower(p.ProjectName), strings.ToLower(toComplete)) {
				matches = append(matches, p.ProjectName)
			}
		}
		return matches, cobra.ShellCompDirectiveNoFileComp
	},
}