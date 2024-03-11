package mpv

import (
	"regexp"
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionAudioDecoders completes audio decoders
//
//	8svx_exp (8SVX exponential)
//	8svx_fib (8SVX fibonacci)
func ActionAudioDecoders() carapace.Action {
	return carapace.ActionExecCommand("mpv", "--ad=help")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		r := regexp.MustCompile(`^ +(?P<name>[^ ]+)[^-]+- (?P<description>.*)$`)

		vals := make([]string, 0)
		for _, line := range lines {
			if matches := r.FindStringSubmatch(line); matches != nil {
				vals = append(vals, matches[1], matches[2])
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}

// ActionVideoDecoders completes video decoders
//
//	012v (Uncompressed 4:2:2 10-bit)
//	4xm (4X Movie)
func ActionVideoDecoders() carapace.Action {
	return carapace.ActionExecCommand("mpv", "--vd=help")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		r := regexp.MustCompile(`^ +(?P<name>[^ ]+)[^-]+- (?P<description>.*)$`)

		vals := make([]string, 0)
		for _, line := range lines {
			if matches := r.FindStringSubmatch(line); matches != nil {
				vals = append(vals, matches[1], matches[2])
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
