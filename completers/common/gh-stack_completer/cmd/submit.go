package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var submitCmd = &cobra.Command{
	Use:   "submit",
	Short: "Create a stack of PRs on GitHub",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(submitCmd).Standalone()

	submitCmd.Flags().Bool("auto", false, "Use auto-generated PR titles without prompting")
	submitCmd.Flags().Bool("draft", false, "Create PRs as drafts")
	submitCmd.Flags().String("remote", "", "Remote to push to (defaults to auto-detected remote)")
	rootCmd.AddCommand(submitCmd)

	carapace.Gen(submitCmd).FlagCompletion(carapace.ActionMap{
		"remote": git.ActionRemotes(),
	})
}
