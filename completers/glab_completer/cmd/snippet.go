package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var snippetCmd = &cobra.Command{
	Use:   "snippet <command> [flags]",
	Short: "Create, view and manage snippets",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(snippetCmd).Standalone()

	snippetCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository using the `OWNER/REPO` or `GROUP/NAMESPACE/REPO` format or full URL or git URL")
	rootCmd.AddCommand(snippetCmd)

	carapace.Gen(snippetCmd).FlagCompletion(carapace.ActionMap{
		"repo": action.ActionRepo(snippetCmd),
	})
}
