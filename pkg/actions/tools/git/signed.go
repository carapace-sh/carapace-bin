package git

import "github.com/carapace-sh/carapace"

// ActionSignedModes completes signed modes for commits and tags
//
//	verbatim (store signatures as-is)
//	warn-verbatim (warn if signature does not verify and store as-is)
func ActionSignedModes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"abort", "abort if signature does not verify",
		"sign-if-invalid", "replace invalid signatures with newly created ones",
		"strip", "strip any signature",
		"strip-if-invalid", "strip only invalid signatures",
		"verbatim", "store signatures as-is",
		"warn-strip", "warn if signature does not verify and strip it",
		"warn-verbatim", "warn if signature does not verify and store as-is",
	).Tag("signed modes").Uid("git", "signed-mode")
}
