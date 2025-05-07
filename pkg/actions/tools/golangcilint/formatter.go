package golangcilint

import (
	"encoding/json"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

// ActionFormatters completes formatters
//
//	gci (Checks if code and import statements are formatted)
//	gofmt (Checks if the code is formatted)
func ActionFormatters(config string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := []string{"formatters", "--json"}
		if config != "" {
			args = append(args, "--config", config)
		}
		return carapace.ActionExecCommand("golangci-lint", args...)(func(output []byte) carapace.Action {
			var formatters struct {
				Enabled []struct {
					Name        string
					Description string
					Deprecated  bool
				}
				Disabled []struct {
					Name        string
					Description string
					Deprecated  bool
				}
			}
			if err := json.Unmarshal(output, &formatters); err != nil {
				return carapace.ActionMessage(err.Error())
			}

			vals := make([]string, 0)
			for _, enabled := range formatters.Enabled {
				vals = append(vals, enabled.Name, enabled.Description, style.Carapace.KeywordPositive)
			}
			for _, disabled := range formatters.Disabled {
				vals = append(vals, disabled.Name, disabled.Description, style.Carapace.KeywordNegative)
			}
			return carapace.ActionStyledValuesDescribed(vals...)
		})
	}).Tag("formatters")
}
