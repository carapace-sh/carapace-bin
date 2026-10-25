package ytdlp

import "github.com/carapace-sh/carapace"

// ActionProtocols completes protocols
//
//	http
//	ftp
func ActionProtocols() carapace.Action {
	return carapace.ActionValues(
		"dash",
		"ftp",
		"http",
		"m3u8",
		"mms",
		"rstp",
		"rtmp",
	).Tag("protocols")
}
