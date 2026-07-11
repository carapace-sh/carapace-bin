package cmd

import (
	"testing"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/sandbox"
)

func TestWgRoot(t *testing.T) {
	sandbox.Package(t, "github.com/carapace-sh/carapace-bin/completers/linux/wg_completer")(func(s *sandbox.Sandbox) {
		s.Run("").
			Expect(carapace.ActionValuesDescribed(
				"show", "Shows the current configuration and device information",
				"showconf", "Shows the current configuration of a given WireGuard interface, for use with setconf",
				"set", "Change the current configuration, add peers, remove peers, or change peers",
				"setconf", "Applies a configuration file to a WireGuard interface",
				"addconf", "Appends a configuration file to a WireGuard interface",
				"syncconf", "Synchronizes a configuration file to a WireGuard interface",
				"genkey", "Generates a new private key and writes it to stdout",
				"genpsk", "Generates a new preshared key and writes it to stdout",
				"pubkey", "Reads a private key from stdin and writes a public key to stdout",
			).Tag("commands"))
	})
}