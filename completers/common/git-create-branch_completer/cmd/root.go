package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "git-create-branch",
	Short: "Create a local branch, optionally with a remote tracking branch",
	Long:  "https://github.com/tj/git-extras/blob/master/Commands.md#git-create-branch",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("from", "", "Start point for branch, defaults to current branch")
	rootCmd.Flags().Bool("help", false, "show help")
	rootCmd.Flags().StringP("remote", "r", "", "Setup remote tracking branch, defaults to origin")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"from":   git.ActionRefs(git.RefOption{}.Default()),
		"remote": git.ActionRemotes(),
	})
}
