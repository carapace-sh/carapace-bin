package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var milestoneCmd = &cobra.Command{
	Use:   "milestone <command>",
	Short: "Manage group or project milestones.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(milestoneCmd).Standalone()

	milestoneCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository. Can use either `OWNER/REPO` or `GROUP/NAMESPACE/REPO` format. Also accepts full URL or Git URL.")
	rootCmd.AddCommand(milestoneCmd)

	carapace.Gen(milestoneCmd).FlagCompletion(carapace.ActionMap{
		"repo": action.ActionRepo(milestoneCmd),
	})
}
