package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var variableCmd = &cobra.Command{
	Use:     "variable",
	Short:   "Manage variables for a GitLab project or group.",
	Aliases: []string{"var"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(variableCmd).Standalone()

	variableCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository. Can use either `OWNER/REPO` or `GROUP/NAMESPACE/REPO` format. Also accepts full URL or Git URL.")
	rootCmd.AddCommand(variableCmd)

	carapace.Gen(variableCmd).FlagCompletion(carapace.ActionMap{
		"repo": action.ActionRepo(variableCmd),
	})
}
