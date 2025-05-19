package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/restic_completer/cmd/action"
	"github.com/spf13/cobra"
)

var diffCmd = &cobra.Command{
	Use:   "diff",
	Short: "Show differences between two snapshots",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(diffCmd).Standalone()
	diffCmd.Flags().Bool("metadata", false, "print changes in metadata")
	rootCmd.AddCommand(diffCmd)

	carapace.Gen(diffCmd).PositionalCompletion(
		action.ActionSnapshotIDs(diffCmd),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionSnapshotIDs(diffCmd).Invoke(c).Filter(c.Args[:1]...).ToA() // TODO use filterargs with shift
		}),
	)
}
