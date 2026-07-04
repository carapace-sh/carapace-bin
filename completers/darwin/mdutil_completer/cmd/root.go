package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mdutil",
	Short: "manage Spotlight indexing",
	Long:  "https://keith.github.io/xcode-manpages/mdutil.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("erase", "E", false, "Erase the Spotlight index on the volume")
	rootCmd.Flags().BoolP("help", "h", false, "Display usage information")
	rootCmd.Flags().BoolP("index", "i", false, "Turn on/off indexing on the volume")
	rootCmd.Flags().BoolP("status", "s", false, "Print indexing status of the volume")
	rootCmd.Flags().BoolP("verbose", "v", false, "Verbose mode")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionDirectories())
}
