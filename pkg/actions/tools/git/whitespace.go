package git

import "github.com/rsteube/carapace"

// ActionWhitespaceModes completes whitespace modes
//
//	error (outputs warnings for a few such errors, and refuses to apply the patch)
//	error-all (is similar to error but shows all errors)
func ActionWhitespaceModes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"nowarn", "turns off the trailing whitespace warning",
		"warn", "outputs warnings for a few such errors, but applies the patch as-is (default)",
		"fix", "outputs warnings for a few such errors, and applies the patch after fixing them",
		"error", "outputs warnings for a few such errors, and refuses to apply the patch",
		"error-all", "is similar to error but shows all errors",
	)
}
