package env

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/conditions"
)

func init() {
	knownVariables["dagger"] = func() variables {
		return variables{
			Condition: conditions.ConditionPath("dagger"),
			Variables: map[string]string{
				"DAGGER_CHECK":                "check if dagger is up to date",
				"DAGGER_NAME":                 "project name",
				"DAGGER_TEMPLATE":             "Template name",
				"DAGGER_PRIVATE_KEY_FILE":     "Private ssh key",
				"DAGGER_PRIVATE_KEY_PASSWORD": "Private ssh key password",
				"DAGGER_UPDATE":               "Update to latest version of specified packages",
				"DAGGER_PLAN":                 "Path to plan",
				"DAGGER_CACHE_FROM":           "External cache sources",
				"DAGGER_CACHE_TO":             "Cache destinations",
				"DAGGER_DRY_RUN":              "Dry run mode",
				"DAGGER_NO_CACHE":             "Disable caching",
				"DAGGER_OUTPUT":               "File path to write the action's output values",
				"DAGGER_OUTPUT_FORMAT":        "Format for output values",
				"DAGGER_PLATFORM":             "Set target build platform",
				"DAGGER_WITH":                 "Set value of dagger value at runtime",
			},
			VariableCompletion: map[string]carapace.Action{
				"DAGGER_CHECK":                carapace.ActionValues("1"),
				"DAGGER_NAME":                 carapace.ActionValues(),
				"DAGGER_TEMPLATE":             carapace.ActionValues(),
				"DAGGER_PRIVATE_KEY_FILE":     carapace.ActionFiles(),
				"DAGGER_PRIVATE_KEY_PASSWORD": carapace.ActionValues(),
				"DAGGER_UPDATE":               carapace.ActionValues("1"),
				"DAGGER_PLAN":                 carapace.ActionDirectories(),
				"DAGGER_CACHE_FROM":           carapace.ActionValues(),
				"DAGGER_CACHE_TO":             carapace.ActionValues(),
				"DAGGER_DRY_RUN":              carapace.ActionValues("1"),
				"DAGGER_NO_CACHE":             carapace.ActionValues("1"),
				"DAGGER_OUTPUT":               carapace.ActionFiles(),
				"DAGGER_OUTPUT_FORMAT":        carapace.ActionValues("plain", "json", "yaml"),
				"DAGGER_PLATFORM":             carapace.ActionValues(),
				"DAGGER_WITH":                 carapace.ActionValues(),
			},
		}
	}
}
