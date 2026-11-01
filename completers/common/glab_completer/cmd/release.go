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

	releaseCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository. You can use either OWNER/REPO or GROUP/NAMESPACE/REPO. The full URL or Git URL is also accepted.")
	rootCmd.AddCommand(releaseCmd)

	carapace.Gen(releaseCmd).FlagCompletion(carapace.ActionMap{
		"repo": action.ActionRepo(releaseCmd),
	})
}
