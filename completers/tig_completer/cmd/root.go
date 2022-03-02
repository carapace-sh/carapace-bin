package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tig",
	Short: "text-mode interface for Git",
	Long:  "https://jonas.github.io/tig/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("C", "C", "", "Start in <path>")
	rootCmd.Flags().BoolP("help", "h", false, "Show help message and exit")
	rootCmd.Flags().BoolP("version", "v", false, "Show version and exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"C": carapace.ActionDirectories(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		git.ActionRefs(git.RefOptionDefault),
	)
}
