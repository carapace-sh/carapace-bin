package nix

import (
	"github.com/carapace-sh/carapace"
)

// ActionInstallables completes Nix installables from flake refs.
//
// Examples include flake references like `nixpkgs#hello`.
func ActionInstallables() carapace.Action {
	return ActionFlakeRefs()
}
