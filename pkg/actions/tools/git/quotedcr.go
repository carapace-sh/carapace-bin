package git

import "github.com/carapace-sh/carapace"

// ActionQuotedCrModes completes quoted-cr modes
//
//	nowarn (Git will do nothing when such a CRLF is found)
//	warn (Git will issue a warning for each message if such a CRLF is found)
func ActionQuotedCrModes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"nowarn", "Git will do nothing when such a CRLF is found",
		"warn", "Git will issue a warning for each message if such a CRLF is found",
		"strip", "Git will convert those CRLF to LF",
	).Tag("quoted cr modes").Uid("git", "quoted-cr-mode")
}
