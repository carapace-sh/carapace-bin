package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
	"github.com/spf13/cobra"
)

var flake_lockCmd = &cobra.Command{
	Use:   "lock [flags] [flake-url]",
	Short: "create missing lockfile entries",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(flake_lockCmd).Standalone()

	flake_lockCmd.Flags().Bool("commit-lock-file", false, "Commit changes to the flake's lock file")
	flake_lockCmd.Flags().String("inputs-from", "", "Use the inputs of the specified flake as registry entries")
	flake_lockCmd.Flags().Bool("no-registries", false, "Don't allow lookups in the flake registries")
	flake_lockCmd.Flags().Bool("no-update-lock-file", false, "Do not allow any updates to the flake's lock file")
	flake_lockCmd.Flags().String("output-lock-file", "", "Write the given lock file instead of flake.lock")
	flake_lockCmd.Flags().String("override-input", "", "Override a specific flake input (e.g. dwarffs/nixpkgs)")
	flake_lockCmd.Flags().Bool("recreate-lock-file", false, "Recreate the flake's lock file from scratch")
	flake_lockCmd.Flags().String("reference-lock-file", "", "Read the given lock file instead of flake.lock")
	flake_lockCmd.Flags().String("update-input", "", "Update a specific flake input (ignoring its previous entry in the lock file")

	flake_lockCmd.Flag("override-input").Nargs = 2

	addEvaluationFlags(flake_lockCmd)
	addLoggingFlags(flake_lockCmd)

	carapace.Gen(flake_lockCmd).FlagCompletion(carapace.ActionMap{
		"inputs-from": carapace.Batch(
			carapace.ActionDirectories(),
			nix.ActionFlakes(),
		).ToA(),
		"output-lock-file":    carapace.ActionFiles(),
		"reference-lock-file": carapace.ActionFiles("lock"),
	})
	carapace.Gen(flake_lockCmd).PositionalCompletion(carapace.Batch(
		carapace.ActionDirectories(),
		nix.ActionFlakes(),
	).ToA())

	flakeCmd.AddCommand(flake_lockCmd)
}
