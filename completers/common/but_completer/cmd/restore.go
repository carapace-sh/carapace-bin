package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/but"
	"github.com/spf13/cobra"
)

var restoreCmd = &cobra.Command{
	Use:     "restore",
	Short:   "Restore to a specific oplog snapshot",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: "operation history",
}

func init() {
	carapace.Gen(restoreCmd).Standalone()

	restoreCmd.Flags().BoolP("force", "f", false, "Skip confirmation prompt")
	restoreCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(restoreCmd)

	carapace.Gen(restoreCmd).PositionalCompletion(
		but.ActionOplogEntries(),
	)
}
