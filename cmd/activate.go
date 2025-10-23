package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var activateCmd = &cobra.Command{
	Use:   "activate [project-name]",
	Short: "Activate a dbt Cloud project by name",
	Args:  cobra.ExactArgs(1),
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		found := false
		for _, p := range dbtConfig.Projects {
			if strings.EqualFold(p.ProjectName, name) {
				// fmt.Printf("Activated project: %s (ID: %d)\n", p.ProjectName, p.ProjectId)
				name = p.ProjectName
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("Unable to find project %s in dbt_cloud.yml try running list to see all available projects", name)
		}

		dbtCmd := exec.Command("dbt", "cloud", "set", "active-project", "--project-name", name)
		// Connect stdout/stderr so it behaves like a real shell command
		dbtCmd.Stdout = os.Stdout
		dbtCmd.Stderr = os.Stderr

		// Optional: forward stdin too (for prompts)
		dbtCmd.Stdin = os.Stdin

		fmt.Printf("Running: dbt cloud set active-project --project-name %s\n", name)

		if err := dbtCmd.Run(); err != nil {
			return fmt.Errorf("failed to activate project: %w", err)
		}
		return nil
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