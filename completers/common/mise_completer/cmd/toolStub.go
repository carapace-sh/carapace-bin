// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var toolStubCmd = &cobra.Command{
	Use:   "tool-stub",
	Short: "Run a tool stub file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(toolStubCmd).Standalone()

	rootCmd.AddCommand(toolStubCmd)
}
