package nix

import (
	"encoding/json"
	"fmt"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

type devShellsSchema struct {
	DevShells map[string]map[string]struct {
		Name string `json:"name"`
	} `json:"devShells"`
}

// ActionDevShells completes development shells from given flake
func ActionDevShells(flake string) carapace.Action {
	return carapace.ActionExecCommand("nix", "flake", "show", "--json", flake)(func(output []byte) carapace.Action {
		devShells := &devShellsSchema{}
		if err := json.Unmarshal(output, devShells); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		return carapace.ActionMultiPartsN(".", 2, func(c carapace.Context) carapace.Action {
			if len(c.Parts) == 0 {
				currentSystem, err := getCurrentSystem(c)
				if err != nil {
					return carapace.ActionMessage(err.Error())
				}

				systems := make([]string, 0)
				for system := range devShells.DevShells {
					systems = append(systems, system)
				}

				return carapace.ActionValues(systems...).StyleF(func(s string, sc style.Context) string {
					if s == currentSystem {
						return style.Blue
					}
					return style.Default
				}).Suffix(".")
			}

			vals := make([]string, 0)
			currentSystem, ok := devShells.DevShells[c.Parts[0]]
			if !ok {
				return carapace.ActionValues()
			}

			for name, value := range currentSystem {
				vals = append(vals, name, value.Name)
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}

func getCurrentSystem(c carapace.Context) (string, error) {
	output, err := c.Command("nix", "show-config", "--json").Output()
	if err != nil {
		return "", fmt.Errorf("nix show-config: %w", err)
	}

	config := struct {
		System struct {
			Value string `json:"value"`
		} `json:"system"`
	}{}
	if err := json.Unmarshal(output, &config); err != nil {
		return "", fmt.Errorf("nix show-config: %w", err)
	}

	return config.System.Value, nil
}
