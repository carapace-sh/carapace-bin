package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var queueRetryBatchCmd = &cobra.Command{
	Use:   "queue:retry-batch [--isolated [ISOLATED]] [--] <id>",
	Short: "Retry the failed jobs for a batch",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(queueRetryBatchCmd).Standalone()

	queueRetryBatchCmd.Flags().String("isolated", "", "Do not run the command if another instance of the command is already running")
	rootCmd.AddCommand(queueRetryBatchCmd)
}
