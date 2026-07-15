package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/herdr_completer/cmd/action"
	"github.com/spf13/cobra"
)

var wait_outputCmd = &cobra.Command{
	Use:   "output <pane_id>",
	Short: "Wait for matching pane output",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wait_outputCmd).Standalone()

	wait_outputCmd.Flags().String("lines", "", "")
	wait_outputCmd.Flags().String("match", "", "")
	wait_outputCmd.Flags().Bool("raw", false, "")
	wait_outputCmd.Flags().Bool("regex", false, "")
	wait_outputCmd.Flags().String("source", "", "")
	wait_outputCmd.Flags().String("timeout", "", "")
	waitCmd.AddCommand(wait_outputCmd)

	carapace.Gen(wait_outputCmd).PositionalCompletion(action.ActionPanes(wait_outputCmd))

	carapace.Gen(wait_outputCmd).FlagCompletion(carapace.ActionMap{
		"source": carapace.ActionValues("visible", "recent", "recent-unwrapped"),
	})
}
