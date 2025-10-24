package commands

import (
	"fmt"

	"github.com/hursty1/go_dbt_cli/internal/config"
	"github.com/hursty1/go_dbt_cli/internal/dbtcloud"
)

func Run(config config.CommandConfig) error {
	dbtCloudConfig, err := dbtcloud.ReadDbtCloudConfig(config.DbtFilePath)
	if err != nil {
		return err
	}
	// pf("RUN COMMAND")
	pf("CONFIG: %v\n",config)
	if config.Active {
		DisplayActive(dbtCloudConfig)
		// return nil
	}
	if config.List {
		DisplayAllProjects(dbtCloudConfig)
		// return nil
	}

	//process change
	if config.Action != "" {
		switch config.Action {
		case "action":
			fmt.Println(config.ProjectName)
		default:
			fmt.Println("Invalid Action")
		}
	}

	return nil
}

// dbt cloud set active-project --project-name "Name of Game 1"
var pf = fmt.Printf

func ActivateProject(config dbtcloud.DbtCloudConfig, project string) error {
	//verify project exists
	//if exists execute cli command dbt cloud set active-project --project-name "Name of Game 1"
	//

	return nil
}
func DisplayActive(config dbtcloud.DbtCloudConfig) {
	activeProject := ""
	for _, project := range config.Projects {
		if project.ProjectId == config.Context.ActiveProject {
			activeProject = project.ProjectName
		}
	}
	pf("\nActive Project: %v", activeProject)
	pf("\nActive Host:    %v", config.Context.ActiveHost)
}

func DisplayAllProjects(config dbtcloud.DbtCloudConfig) {
	
	for _, project := range config.Projects {
		pf("\n")
		DisplayProjectSummary(project)
	}
}

func DisplayProjectSummary(project dbtcloud.DbtProject) {
	pf("\nProject Name: %s", project.ProjectName)
	pf("\nProject Id: %s", project.ProjectId)
	pf("\nAccount Name: %s", project.AccountName)
}