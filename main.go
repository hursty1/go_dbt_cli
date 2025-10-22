package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/hursty1/dbt_helper_go/internal/commands"
	"github.com/hursty1/dbt_helper_go/internal/config"
)

func main() {
	// fmt.Println("Main")

	listProjects := flag.Bool("list", false, "List all Configured projects")
	activeProject := flag.Bool("active", false, "Show active Project")


	flag.BoolVar(listProjects, "l", false, "List all configured dbt projects")
	flag.BoolVar(activeProject, "a", false, "Show active dbt Project")

	args := os.Args[1:]
	flags := []string{}
	positional := []string{}

	for _, arg := range args {
		if strings.HasPrefix(arg, "-") {
			flags = append(flags, arg)
		} else {
			positional = append(positional, arg)
		}
	}

	flag.CommandLine.Parse(flags)

	//check if dbt is in correct spot
	homeDir, err := os.UserHomeDir()
	if err != nil {panic(err)}

	dbtPath := filepath.Join(homeDir, ".dbt")
	dbtFile := filepath.Join(dbtPath, "dbt_cloud.yml")

	info, err := os.Stat(dbtPath)
	if os.IsNotExist(err){
		fmt.Printf("Error: DBT folder not found at %s\n", dbtPath)
		os.Exit(1)
	} else if err != nil {
		fmt.Printf("Error accessing DBT folder: %v\n", err)
		os.Exit(1)
	}
	if !info.IsDir() {
        fmt.Printf("Error: %s exists but is not a directory\n", dbtPath)
        os.Exit(1)
    }

	//dbt_cloud.yml
	//check / open file
	if _, err := os.Stat(dbtFile); os.IsNotExist(err) {
        fmt.Printf("Error: dbt_cloud.yml not found in %s\n", dbtPath)
        os.Exit(1)
    }

	config := config.CommandConfig{
		List: *listProjects,
		Active: *activeProject,
		DbtFilePath: dbtFile,
	}

	err = commands.Run(config)
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}