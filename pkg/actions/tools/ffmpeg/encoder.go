package ffmpeg

import (
	"regexp"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

type EncoderOpts struct {
	Audio    bool
	Subtitle bool
	Video    bool
}

func (o EncoderOpts) Default() EncoderOpts {
	o.Audio = true
	o.Subtitle = true
	o.Video = true
	return o
}

// ActionEncoders completes encoders
//
//	ac3 (ATSC A/52A (AC-3))
//	ac3_fixed (ATSC A/52A (AC-3) (codec ac3))
func ActionEncoders(opts EncoderOpts) carapace.Action {
	return carapace.ActionExecCommand("ffmpeg", "-hide_banner", "-encoders")(func(output []byte) carapace.Action {
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
	}).Tag("encoders")
}
