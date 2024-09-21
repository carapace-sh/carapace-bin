package ffmpeg

import (
	"regexp"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

type CodecOpts struct {
	Attachment bool
	Audio      bool
	Data       bool
	Subtitle   bool
	Video      bool
}

func (o CodecOpts) Default() CodecOpts {
	o.Attachment = true
	o.Audio = true
	o.Data = true
	o.Subtitle = true
	o.Video = true
	return o
}

// ActionCodecs completes codecs
//
//	4gv (4GV (Fourth Generation Vocoder))
//	4xm (4X Movie)
func ActionCodecs(opts CodecOpts) carapace.Action {
	return actionCodecs(opts, nil)
}

// ActionEncodableCodecs completes codecs with encoding support
//
//	amv (AMV Video)
//	anull (Null audio codec)
func ActionEncodableCodecs(opts CodecOpts) carapace.Action {
	return actionCodecs(opts, func(s string) bool {
		return s[1] != 'E'
	})
}

// ActionDecodableCodecs completes codecs with decoding support
//
//	avrn (Avid AVI Codec)
//	avrp (Avid 1:1 10-bit RGB Packer)
func ActionDecodableCodecs(opts CodecOpts) carapace.Action {
	return actionCodecs(opts, func(s string) bool {
		return s[0] != 'D'
	})
}

func actionCodecs(opts CodecOpts, filter func(s string) bool) carapace.Action {
	return carapace.ActionExecCommand("ffmpeg", "-hide_banner", "-codecs")(func(output []byte) carapace.Action {
		_, content, ok := strings.Cut(string(output), " -------")
		if !ok {
			return carapace.ActionMessage("failed to parse codecs")
		}

		lines := strings.Split(content, "\n")
		r := regexp.MustCompile(`^ (?P<decoding>.)(?P<encoding>.)(?P<type>.).{3} (?P<codec>[^ ]+) +(?P<description>.*)$`)

		vals := make([]string, 0)
		for _, line := range lines[10 : len(lines)-1] {
			if matches := r.FindStringSubmatch(line); matches != nil {
				if filter != nil && filter(line[1:7]) {
					continue
				}

				switch matches[3] {
				case "A":
					if opts.Audio {
						vals = append(vals, matches[4], matches[5], style.Yellow)
					}
				case "D":
					if opts.Data {
						vals = append(vals, matches[4], matches[5], style.Cyan)
					}
				case "S":
					if opts.Subtitle {
						vals = append(vals, matches[4], matches[5], style.Magenta)
					}
				case "T":
					if opts.Attachment {
						vals = append(vals, matches[4], matches[5], style.Green)
					}
				case "V":
					if opts.Video {
						vals = append(vals, matches[4], matches[5], style.Blue)
					}
				}
			}
		}

		if filter == nil || !filter("D      ") {
			vals = append(vals, "copy", "copy the codec of the input", style.Default)
		}
		return carapace.ActionStyledValuesDescribed(vals...)
	}).Tag("codecs")
}
