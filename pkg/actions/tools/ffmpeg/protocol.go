package ffmpeg

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

// ActionProtocols completes protocols
//
//	concatf
//	crypto
func ActionProtocols() carapace.Action {
	return carapace.ActionExecCommand("ffmpeg", "-hide_banner", "-protocols")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")

		found := false
		vals := make([]string, 0)
		for _, line := range lines[2 : len(lines)-1] {
			if !found && line == "Output:" {
				found = true
				continue
			}

			switch found {
			case true:
				vals = append(vals, strings.TrimSpace(line), style.Yellow) // output
			default:
				vals = append(vals, strings.TrimSpace(line), style.Blue) // input
			}

		}
		return carapace.ActionStyledValues(vals...)
	}).Tag("protocols")
}
