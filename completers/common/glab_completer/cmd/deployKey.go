package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var deployKeyCmd = &cobra.Command{
	Use:   "deploy-key <command>",
	Short: "Manage deploy keys.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deployKeyCmd).Standalone()

	deployKeyCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository. Can use either `OWNER/REPO` or `GROUP/NAMESPACE/REPO` format. Also accepts full URL or Git URL.")
	rootCmd.AddCommand(deployKeyCmd)

	carapace.Gen(deployKeyCmd).FlagCompletion(carapace.ActionMap{
		"repo": action.ActionRepo(deployKeyCmd),
	})
}
