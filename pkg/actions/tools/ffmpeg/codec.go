package ffmpeg

import (
	"regexp"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

type CodecOpts struct {
	Audio    bool
	Subtitle bool
	Video    bool
	Decoding bool
	Encoding bool
}

func (o CodecOpts) Default() CodecOpts {
	o.Audio = true
	o.Subtitle = true
	o.Video = true
	o.Decoding = true
	o.Encoding = true
	return o
}

// ActionCodecs completes codecs
//
//	4gv (4GV (Fourth Generation Vocoder))
//	4xm (4X Movie)
func ActionCodecs(opts CodecOpts) carapace.Action {
	return carapace.ActionExecCommand("ffmpeg", "-hide_banner", "-codecs")(func(output []byte) carapace.Action {
		_, content, ok := strings.Cut(string(output), " -------")
		if !ok {
			return carapace.ActionMessage("failed to parse codecs")
		}

		lines := strings.Split(content, "\n")
		r := regexp.MustCompile(`^ .(?P<decoding>.)(?P<encoding>.)(?P<type>.).{3} (?P<codec>[^ ]+) +(?P<description>.*)$`)

		vals := make([]string, 0)
		for _, line := range lines[10 : len(lines)-1] {
			if matches := r.FindStringSubmatch(line); matches != nil {
				decoding := matches[1] == "D"
				encoding := matches[2] == "E"

				if opts.Decoding && opts.Encoding {
					if !(decoding || encoding) {
						continue
					}
				} else {
					if opts.Decoding && !decoding {
						continue
					}
					if opts.Encoding && !encoding {
						continue
					}
				}

				switch matches[3] {
				case "A":
					if opts.Audio {
						vals = append(vals, matches[4], matches[5], style.Yellow)
					}
				case "S":
					if opts.Subtitle {
						vals = append(vals, matches[4], matches[5], style.Magenta)
					}
				case "V":
					if opts.Video {
						vals = append(vals, matches[4], matches[5], style.Blue)
					}
				}
			}
		}
		if opts.Decoding {
			vals = append(vals, "copy", "copy the codec of the input", style.Default)
		}
		return carapace.ActionStyledValuesDescribed(vals...)
	}).Tag("codecs")
}
