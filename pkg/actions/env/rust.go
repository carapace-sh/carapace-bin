package env

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/conditions"
	"github.com/rsteube/carapace/pkg/style"
)

func init() {
	knownVariables["rust"] = func() variables {
		_bool := carapace.ActionValuesDescribed("0", "disabled", "1", "enabled").StyleF(style.ForKeyword)
		return variables{
			Condition: conditions.ConditionPath("rustc"),
			Variables: map[string]string{
				"RUST_TEST_THREADS":   "The test framework Rust provides executes tests in parallel",
				"RUST_TEST_NOCAPTURE": "Synonym for the --nocapture flag",
				"RUST_MIN_STACK":      "Sets the minimum stack size for new threads",
				"RUST_BACKTRACE":      "Produces a backtrace in the output of a program which panics",
			},
			VariableCompletion: map[string]carapace.Action{
				"RUST_TEST_THREADS":   carapace.ActionValues(),
				"RUST_TEST_NOCAPTURE": _bool,
				"RUST_MIN_STACK":      carapace.ActionValues(),
				"RUST_BACKTRACE":      _bool,
			},
		}
	}
}
