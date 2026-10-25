package git

import "github.com/carapace-sh/carapace"

// ActionWhitespaceModes completes whitespace modes
//
//	error (outputs warnings for a few such errors, and refuses to apply the patch)
//	error-all (is similar to error but shows all errors)
func ActionWhitespaceModes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"error", "outputs warnings for a few such errors, and refuses to apply the patch",
		"error-all", "is similar to error but shows all errors",
		"fix", "outputs warnings for a few such errors, and applies the patch after fixing them",
		"nowarn", "turns off the trailing whitespace warning",
		"warn", "outputs warnings for a few such errors, but applies the patch as-is (default)",
	).Uid("git", "whitespace-mode")
}

// ActionWhitespaceProblems completes whitespace problems
//
//	blank-at-eof (treats blank lines added at the end of file as an error)
//	trailing-space (is a short-hand to cover both blank-at-eol and blank-at-eof)
func ActionWhitespaceProblems() carapace.Action {
	return carapace.ActionValuesDescribed(
		"blank-at-eof", "treats blank lines added at the end of file as an error",
		"blank-at-eol", "treats trailing whitespaces at the end of the line as an error",
		"cr-at-eol", "treats a carriage-return at the end of line as part of the line terminator",
		"indent-with-non-tab", "treats a line that is indented with space characters instead of the equivalent tabs as an error",
		"space-before-tab", "treats a space character that appears immediately before a tab character in the initial indent part of the line as an error",
		"tab-in-indent", "treats a tab character in the initial indent part of the line as an error",
		"tabwidth", "tells how many character positions a tab occupies",
		"trailing-space", "is a short-hand to cover both blank-at-eol and blank-at-eof",
	).Uid("git", "whitespace-problem")
}
