package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "git-release",
	Short: "Commit, tag and push changes to the repository",
	Long:  "https://github.com/tj/git-extras/blob/master/Commands.md#git-release",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("c", "c", false, "Generate changelog")
	rootCmd.Flags().Bool("help", false, "show help")
	rootCmd.Flags().StringS("m", "m", "", "Commit message")
	rootCmd.Flags().Bool("no-empty-commit", false, "Don't create an empty commit")
	rootCmd.Flags().String("prefix", "", "Tag prefix")
	rootCmd.Flags().StringS("r", "r", "", "Remote to push to")
	rootCmd.Flags().BoolS("s", "s", false, "Create a signed tag")
	rootCmd.Flags().String("semver", "", "Semantic version: major, minor, or patch")
	rootCmd.Flags().StringS("u", "u", "", "Sign the tag with the given key id")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"r":      git.ActionRemotes(),
		"semver": carapace.ActionValues("major", "minor", "patch"),
		"u":      os.ActionGpgKeyIds(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		git.ActionTags(),
	)
}
