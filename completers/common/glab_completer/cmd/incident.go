package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var incidentCmd = &cobra.Command{
	Use:   "incident [command] [flags]",
	Short: "Work with GitLab incidents.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(incidentCmd).Standalone()

	incidentCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository. Can use either `OWNER/REPO` or `GROUP/NAMESPACE/REPO` format. Also accepts full URL or Git URL.")
	rootCmd.AddCommand(incidentCmd)

	carapace.Gen(incidentCmd).FlagCompletion(carapace.ActionMap{
		"repo": action.ActionRepo(incidentCmd),
	})
}
