package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/but"
	"github.com/spf13/cobra"
)

var oplog_restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restore to a specific oplog snapshot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(oplog_restoreCmd).Standalone()

	oplog_restoreCmd.Flags().BoolP("force", "f", false, "Skip confirmation prompt")
	oplog_restoreCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	oplogCmd.AddCommand(oplog_restoreCmd)

	carapace.Gen(oplog_restoreCmd).PositionalCompletion(
		but.ActionOplogEntries(),
	)
}
