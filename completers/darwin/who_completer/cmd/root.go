package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "who",
	Short: "display who is logged in",
	Long:  "https://keith.github.io/xcode-manpages/who.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("H", "H", false, "Print column headings")
	rootCmd.Flags().BoolS("T", "T", false, "Show message state")
	rootCmd.Flags().BoolS("a", "a", false, "Show all")
	rootCmd.Flags().BoolS("b", "b", false, "Show last reboot")
	rootCmd.Flags().BoolS("m", "m", false, "Show hostname and user associated with stdin")
	rootCmd.Flags().BoolS("q", "q", false, "Count all login sessions")
	rootCmd.Flags().BoolS("s", "s", false, "Print name, line, time (default)")
	rootCmd.Flags().BoolS("u", "u", false, "Show idle time")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
