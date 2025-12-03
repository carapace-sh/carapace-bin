package jj

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net/ssh"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
)

// ActionSigningKeys completes signing keys based on the user's configuration
func ActionSigningKeys() carapace.Action {
	return actionExecJJ("config", "get", "signing.backend")(func(output []byte) carapace.Action {
		backend := strings.TrimSpace(string(output))

		if backend == "ssh" {
			return ssh.ActionSigningKeys()
		} else if backend == "gpg" {
			return os.ActionGpgKeyIds()
		}

		// TODO: complete gpg signing keys
		return carapace.ActionValues()
	})
}
