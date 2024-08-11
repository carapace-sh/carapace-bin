package ffmpeg

import (
	"regexp"
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionCodecs completes codecs
//
//	4gv (4GV (Fourth Generation Vocoder))
//	4xm (4X Movie)
func ActionCodecs(codecType rune) carapace.Action {
	return carapace.ActionExecCommand("ffmpeg", "-hide_banner", "-codecs")(func(outputCodecs []byte) carapace.Action {
		encodersCmd := carapace.ActionExecCommand("ffmpeg", "-hide_banner", "-encoders")
		return encodersCmd(func(outputEncoders []byte) carapace.Action {
			linesCodecs := strings.Split(strings.TrimSpace(string(outputCodecs)), "\n")
			linesEncoders := strings.Split(strings.TrimSpace(string(outputEncoders)), "\n")

			r := regexp.MustCompile(`^ (?P<prefix>.{6}) (?P<codec>\w+) +(?P<description>.*)$`)

			vals := make(map[string]string)

			// copy
			vals["copy"] = "copy the codec of the input"

			// Parse codecs
			for _, line := range linesCodecs {
				if r.MatchString(line) {
					matches := r.FindStringSubmatch(line)
					prefix := matches[1]
					if (codecType == '\u0000' || byte(codecType) == prefix[2]) && prefix[1] == 'E' {
						vals[matches[2]] = matches[3]
					}
				}
			}

			// Parse encoders
			for _, line := range linesEncoders {
				if r.MatchString(line) {
					matches := r.FindStringSubmatch(line)
					prefix := matches[1]
					if codecType == '\u0000' || (byte(codecType) == prefix[0]) {
						vals[matches[2]] = matches[3]
					}
				}
			}

			result := []string{}
			for codec, description := range vals {
				result = append(result, codec, description)
			}

			return carapace.ActionValuesDescribed(result...)
		})
	})
}
