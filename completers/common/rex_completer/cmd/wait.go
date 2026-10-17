package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rex_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/completers/common/rex_completer/cmd/common"
	"github.com/spf13/cobra"
)

var waitCmd = &cobra.Command{
	Use:   "wait TARGET",
	Short: "Wait for a block, window, or session to end",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waitCmd).Standalone()

	waitCmd.Flags().String("for", "", "give up after DURATION")
	waitCmd.Flags().Bool("json", false, "output the terminating payload as JSON")

	common.AddWindowScopeFlags(waitCmd)

	rootCmd.AddCommand(waitCmd)

	carapace.Gen(waitCmd).PositionalCompletion(
		carapace.Batch(
			action.ActionSessions(waitCmd),
			action.ActionWindows(waitCmd),
			action.ActionBlocks(waitCmd),
		).ToA(),
	)
}
