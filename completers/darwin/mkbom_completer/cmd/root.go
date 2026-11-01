package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mkbom",
	Short: "create a bill-of-materials file",
	Long:  "https://keith.github.io/xcode-manpages/mkbom.8.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("h", "h", false, "Print full usage")
	rootCmd.Flags().StringS("i", "i", "", "Use the information in filelist to construct the bom file")
	rootCmd.Flags().BoolS("s", "s", false, "Create a simplified bom containing only file paths")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"i": carapace.ActionFiles(),
	})
}
