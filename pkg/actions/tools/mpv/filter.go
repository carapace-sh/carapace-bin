package mpv

import (
	"regexp"
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionAudioFilters completes audio filters
//
//	abench (Benchmark part of a filtergraph.)
//	acompressor (Audio compressor.)
func ActionAudioFilters() carapace.Action {
	// TODO add value completion (`FILTER=VALUE`) and split up action internally. see `mpv --af=FILTER=help`
	return carapace.ActionExecCommand("mpv", "--af=help")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		r := regexp.MustCompile(`^ +(?P<name>[^ ]+) +(?P<description>.*)$`)

		vals := make([]string, 0)
		for _, line := range lines {
			if matches := r.FindStringSubmatch(line); matches != nil {
				vals = append(vals, matches[1], matches[2])
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}

// ActionVideoFilters completes video filters
//
//	addroi (Add region of interest to frame.)
//	alphaextract (Extract an alpha channel as a grayscale image component.)
func ActionVideoFilters() carapace.Action {
	// TODO add value completion (`FILTER=VALUE`) and split up action internally. see `mpv --vf=FILTER=help`
	return carapace.ActionExecCommand("mpv", "--vf=help")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		r := regexp.MustCompile(`^ +(?P<name>[^ ]+) +(?P<description>.*)$`)

		vals := make([]string, 0)
		for _, line := range lines {
			if matches := r.FindStringSubmatch(line); matches != nil {
				vals = append(vals, matches[1], matches[2])
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
