// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).
// Completions below the flag definitions were wired by hand.

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var enCmd = &cobra.Command{
	Use:   "en",
	Short: "Start a subshell with mise's tools and environment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(enCmd).Standalone()

	enCmd.Flags().BoolP("help", "h", false, "Print help")
	enCmd.Flags().StringP("shell", "s", "", "Shell to start; defaults to $SHELL")
	rootCmd.AddCommand(enCmd)

	carapace.Gen(enCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
