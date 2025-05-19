package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var jobCmd = &cobra.Command{
	Use:   "job <command> [flags]",
	Short: "Work with GitLab CI/CD jobs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(jobCmd).Standalone()

	jobCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository. Can use either `OWNER/REPO` or `GROUP/NAMESPACE/REPO` format. Also accepts full URL or Git URL.")
	rootCmd.AddCommand(jobCmd)

	carapace.Gen(jobCmd).FlagCompletion(carapace.ActionMap{
		"repo": action.ActionRepo(jobCmd),
	})
}
