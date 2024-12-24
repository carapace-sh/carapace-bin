package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
	"github.com/spf13/cobra"
)

var flake_updateCmd = &cobra.Command{
	Use:   "update [flags] [flake-url]",
	Short: "update flake lock file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(flake_updateCmd).Standalone()

	flake_updateCmd.Flags().Bool("commit-lock-file", false, "Commit changes to the flake's lock file")
	flake_updateCmd.Flags().String("inputs-from", "", "Use the inputs of the specified flake as registry entries")
	flake_updateCmd.Flags().Bool("no-registries", false, "Don't allow lookups in the flake registries")
	flake_updateCmd.Flags().String("output-lock-file", "", "Write the given lock file instead of flake.lock")
	flake_updateCmd.Flags().String("override-input", "", "Override a specific flake input (e.g. dwarffs/nixpkgs)")
	flake_updateCmd.Flags().String("reference-lock-file", "", "Read the given lock file instead of flake.lock")

	flake_updateCmd.Flag("override-input").Nargs = 2

	addEvaluationFlags(flake_updateCmd)
	addLoggingFlags(flake_updateCmd)

	carapace.Gen(flake_updateCmd).FlagCompletion(carapace.ActionMap{
		"inputs-from": carapace.Batch(
			carapace.ActionDirectories(),
			nix.ActionFlakes(),
		).ToA(),
		"output-lock-file":    carapace.ActionFiles(),
		"reference-lock-file": carapace.ActionFiles("lock"),
	})
	carapace.Gen(flake_updateCmd).PositionalCompletion(carapace.Batch(
		carapace.ActionDirectories(),
		nix.ActionFlakes(),
	).ToA())

	flakeCmd.AddCommand(flake_updateCmd)
}
