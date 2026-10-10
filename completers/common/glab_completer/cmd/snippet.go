package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var snippetCmd = &cobra.Command{
	Use:   "snippet <command> [flags]",
	Short: "Create snippets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(snippetCmd).Standalone()

	snippetCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository. You can use either OWNER/REPO or GROUP/NAMESPACE/REPO. The full URL or Git URL is also accepted.")
	rootCmd.AddCommand(snippetCmd)

	carapace.Gen(snippetCmd).FlagCompletion(carapace.ActionMap{
		"repo": action.ActionRepo(snippetCmd),
	})
}
