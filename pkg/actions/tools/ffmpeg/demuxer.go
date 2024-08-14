package ffmpeg

import (
	"regexp"
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionDemuxers completes demuxers
//
//	aax (CRI AAX)
//	ac3 (raw AC-3)
func ActionDemuxers() carapace.Action {
	return carapace.ActionExecCommand("ffmpeg", "-hide_banner", "-demuxers")(func(output []byte) carapace.Action {
		_, content, ok := strings.Cut(string(output), " ---")
		if !ok {
			return carapace.ActionMessage("failed to parse demuxers")
		}

		lines := strings.Split(content, "\n")
		r := regexp.MustCompile(`^.{5}(?P<name>[^ ]+) +(?P<description>.*)$`)

		vals := make([]string, 0)
		for _, line := range lines {
			if matches := r.FindStringSubmatch(line); matches != nil {
				vals = append(vals, matches[1], matches[2])
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Tag("demuxers")
}
