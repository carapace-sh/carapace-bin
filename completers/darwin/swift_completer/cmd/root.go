package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "swift",
	Short: "Swift REPL and compiler",
	Long:  "https://keith.github.io/xcode-manpages/swift.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.Flags().BoolS("Help|-h", "Help|-h", false, "")
	rootCmd.Flags().BoolS("Version|-version", "Version|-version", false, "")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
