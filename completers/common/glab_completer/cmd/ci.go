package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var ciCmd = &cobra.Command{
	Use:     "ci <command> [flags]",
	Short:   "Work with GitLab CI/CD pipelines and jobs.",
	Aliases: []string{"pipe", "pipeline"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ciCmd).Standalone()

	ciCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository. Can use either `OWNER/REPO` or `GROUP/NAMESPACE/REPO` format. Also accepts full URL or Git URL.")
	rootCmd.AddCommand(ciCmd)

	carapace.Gen(ciCmd).FlagCompletion(carapace.ActionMap{
		"repo": action.ActionRepo(ciCmd),
	})
}
