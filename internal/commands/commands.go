package commands

import (
	"fmt"

	"github.com/hursty1/dbt_helper_go/internal/config"
	"github.com/hursty1/dbt_helper_go/internal/dbtcloud"
)

func Run(config config.CommandConfig) error {
	dbtCloudConfig, err := dbtcloud.ReadDbtCloudConfig(config.DbtFilePath)
	if err != nil {
		return err
	}
	// pf("RUN COMMAND")
	// pf("CONFIG: %v",config)
	if config.Active {
		DisplayActive(dbtCloudConfig)
		// return nil
	}
	if config.List {
		DisplayAllProjects(dbtCloudConfig)
		// return nil
	}

	//process change

	return nil
}
var pf = fmt.Printf
func DisplayActive(config dbtcloud.DbtCloudConfig) {
	activeProject := ""
	for _, project := range config.Projects {
		if project.AccountId == config.Context.ActiveProject {
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