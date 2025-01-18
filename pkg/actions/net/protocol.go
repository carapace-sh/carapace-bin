package net

import "github.com/carapace-sh/carapace"

// ActionProtocols completes protocols
//
//	ftp (File Transfer Protocol)
//	http (Hypertext Transfer Protocol)
func ActionProtocols() carapace.Action {
	return carapace.ActionValuesDescribed(
		"ftp", "File Transfer Protocol",
		"http", "Hypertext Transfer Protocol",
		"https", "Hypertext Transfer Protocol Secure",
		"imap", "Internet Message Access Protocol",
		"nntp", "Network News Transfer Protocol",
	).Tag("protocols")
}
