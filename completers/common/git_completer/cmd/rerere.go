package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rerereCmd = &cobra.Command{
	Use:     "rerere",
	Short:   "Perform merge without touching index or working tree",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_interrogator].ID,
}

func init() {
	carapace.Gen(rerereCmd).Standalone()

	rerereCmd.Flags().Bool("rerere-autoupdate", false, "register clean resolutions in index")
	rootCmd.AddCommand(rerereCmd)

	carapace.Gen(rerereCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"clear", "Reset the metadata used by rerere if a merge resolution is to be aborted",
			"forget", "Reset the conflict resolutions which rerere has recorded for the current conflict in <pathspec>",
			"diff", "Display diffs for the current state of the resolution",
			"status", "Print paths with conflicts whose merge resolution rerere will record",
			"remaining", "Print paths with conflicts that have not been autoresolved by rerere",
			"gc", "Prune records of conflicted merges that occurred a long time ago",
		),
		// TODO forget <pathspec>
	)
}
