package env

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/cmd/carapace/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/env"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/carapace-sh/carapace/pkg/style"
)

func init() {
	knownVariables["carapace"] = func() variables {
		_bool := carapace.ActionValuesDescribed("0", "disabled", "1", "enabled").StyleF(style.ForKeyword)
		return variables{
			Variables: map[string]string{
				// carapace
				"CARAPACE_COVERDIR":      "coverage directory for sandbox tests",
				"CARAPACE_COMPLINE":      "current command line",
				"CARAPACE_ENV":           "register get-env, set-env and unset-env",
				"CARAPACE_HIDDEN":        "show hidden commands/flags",
				"CARAPACE_LENIENT":       "allow unknown flags",
				"CARAPACE_LOG":           "enable logging",
				"CARAPACE_MATCH":         "match case insensitive",
				"CARAPACE_MERGEFLAGS":    "merge flags to single tag group",
				"CARAPACE_NOSPACE":       "nospace suffixes",
				"CARAPACE_SANDBOX":       "mock context for sandbox tests",
				"CARAPACE_TOOLTIP":       "enable tooltip style",
				"CARAPACE_UNFILTERED":    "skip the final filtering step",
				"CARAPACE_ZSH_HASH_DIRS": "zsh hash directories",
				// carapace-bin
				"CARAPACE_EXCLUDES": "internal completers to exclude",
				"CARAPACE_BRIDGES":  "implicit bridges",
			},
			VariableCompletion: map[string]carapace.Action{
				// carapace
				"CARAPACE_COVERDIR": carapace.ActionDirectories(),
				"CARAPACE_COMPLINE": bridge.ActionCarapaceBin().Split(),
				"CARAPACE_ENV":      _bool,
				"CARAPACE_HIDDEN": carapace.ActionStyledValuesDescribed(
					"0", "disabled", style.Carapace.KeywordNegative,
					"1", "enabled", style.Carapace.KeywordPositive,
					"2", "enabled including _carapace", style.Carapace.KeywordPositive,
				),
				"CARAPACE_LENIENT": _bool,
				"CARAPACE_LOG":     _bool,
				"CARAPACE_MATCH": carapace.ActionValuesDescribed(
					"0", "CASE_SENSITIVE",
					"1", "CASE_INSENSITIVE",
				).StyleF(style.ForKeyword),
				"CARAPACE_MERGEFLAGS": _bool,
				"CARAPACE_NOSPACE": carapace.ActionValuesDescribed(
					"*", "match all",
				).UniqueList(""),
				"CARAPACE_SANDBOX":    carapace.ActionValues(),
				"CARAPACE_TOOLTIP":    _bool,
				"CARAPACE_UNFILTERED": _bool,
				// carapace-bin
				"CARAPACE_EXCLUDES": carapace.Batch(
					carapace.ActionCallback(func(c carapace.Context) carapace.Action {
						c.Setenv(env.CARAPACE_EXCLUDES, "")
						return action.ActionCompleters(action.CompleterOpts{Internal: true}).Invoke(c).ToA()
					}),
					carapace.ActionValuesDescribed("*", "exclude all"),
				).ToA().UniqueList(","),
				"CARAPACE_BRIDGES": carapace.ActionStyledValues(
					"bash", "#d35673",
					"fish", "#7ea8fc",
					"inshellisense", style.Default,
					"zsh", "#efda53",
				).UniqueList(","),
			},
		}
	}

}
