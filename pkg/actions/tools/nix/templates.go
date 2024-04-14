package nix

import (
	"encoding/json"
	"time"

	"github.com/carapace-sh/carapace"
)

// ActionTemplates completes the templates outputs of the system template flake
func ActionTemplates() carapace.Action {
	return carapace.ActionExecCommand("nix", "flake", "show", "--json", "templates")(func(output []byte) carapace.Action {
		templatesSchema := struct {
			Templates map[string]struct {
				Description string `json:"description"`
			} `json:"templates"`
		}{}
		if err := json.Unmarshal(output, &templatesSchema); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for template, value := range templatesSchema.Templates {
			vals = append(vals, template, value.Description)
		}
		return carapace.ActionValuesDescribed(vals...).Prefix("templates#")
	}).Cache(time.Hour)
}
