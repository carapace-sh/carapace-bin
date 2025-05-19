package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var releaseCmd = &cobra.Command{
	Use:   "release <command> [flags]",
	Short: "Manage GitLab releases.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(releaseCmd).Standalone()

	releaseCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository. Can use either `OWNER/REPO` or `GROUP/NAMESPACE/REPO` format. Also accepts full URL or Git URL.")
	rootCmd.AddCommand(releaseCmd)

	carapace.Gen(releaseCmd).FlagCompletion(carapace.ActionMap{
		"repo": action.ActionRepo(releaseCmd),
	})
}
