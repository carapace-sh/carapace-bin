// package virtualbox contains virtualbox related actions
package virtualbox

import (
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
)

// ActionMachines completes virtualbox machines
//   machine1 (34234234-87b5-4027-a683-c239a8cbd180)
//   machine2 (32423523-87b5-4027-a683-c239a8cbd180)
func ActionMachines() carapace.Action {
	return carapace.ActionExecCommand("vboxmanage", "list", "vms")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		r := regexp.MustCompile(`^"(?P<name>[^"]+)" {(?P<chksum>.*)}$`)

		vals := make([]string, 0)
		for _, line := range lines {
			if r.MatchString(line) {
				matches := r.FindStringSubmatch(line)
				vals = append(vals, matches[1], matches[2])
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
