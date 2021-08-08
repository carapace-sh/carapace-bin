package action

import (
	"io/ioutil"
	"os"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/util"
	"gopkg.in/yaml.v3"
)

type pubspec struct {
	Dependencies    map[string]string `yaml:"dependencies"`
	DevDependencies map[string]string `yaml:"dev_dependencies"`
}

func loadPubspec() (*pubspec, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	path, err := util.FindReverse(wd, "pubspec.yaml")
	if err != nil {
		return nil, err
	}

	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var p *pubspec
	err = yaml.Unmarshal(content, &p)
	return p, err
}

func ActionDependencies() carapace.Action { // TODO configure which dependencies to include
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		p, err := loadPubspec()
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for name, version := range p.Dependencies {
			vals = append(vals, name, version)
		}

		for name, version := range p.DevDependencies {
			vals = append(vals, name, version)
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
