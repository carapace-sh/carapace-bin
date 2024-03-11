package env

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/conditions"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/carapace-sh/carapace/pkg/style"
)

func init() {
	knownVariables["zoxide"] = func() variables {
		_bool := carapace.ActionValuesDescribed("0", "disabled", "1", "enabled").StyleF(style.ForKeyword)
		return variables{
			Condition: conditions.ConditionPath("zoxide"),
			Variables: map[string]string{
				"_ZO_DATA_DIR":         "Path for zoxide data files",
				"_ZO_ECHO":             "Print the matched directory before navigating to it when set to 1",
				"_ZO_EXCLUDE_DIRS":     "List of directory globs to be excluded",
				"_ZO_FZF_OPTS":         "Custom flags to pass to fzf",
				"_ZO_MAXAGE":           "Maximum total age after which entries start getting deleted",
				"_ZO_RESOLVE_SYMLINKS": "Resolve symlinks when storing paths",
			},
			VariableCompletion: map[string]carapace.Action{
				"_ZO_DATA_DIR":         carapace.ActionDirectories(),
				"_ZO_ECHO":             _bool,
				"_ZO_EXCLUDE_DIRS":     carapace.ActionDirectories().List(","), // TODO verify
				"_ZO_FZF_OPTS":         bridge.ActionCarapaceBin("fzf").Split(),
				"_ZO_MAXAGE":           carapace.ActionValues(),
				"_ZO_RESOLVE_SYMLINKS": _bool,
			},
		}
	}
}
