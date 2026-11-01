package action

import (
	"encoding/json"

	"github.com/carapace-sh/carapace"
)

type packageDescription struct {
	Products []struct {
		Name string `json:"name"`
	} `json:"products"`
	Targets []struct {
		Name string `json:"name"`
	} `json:"targets"`
}

type executable struct {
	Package string `json:"package"`
	Name    string `json:"name"`
}

func ActionProducts() carapace.Action {
	return actionDescription(func(description packageDescription) []string {
		values := make([]string, 0, len(description.Products))
		for _, product := range description.Products {
			values = append(values, product.Name)
		}
		return values
	})
}

func ActionTargets() carapace.Action {
	return actionDescription(func(description packageDescription) []string {
		values := make([]string, 0, len(description.Targets))
		for _, target := range description.Targets {
			values = append(values, target.Name)
		}
		return values
	})
}

func ActionTargetsAndProducts() carapace.Action {
	return actionDescription(func(description packageDescription) []string {
		values := make([]string, 0, len(description.Products)+len(description.Targets))
		visited := make(map[string]bool)
		for _, target := range description.Targets {
			values = append(values, target.Name)
			visited[target.Name] = true
		}
		for _, product := range description.Products {
			if !visited[product.Name] {
				values = append(values, product.Name)
			}
		}
		return values
	})
}

func ActionExecutables() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("swift", "package", "show-executables", "--format", "json")(func(output []byte) carapace.Action {
			var executables []executable
			if err := json.Unmarshal(output, &executables); err != nil {
				return carapace.ActionMessage("failed to parse executables: " + err.Error())
			}

			values := make([]string, 0, len(executables))
			for _, executable := range executables {
				if executable.Package == "" { // skip executables of dependencies as they can't be run
					values = append(values, executable.Name)
				}
			}
			return carapace.ActionValues(values...)
		})
	})
}

func actionDescription(f func(packageDescription) []string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("swift", "package", "describe", "--type", "json")(func(output []byte) carapace.Action {
			var description packageDescription
			if err := json.Unmarshal(output, &description); err != nil {
				return carapace.ActionMessage("failed to parse package description: " + err.Error())
			}
			return carapace.ActionValues(f(description)...)
		})
	})
}

type supportedFeatures struct {
	Features struct {
		Optional     []feature `json:"optional"`
		Upcoming     []feature `json:"upcoming"`
		Experimental []feature `json:"experimental"`
	} `json:"features"`
}

type feature struct {
	Name       string `json:"name"`
	Migratable bool   `json:"migratable"`
}

func ActionUpcomingFeatures() carapace.Action {
	return actionFeatures(func(features supportedFeatures) []feature {
		return features.Features.Upcoming
	})
}

func ActionExperimentalFeatures() carapace.Action {
	return actionFeatures(func(features supportedFeatures) []feature {
		return features.Features.Experimental
	})
}

func ActionMigratableFeatures() carapace.Action {
	return actionFeatures(func(features supportedFeatures) []feature {
		migratable := make([]feature, 0)
		for _, group := range [][]feature{
			features.Features.Optional,
			features.Features.Upcoming,
			features.Features.Experimental,
		} {
			for _, feature := range group {
				if feature.Migratable {
					migratable = append(migratable, feature)
				}
			}
		}
		return migratable
	})
}

func actionFeatures(f func(supportedFeatures) []feature) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("swift", "-print-supported-features")(func(output []byte) carapace.Action {
			var features supportedFeatures
			if err := json.Unmarshal(output, &features); err != nil {
				return carapace.ActionMessage("failed to parse supported features: " + err.Error())
			}

			selected := f(features)
			values := make([]string, 0, len(selected))
			for _, feature := range selected {
				values = append(values, feature.Name)
			}
			return carapace.ActionValues(values...)
		})
	})
}
