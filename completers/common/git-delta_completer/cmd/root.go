package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "git-delta",
	Short: "Lists files that differ from another branch",
	Long:  "https://github.com/tj/git-extras/blob/master/Commands.md#git-delta",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("help", false, "show help")

	carapace.Gen(rootCmd).PositionalCompletion(
		git.ActionLocalBranches(),
		carapace.ActionValuesDescribed(
			"M", "Modified",
			"T", "Type changed",
			"R", "Renamed",
			"D", "Deleted",
			"A", "Added",
			"C", "Copied",
			"U", "Unmerged",
			"X", "Unknown",
			"B", "Broken pairing",
		),
	)
}
