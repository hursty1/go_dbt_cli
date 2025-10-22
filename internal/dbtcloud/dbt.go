package dbtcloud

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"gopkg.in/yaml.v3"
)

type StringInt int

func (s *StringInt) UnmarshalYAML(value *yaml.Node) error {
	var asInt int
	var asStr string

	switch value.Tag {
	case "!!str":
		asStr = value.Value
		i, err := strconv.Atoi(asStr)
		if err != nil {
			return fmt.Errorf("invalid integer string %q: %v", asStr, err)
		}
		*s = StringInt(i)
	case "!!int":
		if err := value.Decode(&asInt); err != nil {
			return err
		}
		*s = StringInt(asInt)
	default:
		return fmt.Errorf("unexpected YAML type: %s", value.Tag)
	}

	return nil
}

type DbtProject struct {
	ProjectName string `yaml:"project-name"`
	ProjectId   StringInt    `yaml:"project-id"`
	AccountName string `yaml:"account-name"`
	AccountId   StringInt    `yaml:"account-id"`
	AccountHost string `yaml:"account-host"`
	TokenName   string `yaml:"token-name"`
	TokenValue  string `yaml:"token-value"`
}

type DbtContext struct {
	ActiveHost    string `yaml:"active-host"`
	ActiveProject StringInt `yaml:"active-project"`
}
type DbtCloudConfig struct {
	Version  StringInt        `yaml:"version"`
	Context  DbtContext `yaml:"context"`
	Projects []DbtProject `yaml:"projects"`
}

func ReadDbtCloudConfig(filePath string) (DbtCloudConfig, error) {

	dbtConfig := DbtCloudConfig{}

	file, err := os.Open(filePath)
    if err != nil {
        fmt.Printf("Error opening file: %v\n", err)
        return dbtConfig, err
    }
    defer file.Close()

	contents, err := io.ReadAll(file)
    if err != nil {
        fmt.Printf("Error reading file: %v\n", err)
        return dbtConfig, err
    }

	if err := yaml.Unmarshal(contents, &dbtConfig); err != nil {
		fmt.Printf("Error Parsing YAML file: %v \n", err)
		return dbtConfig, err
	}
	// fmt.Printf("Contents are: %v\n", dbtConfig)
	return dbtConfig, nil
}