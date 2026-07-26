package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "git-commits-since",
	Short: "Show commit logs since some date",
	Long:  "https://github.com/tj/git-extras/blob/master/Commands.md#git-commits-since",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("help", false, "show help")
	rootCmd.Flags().StringP("ref", "r", "", "Show commits since given ref")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"ref": git.ActionRefs(git.RefOption{}.Default()),
	})
}
