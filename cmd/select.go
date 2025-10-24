package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var selectCmd = &cobra.Command{
	Use:   "select",
	Short: "Select a project to activate",
	RunE: func(cmd *cobra.Command, args []string) error {
		selectedProject, err := selectProject()
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		err = RunActivateProjectCommand(selectedProject)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		return nil
	},
}


func selectProject() (string, error) {
	if len(dbtConfig.Projects) == 0 {
		return "", fmt.Errorf("No Projects configurations found.")
	}

	keys := make([]string, 0, len(dbtConfig.Projects))
	for _, k := range dbtConfig.Projects {
		keys = append(keys, k.ProjectName)
	}

	prompt := promptui.Select{
		Label: "Select a device to open connection",
		Items: keys,
		Size: 10,
	}

	_, key, err := prompt.Run()
	if err != nil {
		return "", fmt.Errorf("prompt failed: %w", err)
	}

	return key, nil
}


func RunActivateProjectCommand(name string) error {
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
}