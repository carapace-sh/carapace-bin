package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mdimport",
	Short: "import files into Spotlight index",
	Long:  "https://keith.github.io/xcode-manpages/mdimport.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("help", "h", false, "Display usage information")
	rootCmd.Flags().BoolP("list", "A", false, "List all installed metadata importers")
	rootCmd.Flags().BoolP("test", "t", false, "Test the metadata importer")
	rootCmd.Flags().BoolP("verbose", "v", false, "Verbose mode")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
