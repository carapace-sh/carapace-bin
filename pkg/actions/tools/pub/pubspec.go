package pub

import (
	"fmt"
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

func loadPubspec(c carapace.Context) (*pubspec, error) {
	path, err := util.FindReverse(c.Dir, "pubspec.yaml")
	if err != nil {
		return nil, err
	}

	content, err := os.ReadFile(path)
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

func loadPubspecLock(c carapace.Context) (*pubspecLock, error) {
	path, err := util.FindReverse(c.Dir, "pubspec.lock")
	if err != nil {
		return nil, err
	}

	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var p *pubspecLock
	err = yaml.Unmarshal(content, &p)
	return p, err
}

// ActionDependencies completes pubspec dependencies
//
//	build_runner (^1.5.0)
//	build_web_compilers (^2.1.0)
func ActionDependencies() carapace.Action { // TODO configure which dependencies to include
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		p, err := loadPubspec(c)
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

type HostedExecutablesOpts struct {
	Name    string
	Version string
}

// ActionHostedExecutables completes executables from pub_cache
//
//	dcat
//	dgrep
func ActionHostedExecutables(opts HostedExecutablesOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if opts.Version == "" {
			pl, err := loadPubspecLock(c)
			if err != nil {
				return carapace.ActionMessage(err.Error())
			}
			if pkg, ok := pl.Packages[opts.Name]; !ok {
				return carapace.ActionMessage(fmt.Sprintf("Missing '%v' in pubspec.lock", opts.Name))
			} else {
				opts.Version = pkg.Version
			}
		}

		home, err := os.UserHomeDir()
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		files, err := os.ReadDir(fmt.Sprintf("%v/.pub-cache/hosted/pub.dartlang.org/%v-%v/bin", home, opts.Name, opts.Version))
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
