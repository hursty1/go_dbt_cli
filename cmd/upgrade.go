package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"

	"github.com/hursty1/go_dbt_cli/internal/version"
	"github.com/spf13/cobra"
)

var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Installs the latest version of the CLI.",
	Run: func(cmd *cobra.Command, args []string) {
		if !IsOutdated(version.Get()) {
			fmt.Println("dbthelper is up to date.")
			return
		}

		fmt.Println("Upgrading dbthelper to the latest version...")
		command := exec.Command("go", "install", "github.com/hursty1/go_dbt_cli/cmd/dbthelper@latest")
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr

		if err := command.Run(); err != nil {
			fmt.Printf("Upgrade failed: %v\n", err)
			return
		}

		fmt.Println("Upgrade completed. Verifying version...")

		command = exec.Command("dbthelper", "version")
		b, err := command.Output()
		if err != nil {
			fmt.Println("Error checking new version:", err)
			return
		}

		re := regexp.MustCompile(`v\d+\.\d+\.\d+`)
		v := re.FindString(string(b))
		fmt.Printf("Successfully upgraded to %s!\n", v)
		os.Exit(0)
	},
}