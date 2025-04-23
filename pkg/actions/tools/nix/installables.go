package nix

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionInstallables completes nix packages and flakes
//
// An installable nix derivation can either be a nix package (legacyPackages)
// or a nixosModules output of a flake. There is no completion support implemented
// for flake outputs due to the processing time required, so ActionInstallables
// completes packages from the global/system/user channels available and flake names from the
// global/system/user registry
func ActionInstallables() carapace.Action {
	return carapace.ActionMultiParts("#", func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return carapace.Batch(
				ActionFlakes(),
				carapace.ActionMultiParts("/", func(c carapace.Context) carapace.Action {
					switch len(c.Parts) {
					case 0:
						return ActionLocalChannels().Suffix("#")
					case 1:
						return carapace.ActionMessage("TODO: support branch completion") // TODO support branch completion
					default:
						return carapace.ActionValues()
					}
				}),
			).ToA()

		case 1:
			switch c.Parts[0] {
			case ".":
				return ActionFlakeAttributes(".")
			default:
				return ActionPackages(strings.SplitN(c.Parts[0], "/", 2)[0])
			}

		default:
			return carapace.ActionValues()
		}
	})
}
