package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/ruff_completer/cmd/action"
)

func init() {
	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"color":  action.ActionColors(),
		"config": carapace.ActionFiles(".toml"),
	})

	carapace.Gen(checkCmd).FlagCompletion(carapace.ActionMap{
		"cache-dir":        carapace.ActionDirectories(),
		"extend-fixable":   action.ActionRules(),
		"extend-ignore":    action.ActionRules(),
		"extend-select":    action.ActionRules(),
		"extend-unfixable": action.ActionRules(),
		"fixable":          action.ActionRules(),
		"ignore":           action.ActionRules(),
		"output-file":      carapace.ActionFiles(),
		"output-format":    action.ActionOutputFormats(),
		"select":           action.ActionRules(),
		"stdin-filename":   carapace.ActionFiles(),
		"target-version":   action.ActionPythonVersions(),
		"unfixable":        action.ActionRules(),
	})

	carapace.Gen(checkCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)

	carapace.Gen(ruleCmd).FlagCompletion(carapace.ActionMap{
		"output-format": carapace.ActionValues("json", "text"),
	})

	carapace.Gen(ruleCmd).PositionalCompletion(
		action.ActionRuleCodes(),
	)

	carapace.Gen(configCmd).FlagCompletion(carapace.ActionMap{
		"output-format": carapace.ActionValues("json", "text"),
	})

	carapace.Gen(configCmd).PositionalCompletion(
		action.ActionConfigOptions(),
	)

	carapace.Gen(linterCmd).FlagCompletion(carapace.ActionMap{
		"output-format": carapace.ActionValues("json", "text"),
	})

	carapace.Gen(generateShellCompletionCmd).PositionalCompletion(
		action.ActionShells(),
	)

	carapace.Gen(formatCmd).FlagCompletion(carapace.ActionMap{
		"cache-dir":      carapace.ActionDirectories(),
		"output-format":  action.ActionOutputFormats(),
		"stdin-filename": carapace.ActionFiles(),
		"target-version": action.ActionPythonVersions(),
	})

	carapace.Gen(formatCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)

	carapace.Gen(analyze_graphCmd).FlagCompletion(carapace.ActionMap{
		"direction":      action.ActionDirections(),
		"python":         carapace.ActionFiles(),
		"target-version": action.ActionPythonVersions(),
	})

	carapace.Gen(analyze_graphCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)

	carapace.Gen(versionCmd).FlagCompletion(carapace.ActionMap{
		"output-format": carapace.ActionValues("json", "text"),
	})
}
