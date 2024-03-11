package nix

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionPaths completes paths
//
// A path can be one of:
//   - a local file path (default.nix)
//   - an http/https URL (https://releases.nixos.org/../nixexprs.tar.xz
//   - a channel: specifier (channel:nixos-22.05)
//   - a local channel (<nixpkgs>)
func ActionPaths() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		batch := carapace.Batch(
			ActionLocalChannels().Prefix("<").Suffix(">"),
			carapace.ActionFiles(".nix"),
		)

		if !strings.HasPrefix(c.Value, "channel:") {
			batch = append(batch, carapace.ActionValues("channel:").NoSpace(':'))
		} else {
			batch = append(batch, ActionRemoteChannels().Prefix("channel:"))
		}

		return batch.ToA()
	})
}
