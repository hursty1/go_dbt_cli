package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/hursty1/go_dbt_cli/internal/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the current version of dbthelper",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("dbthelper version:", version.Get())
		latest, err := LatestVersion()
		if err != nil {
			fmt.Println("update check failed:", err)
			// return false
		}
		fmt.Println("Latest Version is:", latest)
	},
}


func LatestVersion() (string, error) {
    resp, err := http.Get("https://api.github.com/repos/hursty1/go_dbt_cli/tags")
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    var tags []struct {
        Name string `json:"name"`
    }
    if err := json.NewDecoder(resp.Body).Decode(&tags); err != nil {
        return "", err
    }

    if len(tags) == 0 {
        return "", fmt.Errorf("no tags found")
    }

    // first tag is latest (GitHub sorts by date)
    return tags[0].Name, nil
}

func IsOutdated(current string) bool {
    latest, err := LatestVersion()
    if err != nil {
        fmt.Println("update check failed:", err)
        return false
    }

    current = strings.TrimSpace(current)
    latest = strings.TrimSpace(latest)

    if current == "" {
        fmt.Println("unknown current version, skipping update check")
        return false
    }

    if current == "dev" {
        fmt.Println("development build detected — upgrading to latest release...")
        return true
    }

    if latest != current {
        fmt.Printf("A new version is available: %s → %s\n", current, latest)
        return true
    }

    return false
}