package cargo

import (
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
)

// ActionNightlyFlags completes nightly flags
//
//	allow-features (Allow *only* the listed unstable features)
//	avoid-dev-deps (Avoid installing dev-dependencies if possible)
func ActionNightlyFlags() carapace.Action {
	return carapace.ActionExecCommand("cargo", "-Z", "help")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		re := regexp.MustCompile(`^ +-Z (?P<name>[^ ]+) +-- (?P<description>.*)$`)

		vals := make([]string, 0)
		for _, line := range lines {
			if matches := re.FindStringSubmatch(line); matches != nil {
				vals = append(vals, matches[1], matches[2])
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
