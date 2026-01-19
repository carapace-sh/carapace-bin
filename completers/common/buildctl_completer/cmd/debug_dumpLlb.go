package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_dumpLlbCmd = &cobra.Command{
	Use:   "dump-llb",
	Short: "dump LLB in human-readable format",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_dumpLlbCmd).Standalone()

	debug_dumpLlbCmd.Flags().Bool("dot", false, "Output dot format")
	debugCmd.AddCommand(debug_dumpLlbCmd)

	carapace.Gen(debug_dumpLlbCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
