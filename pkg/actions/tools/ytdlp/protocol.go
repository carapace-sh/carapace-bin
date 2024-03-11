package ytdlp

import "github.com/carapace-sh/carapace"

// ActionProtocols completes protocols
//
//	http
//	ftp
func ActionProtocols() carapace.Action {
	return carapace.ActionValues(
		"http",
		"ftp",
		"m3u8",
		"dash",
		"rstp",
		"rtmp",
		"mms",
	).Tag("protocols")
}
