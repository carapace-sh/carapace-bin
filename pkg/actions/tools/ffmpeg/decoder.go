package ffmpeg

import (
	"regexp"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

type DecoderOpts struct {
	Audio    bool
	Subtitle bool
	Video    bool
}

func (o DecoderOpts) Default() DecoderOpts {
	o.Audio = true
	o.Subtitle = true
	o.Video = true
	return o
}

// ActionDecoders completes decoders
//
//	4xm (4X Movie)
//	8bps (QuickTime 8BPS video)
func ActionDecoders(opts DecoderOpts) carapace.Action {
	return carapace.ActionExecCommand("ffmpeg", "-hide_banner", "-decoders")(func(output []byte) carapace.Action {
		_, content, ok := strings.Cut(string(output), " ------")
		if !ok {
			return carapace.ActionMessage("failed to parse encoders")
		}

		lines := strings.Split(content, "\n")
		r := regexp.MustCompile(`^ (?P<type>.).{5} (?P<name>[^ ]+) +(?P<description>.*)$`)

		vals := make([]string, 0)
		for _, line := range lines {
			if matches := r.FindStringSubmatch(line); matches != nil {
				switch matches[1] {
				case "A":
					if opts.Audio {
						vals = append(vals, matches[2], matches[3], style.Yellow)
					}
				case "S":
					if opts.Subtitle {
						vals = append(vals, matches[2], matches[3], style.Magenta)
					}
				case "V":
					if opts.Video {
						vals = append(vals, matches[2], matches[3], style.Blue)
					}
				}
			}
		}
		return carapace.ActionStyledValuesDescribed(vals...)
	}).Tag("decoders")
}
