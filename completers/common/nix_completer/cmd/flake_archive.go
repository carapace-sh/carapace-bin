package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/nix_completer/cmd/common"
	"github.com/spf13/cobra"
)

var flake_archiveCmd = &cobra.Command{
	Use:   "archive",
	Short: "copy a flake and all its inputs to a store",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(flake_archiveCmd).Standalone()

	flake_archiveCmd.Flags().Bool("dry-run", false, "Show what this command would do without doing it.")
	flake_archiveCmd.Flags().Bool("no-check-sigs", false, "Do not require that paths are signed by trusted keys")
	flake_archiveCmd.Flags().String("to", "", "URI of the destination Nix store")
	flakeCmd.AddCommand(flake_archiveCmd)

	common.AddEvaluationFlags(flake_archiveCmd)
	common.AddFlakeFlags(flake_archiveCmd)
	common.AddLoggingFlags(flake_archiveCmd)

	carapace.Gen(flake_archiveCmd).FlagCompletion(carapace.ActionMap{})
}
