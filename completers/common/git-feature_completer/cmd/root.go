package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "git-feature",
	Short: "Create/Merge a feature branch",
	Long:  "https://github.com/tj/git-extras/blob/master/Commands.md#git-feature",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var finishCmd = &cobra.Command{
	Use:   "finish",
	Short: "Merge and delete a feature branch",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("alias", "a", "", "Branch prefix, default: feature")
	rootCmd.Flags().String("from", "", "Start point for branch")
	rootCmd.Flags().Bool("help", false, "show help")
	rootCmd.Flags().StringP("remote", "r", "", "Remote tracking, defaults to origin")
	rootCmd.Flags().StringP("separator", "s", "", "Separator, default: /")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"from":   git.ActionRefs(git.RefOption{}.Default()),
		"remote": git.ActionRemotes(),
	})

	finishCmd.Flags().Bool("help", false, "show help")
	finishCmd.Flags().Bool("squash", false, "Squash merge on finish")
	carapace.Gen(finishCmd).Standalone()
	rootCmd.AddCommand(finishCmd)
}
