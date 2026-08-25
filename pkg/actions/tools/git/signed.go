package git

import "github.com/carapace-sh/carapace"

// ActionSignedModes completes signed modes for commits and tags
//
//	verbatim (store signatures as-is)
//	warn-verbatim (warn if signature does not verify and store as-is)
func ActionSignedModes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"verbatim", "store signatures as-is",
		"warn-verbatim", "warn if signature does not verify and store as-is",
		"warn-strip", "warn if signature does not verify and strip it",
		"strip", "strip any signature",
		"strip-if-invalid", "strip only invalid signatures",
		"sign-if-invalid", "replace invalid signatures with newly created ones",
		"abort", "abort if signature does not verify",
	).Tag("signed modes").Uid("git", "signed-mode")
}
