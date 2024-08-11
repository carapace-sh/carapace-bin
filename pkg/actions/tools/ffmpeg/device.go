package ffmpeg

import (
	"regexp"
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionDevices completes devices
//
//	alsa (ALSA audio output)
//	fbdev (Linux framebuffer)
func ActionDevices(deviceType rune) carapace.Action {
	return carapace.ActionExecCommand("ffmpeg", "-hide_banner", "-devices")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		r := regexp.MustCompile(`^ (?P<prefix>.{2}) (?P<devices>[^ ]+) +(?P<description>.*)$`)

		vals := make([]string, 0)
		for _, line := range lines[4 : len(lines)-1] {
			if r.MatchString(line) {
				matches := r.FindStringSubmatch(line)
				if deviceType == '\u0000' || strings.IndexByte(matches[1], byte(deviceType)) != -1 {
					vals = append(vals, matches[2], matches[3])
				}
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
