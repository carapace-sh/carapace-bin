package rails

import "github.com/carapace-sh/carapace"

// ActionEnvironments completes Rails environment names
//
//	development
//	test
func ActionEnvironments() carapace.Action {
	return carapace.ActionValues(
		"development",
		"test",
		"production",
		"staging",
	).Tag("environments").Uid("rails", "environment")
}
