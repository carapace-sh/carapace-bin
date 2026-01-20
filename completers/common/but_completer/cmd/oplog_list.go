package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/but"
	"github.com/spf13/cobra"
)

var oplog_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List operation history",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(oplog_listCmd).Standalone()

	oplog_listCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	oplog_listCmd.Flags().String("since", "", "Start from this oplog SHA instead of the head")
	oplog_listCmd.Flags().BoolP("snapshot", "s", false, "Show only on-demand snapshot entries")
	oplogCmd.AddCommand(oplog_listCmd)

	carapace.Gen(oplog_listCmd).FlagCompletion(carapace.ActionMap{
		"since": but.ActionOplogEntries(),
	})
}
