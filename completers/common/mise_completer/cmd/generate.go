// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:     "generate",
	Short:   "Generate project files such as CI workflows, git hooks, and scripts",
	Aliases: []string{"gen", "g"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(generateCmd).Standalone()

	generateCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(generateCmd)
}
