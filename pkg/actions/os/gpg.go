package os

import (
	"os/exec"
	"strings"

	"github.com/rsteube/carapace"
)

// ActionGpgKeyIds completes GPG key ids
//   ABCDEF1234567890 (some GPG key)
//   ABCDEF1234567891 (another GPG key)
func ActionGpgKeyIds() carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		if output, err := exec.Command("sh", "-c", "gpg --list-keys --with-colons | awk -F: '/^pub:|^uid:/ { print $5 $10}'").Output(); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			lines := strings.Split(string(output), "\n")
			return carapace.ActionValuesDescribed(lines[:len(lines)-1]...)
		}
	})
}
