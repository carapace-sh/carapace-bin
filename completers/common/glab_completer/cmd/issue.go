package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var issueCmd = &cobra.Command{
	Use:   "issue [command] [flags]",
	Short: "Work with GitLab issues.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(issueCmd).Standalone()

	issueCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository. Can use either `OWNER/REPO` or `GROUP/NAMESPACE/REPO` format. Also accepts full URL or Git URL.")
	rootCmd.AddCommand(issueCmd)

	carapace.Gen(issueCmd).FlagCompletion(carapace.ActionMap{
		"repo": action.ActionRepo(issueCmd),
	})
}
