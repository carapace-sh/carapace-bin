package env

import (
	"os"
	"os/exec"

	"github.com/rsteube/carapace"
	spec "github.com/rsteube/carapace-spec"
	"github.com/rsteube/carapace/pkg/xdg"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

type variables struct {
	Condition  func(c carapace.Context) bool
	Names      map[string]string
	Completion map[string]carapace.Action
}

func (v *variables) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var env struct {
		Names      map[string]string
		Completion map[string][]string
	}
	if err := unmarshal(&env); err != nil {
		return err
	}
	v.Names = env.Names
	v.Completion = make(map[string]carapace.Action)
	for name, completion := range env.Completion {
		v.Completion[name] = spec.NewAction(completion).Parse(&cobra.Command{})
	}
	return nil
}

func checkPath(s string) func(c carapace.Context) bool {
	return func(c carapace.Context) bool {
		_, err := exec.LookPath(s) // TODO copy function to carapace as this needs to use carapace.Context$Env
		return err == nil
	}
}

var knownVariables = map[string]variables{}

func ActionKnownEnvironmentVariables() carapace.Action {
	return carapace.Batch(
		actionKnownEnvironmentVariables(),
		actionCustomEnvironmentVariables(),
	).ToA()
}

func actionKnownEnvironmentVariables() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		vals := make([]string, 0)
		for _, v := range knownVariables {
			if v.Condition != nil && !v.Condition(c) {
				continue
			}

			for name, description := range v.Names {
				vals = append(vals, name, description)
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Tag("known environment variables")
}

func ActionEnvironmentVariableValues(s string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if custom, err := loadCustomVariables(); err == nil {
			if action, ok := custom.Completion[s]; ok {
				return action
			}
		}

		for _, v := range knownVariables {
			if action, ok := v.Completion[s]; ok {
				return action
			}
		}
		return carapace.ActionFiles()
	})
}

func actionCustomEnvironmentVariables() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		v, err := loadCustomVariables()
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for name, description := range v.Names {
			vals = append(vals, name, description)
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Tag("custom environment variables")
}

func loadCustomVariables() (*variables, error) {
	dir, err := xdg.UserConfigDir()
	if err != nil {
		return nil, err
	}

	content, err := os.ReadFile(dir + "/carapace/env.yaml")
	if err != nil {
		if !os.IsNotExist(err) {
			return nil, err
		}
		return &variables{
			Names:      map[string]string{},
			Completion: make(map[string]carapace.Action),
		}, nil
	}

	var v variables
	if err := yaml.Unmarshal(content, &v); err != nil {
		return nil, err
	}

	return &v, nil
}
