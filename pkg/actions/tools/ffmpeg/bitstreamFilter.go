package ffmpeg

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionBitstreamFilters completes bitstream filters
//
//	dca_core
//	dts2pts
func ActionBitstreamFilters() carapace.Action {
	return carapace.ActionExecCommand("ffmpeg", "-hide_banner", "-bsfs")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")

		vals := make([]string, 0)
		for _, line := range lines[1:] {
			if line != "" {
				vals = append(vals, line)
			}
		}
		return carapace.ActionValues(vals...)
	}).Tag("bitstream filters")
}
