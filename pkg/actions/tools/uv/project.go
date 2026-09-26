package uv

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/traverse"
	toml "github.com/pelletier/go-toml"
)

func pyprojectPath(c carapace.Context) (string, error) {
	dir, err := traverse.Parent("pyproject.toml")(c)
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "pyproject.toml"), nil
}

type pyproject struct {
	Project struct {
		Scripts              map[string]string   `toml:"scripts"`
		GuiScripts           map[string]string   `toml:"gui-scripts"`
		Dependencies         []string            `toml:"dependencies"`
		OptionalDependencies map[string][]string `toml:"optional-dependencies"`
	} `toml:"project"`
	DependencyGroups map[string]interface{} `toml:"dependency-groups"`
}

func loadPyproject(c carapace.Context) (*pyproject, error) {
	path, err := pyprojectPath(c)
	if err != nil {
		return nil, err
	}
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var config pyproject
	if err := toml.Unmarshal(content, &config); err != nil {
		return nil, err
	}
	return &config, nil
}

// ActionDependencyGroups completes dependency groups of the project
//
//	dev
//	lint.mypy
func ActionDependencyGroups() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		config, err := loadPyproject(c)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for key, value := range config.DependencyGroups {
			switch value.(type) {
			case []interface{}:
				vals = append(vals, key)
			case map[string]interface{}: // nested groups are referenced with a dotted name
				for subkey, subvalue := range value.(map[string]interface{}) {
					if _, ok := subvalue.([]interface{}); ok {
						vals = append(vals, fmt.Sprintf("%v.%v", key, subkey))
					}
				}
			}
		}
		return carapace.ActionValues(vals...).
			UidF(Uid("dependency-group")).
			QueryF(Uid("dependency-group"))
	}).Tag("dependency groups")
}

// ActionExtras completes extras of the project
//
//	all
//	dev
func ActionExtras() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		config, err := loadPyproject(c)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0, len(config.Project.OptionalDependencies))
		for name := range config.Project.OptionalDependencies {
			vals = append(vals, name)
		}
		return carapace.ActionValues(vals...).
			UidF(Uid("extra")).
			QueryF(Uid("extra"))
	}).Tag("extras")
}

// ActionScripts completes scripts of the project
//
//	app (app:main)
//	tool (tool.cli:main)
func ActionScripts() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		config, err := loadPyproject(c)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for _, scripts := range []map[string]string{config.Project.Scripts, config.Project.GuiScripts} {
			for name, target := range scripts {
				vals = append(vals, name, target)
			}
		}
		return carapace.ActionValuesDescribed(vals...).
			UidF(Uid("script")).
			QueryF(Uid("script"))
	}).Tag("scripts")
}

// ActionProjectDependencies completes dependencies of the project
//
//	click
//	ruff
func ActionProjectDependencies() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		config, err := loadPyproject(c)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		seen := map[string]bool{}
		vals := make([]string, 0)
		add := func(requirements ...string) {
			for _, requirement := range requirements {
				if name := dependencyName(requirement); !seen[name] {
					seen[name] = true
					vals = append(vals, name)
				}
			}
		}

		add(config.Project.Dependencies...)
		for _, requirements := range config.Project.OptionalDependencies {
			add(requirements...)
		}
		for _, value := range config.DependencyGroups {
			switch value := value.(type) {
			case []interface{}:
				for _, requirement := range value {
					if requirement, ok := requirement.(string); ok {
						add(requirement)
					}
				}
			case map[string]interface{}:
				for _, subvalue := range value {
					if requirements, ok := subvalue.([]interface{}); ok {
						for _, requirement := range requirements {
							if requirement, ok := requirement.(string); ok {
								add(requirement)
							}
						}
					}
				}
			}
		}
		return carapace.ActionValues(vals...).
			UidF(Uid("project-dependency")).
			QueryF(Uid("project-dependency"))
	}).Tag("project dependencies")
}

var dependencyNameRE = regexp.MustCompile(`^\s*([A-Za-z0-9][A-Za-z0-9._-]*)`)

func dependencyName(requirement string) string {
	if match := dependencyNameRE.FindStringSubmatch(requirement); match != nil {
		return strings.ToLower(match[1])
	}
	return requirement
}
