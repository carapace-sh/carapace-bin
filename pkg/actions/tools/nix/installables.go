package nix

import "github.com/carapace-sh/carapace"

// ActionInstallables completes nix flake installables.
//
// Installable completion intentionally relies on the flake registry rather than
// legacy nix channels so it works on flake-based systems with channels disabled.
func ActionInstallables() carapace.Action {
	return ActionFlakeRefs()
}
