package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
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

	rootCmd.Flags().BoolP("all", "a", false, "Apply command to all stores on all volumes")
	rootCmd.Flags().BoolP("disable", "d", false, "Disable Spotlight activity for volume")
	rootCmd.Flags().BoolP("erase", "E", false, "Erase the Spotlight index on the volume")
	rootCmd.Flags().BoolP("help", "h", false, "Display usage information")
	rootCmd.Flags().StringP("index", "i", "", "Turn on/off indexing on the volume")
	rootCmd.Flags().BoolP("publish", "p", false, "Publish metadata")
	rootCmd.Flags().BoolP("resolve", "t", false, "Resolve files from file id with an optional volume path")
	rootCmd.Flags().BoolP("status", "s", false, "Print indexing status of the volume")
	rootCmd.Flags().BoolP("verbose", "v", false, "Verbose mode")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"index": carapace.ActionValues("on", "off").StyleF(style.ForKeyword),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionDirectories())
}
