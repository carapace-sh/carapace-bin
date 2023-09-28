package env

import (
	"os"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/internal/condition"
	"github.com/rsteube/carapace-bin/pkg/conditions"
	spec "github.com/rsteube/carapace-spec"
	"github.com/rsteube/carapace/pkg/xdg"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

type variables struct {
	Condition          condition.Condition
	Variables          map[string]string
	VariableCompletion map[string]carapace.Action
	// PlaceholderCompletion map[string]carapace.Action // TODO
}

func (v *variables) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var env struct {
		Condition  []string
		Variables  map[string]string
		Completion struct {
			Variable map[string][]string
			// Placeholder map[string][]string // TODO
		}
	}
	if err := unmarshal(&env); err != nil {
		return err
	}

	conds := make([]condition.Condition, 0)
	for _, c := range env.Condition {
		m, err := conditions.MacroMap.Lookup(c)
		if err != nil {
			return err
		}
		conds = append(conds, m.Parse(c))
	}

	v.Condition = condition.Of(conds...)
	v.Variables = env.Variables
	v.VariableCompletion = make(map[string]carapace.Action)
	for name, completion := range env.Completion.Variable {
		v.VariableCompletion[name] = spec.NewAction(completion).Parse(&cobra.Command{})
	}
	return nil
}

var knownVariables = map[string]variables{}

// ActionKnownEnvironmentVariables completes known environment variables
//
//	GOARCH (The architecture, or processor, for which to compile code)
//	GOOS (The operating system for which to compile code)
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

			for name, description := range v.Variables {
				vals = append(vals, name, description)
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Tag("known environment variables")
}

// ActionEnvironmentVariableValues completes values for given environment variable
func ActionEnvironmentVariableValues(s string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		dir, err := xdg.UserConfigDir()
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		entries, err := os.ReadDir(dir + "/carapace/variables")
		if err != nil {
			return carapace.ActionMessage(err.Error()) // TODO ignore if not exists
		}

		found := false
		batch := carapace.Batch()
		for _, entry := range entries {
			if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".yaml") {
				file := dir + "/carapace/variables/" + entry.Name()
				batch = append(batch, carapace.ActionCallback(func(c carapace.Context) carapace.Action {
					custom, err := loadCustomVariables(file)
					if err != nil {
						return carapace.ActionMessage(err.Error())
					}

					if custom.Condition == nil || custom.Condition(c) {
						if action, ok := custom.VariableCompletion[s]; ok {
							found = true
							return action.Usage(custom.Variables[s])
						}
					}

					for _, v := range knownVariables {
						if action, ok := v.VariableCompletion[s]; ok {
							found = true
							return action.Usage(v.Variables[s])
						}
					}
					return carapace.ActionValues()
				}))
			}
		}

		if a := batch.Invoke(c).Merge().ToA(); found {
			return a
		}
		return carapace.ActionFiles() // fallback
	})
}

func actionCustomEnvironmentVariables() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		dir, err := xdg.UserConfigDir()
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		entries, err := os.ReadDir(dir + "/carapace/variables")
		if err != nil {
			return carapace.ActionMessage(err.Error()) // TODO ignore if not exists
		}

		batch := carapace.Batch()
		for _, entry := range entries {
			if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".yaml") {
				file := dir + "/carapace/variables/" + entry.Name()
				batch = append(batch, carapace.ActionCallback(func(c carapace.Context) carapace.Action {
					custom, err := loadCustomVariables(file)
					if err != nil {
						return carapace.ActionMessage(err.Error())
					}

					if custom.Condition != nil && !custom.Condition(c) {
						return carapace.ActionValues() // skip when condition failed
					}

					vals := make([]string, 0)
					for name, description := range custom.Variables {
						vals = append(vals, name, description)
					}
					return carapace.ActionValuesDescribed(vals...)

				}))
			}
		}
		return batch.ToA()
	}).Tag("custom environment variables")
}

func loadCustomVariables(file string) (*variables, error) {
	content, err := os.ReadFile(file)
	if err != nil {
		if !os.IsNotExist(err) {
			return nil, err
		}
		return &variables{
			Variables:          map[string]string{},
			VariableCompletion: make(map[string]carapace.Action),
		}, nil
	}

	var v variables
	if err := yaml.Unmarshal(content, &v); err != nil {
		return nil, err
	}

	return &v, nil
}
