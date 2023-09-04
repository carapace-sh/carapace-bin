package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var incidentCmd = &cobra.Command{
	Use:   "incident [command] [flags]",
	Short: "Work with GitLab incidents",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(incidentCmd).Standalone()

	incidentCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository using the `OWNER/REPO` or `GROUP/NAMESPACE/REPO` format or full URL or git URL")
	rootCmd.AddCommand(incidentCmd)

	carapace.Gen(incidentCmd).FlagCompletion(carapace.ActionMap{
		"repo": action.ActionRepo(incidentCmd),
	})
}
