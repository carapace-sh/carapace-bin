package ytdlp

import (
	"strings"

	"github.com/rsteube/carapace"
)

// ActionExtractors completes extractors
//
//	1News
//	1tv
func ActionExtractors() carapace.Action {
	return carapace.ActionExecCommand("yt-dlp", "--list-extractors")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		return carapace.ActionValues(lines[:len(lines)-1]...)
	}).Tag("extractors")
}
