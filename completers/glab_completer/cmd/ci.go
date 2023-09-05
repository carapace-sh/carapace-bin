package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var ciCmd = &cobra.Command{
	Use:     "ci <command> [flags]",
	Short:   "Work with GitLab CI/CD pipelines and jobs",
	Aliases: []string{"pipe", "pipeline"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ciCmd).Standalone()

	ciCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository using the `OWNER/REPO` or `GROUP/NAMESPACE/REPO` format or full URL or git URL")
	rootCmd.AddCommand(ciCmd)

	carapace.Gen(ciCmd).FlagCompletion(carapace.ActionMap{
		"repo": action.ActionRepo(ciCmd),
	})
}
