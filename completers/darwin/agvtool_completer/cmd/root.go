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

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"new-version", "Set a new version number",
			"vers", "Set version number",
			"mvers", "Get marketing version",
			"version", "Get version number",
			"what-version", "Print current version",
			"next-version", "Increment version number",
		),
	)
}
