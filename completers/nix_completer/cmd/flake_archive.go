package cmd

import (
	"github.com/rsteube/carapace"
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
	flake_archiveCmd.Flags().Bool("json", false, "Produce output in JSON format")
	flake_archiveCmd.Flags().String("to", "", "URI of the destination Nix store")
	flakeCmd.AddCommand(flake_archiveCmd)

	addEvaluationFlags(flake_archiveCmd)
	addFlakeFlags(flake_archiveCmd)
	addLoggingFlags(flake_archiveCmd)
}
