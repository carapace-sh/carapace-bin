package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mdls",
	Short: "list Spotlight metadata attributes",
	Long:  "https://keith.github.io/xcode-manpages/mdls.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("help", "h", false, "Display usage information")
	rootCmd.Flags().StringP("name", "name", "", "Get the specified attribute")
	rootCmd.Flags().BoolP("plist", "plist", false, "Print in plist format")
	rootCmd.Flags().BoolP("raw", "raw", false, "Print raw attribute data")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
