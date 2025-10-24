package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/hursty1/go_dbt_cli/internal/dbtcloud"
	"github.com/spf13/cobra"
)

var (
	dbtFilePath string
	dbtConfig   dbtcloud.DbtCloudConfig
)

// Root command
var rootCmd = &cobra.Command{
	Use:   "dbtctl",
	Short: "A helper tool for managing dbt Cloud projects",
	Long:  "dbtctl allows you to list, activate, and inspect dbt Cloud projects based on your local ~/.dbt/dbt_cloud.yml file.",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// Locate and load config
		home, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		dbtFilePath = filepath.Join(home, ".dbt", "dbt_cloud.yml")

		if _, err := os.Stat(dbtFilePath); os.IsNotExist(err) {
			return fmt.Errorf("dbt_cloud.yml not found at %s", dbtFilePath)
		}

		dbtConfig, err = dbtcloud.ReadDbtCloudConfig(dbtFilePath)
		if err != nil {
			return fmt.Errorf("failed to load config: %w", err)
		}
		return nil
	},
}

func Execute() {
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(activeCmd)
	rootCmd.AddCommand(activateCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(selectCmd)
	rootCmd.AddCommand(upgradeCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}