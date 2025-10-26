package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var iterationCmd = &cobra.Command{
	Use:   "iteration <command> [flags]",
	Short: "Retrieve iteration information.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iterationCmd).Standalone()

	iterationCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository. Can use either `OWNER/REPO` or `GROUP/NAMESPACE/REPO` format. Also accepts full URL or Git URL.")
	rootCmd.AddCommand(iterationCmd)

	carapace.Gen(iterationCmd).FlagCompletion(carapace.ActionMap{
		"repo": action.ActionRepo(iterationCmd),
	})
}
