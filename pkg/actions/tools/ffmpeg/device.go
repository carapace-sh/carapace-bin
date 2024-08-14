package ffmpeg

import (
	"regexp"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

type DeviceOpts struct {
	Demuxing bool
	Muxing   bool
}

func (o DeviceOpts) Default() DeviceOpts {
	o.Demuxing = true
	o.Muxing = true
	return o
}

// ActionDevices completes devices
//
//	alsa (ALSA audio output)
//	fbdev (Linux framebuffer)
func ActionDevices(opts DeviceOpts) carapace.Action {
	return carapace.ActionExecCommand("ffmpeg", "-hide_banner", "-devices")(func(output []byte) carapace.Action {
		_, content, ok := strings.Cut(string(output), " ---")
		if !ok {
			return carapace.ActionMessage("failed to parse devices")
		}

		lines := strings.Split(content, "\n")
		r := regexp.MustCompile(`^ (?P<demuxer>.)(?P<muxer>.) (?P<devices>[^ ]+) +(?P<description>.*)$`)

		vals := make([]string, 0)
		for _, line := range lines[:len(lines)-1] {
			if matches := r.FindStringSubmatch(line); matches != nil {
				demuxer := matches[1] == "D" && opts.Demuxing
				muxer := matches[2] == "E" && opts.Muxing
				switch {
				case demuxer && muxer:
					vals = append(vals, matches[3], matches[4], style.Magenta)
				case demuxer:
					vals = append(vals, matches[3], matches[4], style.Blue)
				case muxer:
					vals = append(vals, matches[3], matches[4], style.Yellow)
				}
			}
		}
		return carapace.ActionStyledValuesDescribed(vals...)
	}).Tag("devices")
}
