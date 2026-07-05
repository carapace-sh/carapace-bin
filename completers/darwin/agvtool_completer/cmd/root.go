package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "agvtool",
	Short: "Apple Generic Versioning tool",
	Long:  "https://keith.github.io/xcode-manpages/agvtool.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("help", "h", false, "Display usage information")
	rootCmd.Flags().Bool("noscm", false, "Do not use source control")
	rootCmd.Flags().Bool("usecvs", false, "Use CVS for source control")
	rootCmd.Flags().Bool("usesvn", false, "Use Subversion for source control")
}
