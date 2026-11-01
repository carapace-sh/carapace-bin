package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var tokenCmd = &cobra.Command{
	Use:     "token",
	Short:   "Manage personal, project, or group tokens.",
	Aliases: []string{"token"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tokenCmd).Standalone()

	tokenCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository. You can use either OWNER/REPO or GROUP/NAMESPACE/REPO. The full URL or Git URL is also accepted.")
	rootCmd.AddCommand(tokenCmd)

	carapace.Gen(tokenCmd).FlagCompletion(carapace.ActionMap{
		"repo": action.ActionRepo(tokenCmd),
	})
}
