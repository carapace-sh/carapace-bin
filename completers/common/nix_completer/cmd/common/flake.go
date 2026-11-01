package common

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
	"github.com/spf13/cobra"
)

// AddFlakeFlags adds common flake-related flags
// (MixFlakeOptions in nix source).
func AddFlakeFlags(cmd *cobra.Command) {
	cmd.Flags().Bool("commit-lock-file", false, "Commit changes to the flake's lock file")
	cmd.Flags().String("inputs-from", "", "Use the inputs of the specified flake as registry entries")
	cmd.Flags().Bool("no-registries", false, "Don't allow lookups in the flake registries")
	cmd.Flags().Bool("no-update-lock-file", false, "Do not allow any updates to the flake's lock file")
	cmd.Flags().Bool("no-write-lock-file", false, "Do not write the flake's newly generated lock file")
	cmd.Flags().String("output-lock-file", "", "Write the given lock file instead of flake.lock")
	cmd.Flags().String("override-input", "", "Override a specific flake input (e.g. dwarffs/nixpkgs)")
	cmd.Flags().Bool("recreate-lock-file", false, "Recreate the flake's lock file from scratch")
	cmd.Flags().String("reference-lock-file", "", "Read the given lock file instead of flake.lock")
	cmd.Flags().String("update-input", "", "Update a specific flake input (ignoring its previous entry in the lock file")

	cmd.Flag("override-input").Nargs = 2

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"inputs-from": carapace.Batch(
			carapace.ActionDirectories(),
			nix.ActionFlakes(),
		).ToA(),
		"output-lock-file":    carapace.ActionFiles(),
		"reference-lock-file": carapace.ActionFiles("lock"),
	})
}
