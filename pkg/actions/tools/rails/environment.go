package rails

import "github.com/carapace-sh/carapace"

// ActionEnvironments completes Rails environment names
//
//	development
//	test
func ActionEnvironments() carapace.Action {
	return carapace.ActionValues(
		"development",
		"production",
		"staging",
		"test",
	).Tag("environments").Uid("rails", "environment")
}
