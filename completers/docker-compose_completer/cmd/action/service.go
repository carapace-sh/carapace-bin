package action

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func ActionServices(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if services, _, err := loadFile(cmd); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			return carapace.ActionValues(services...)
		}
	})
}

func ActionVolumes(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if _, volumes, err := loadFile(cmd); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			return carapace.ActionValues(volumes...)
		}
	})
}

type compose struct {
	Version  string
	Services map[string]interface{}
	Volumes  map[string]interface{}
}

func fileLocation(cmd *cobra.Command) (string, error) {
	if flag := cmd.Root().Flag("file"); flag.Changed {
		return flag.Value.String(), nil
	} else {
		if path, err := os.Getwd(); err == nil {
			return traverse(path)
		} else {
			return "", err
		}
	}
}

func traverse(path string) (string, error) {
	if _, err := os.Stat(path + "/docker-compose.yml"); err == nil {
		return path + "/docker-compose.yml", nil
	} else {
		if path == "/" {
			return "", errors.New("no docker-compose.yml present")
		}
		return traverse(filepath.Dir(path))
	}
}

func loadFile(cmd *cobra.Command) ([]string, []string, error) {
	file, err := fileLocation(cmd)
	if err != nil {
		return nil, nil, err
	}

	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, nil, err
	}

	var content compose
	if err = yaml.Unmarshal(yamlFile, &content); err != nil {
		return nil, nil, err
	}
	return keys(content.Services), keys(content.Volumes), nil
}

func keys(m map[string]interface{}) []string {
	keys := make([]string, len(m))
	index := 0
	for key := range m {
		keys[index] = key
		index += 1
	}
	return keys
}
