package pub

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/util"
	"gopkg.in/yaml.v3"
)

type version string

func (v *version) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	_ = unmarshal(&s) // ignore struct
	*v = version(s)
	return nil
}

type pubspec struct {
	Dependencies    map[string]version     `yaml:"dependencies"`
	DevDependencies map[string]version     `yaml:"dev_dependencies"`
	Executables     map[string]interface{} `yaml:"executables"`
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

type pubspecLock struct {
	Packages map[string]struct {
		Version string
	}
}

func loadPubspecLock() (*pubspecLock, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	path, err := util.FindReverse(wd, "pubspec.lock")
	if err != nil {
		return nil, err
	}

	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var p *pubspecLock
	err = yaml.Unmarshal(content, &p)
	return p, err
}

// ActionDependencies completes pubspec dependencies
//   build_runner (^1.5.0)
//   build_web_compilers (^2.1.0)
func ActionDependencies() carapace.Action { // TODO configure which dependencies to include
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		p, err := loadPubspec()
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for name, version := range p.Dependencies {
			vals = append(vals, name, string(version))
		}

		for name, version := range p.DevDependencies {
			vals = append(vals, name, string(version))
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}

// ActionHostedExecutables completes executables from pub_cache
//   dcat
//   dgrep
func ActionHostedExecutables(name string, version string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if version == "" {
			pl, err := loadPubspecLock()
			if err != nil {
				return carapace.ActionMessage(err.Error())
			}
			if pkg, ok := pl.Packages[name]; !ok {
				return carapace.ActionMessage(fmt.Sprintf("Missing '%v' in pubspec.lock", name))
			} else {
				version = pkg.Version
			}
		}

		home, err := os.UserHomeDir()
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		files, err := ioutil.ReadDir(fmt.Sprintf("%v/.pub-cache/hosted/pub.dartlang.org/%v-%v/bin", home, name, version))
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for _, file := range files {
			if !file.IsDir() && strings.HasSuffix(file.Name(), ".dart") {
				vals = append(vals, strings.TrimSuffix(filepath.Base(file.Name()), ".dart"))
			}
		}
		return carapace.ActionValues(vals...)
	})
}
