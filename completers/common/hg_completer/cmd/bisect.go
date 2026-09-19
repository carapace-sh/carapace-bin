package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/hg_completer/cmd/action"
	"github.com/spf13/cobra"
)

var bisectCmd = &cobra.Command{
	Use:     "bisect",
	Short:   "subdivision search of changesets",
	GroupID: groups[group_change_navigation].ID,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bisectCmd).Standalone()

	bisectCmd.Flags().BoolP("bad", "b", false, "mark changeset bad")
	bisectCmd.Flags().StringP("command", "c", "", "use command to check changeset state")
	bisectCmd.Flags().BoolP("extend", "e", false, "extend the bisect range")
	bisectCmd.Flags().BoolP("good", "g", false, "mark changeset good")
	bisectCmd.Flags().BoolP("noupdate", "U", false, "do not update to target")
	bisectCmd.Flags().BoolP("reset", "r", false, "reset bisect state")
	bisectCmd.Flags().BoolP("skip", "s", false, "skip testing changeset")
	rootCmd.AddCommand(bisectCmd)

	carapace.Gen(bisectCmd).PositionalCompletion(
		action.ActionRevisions(),
	)
}
