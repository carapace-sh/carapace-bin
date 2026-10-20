// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).
// Completions below the flag definitions were wired by hand.

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/mise"
	"github.com/spf13/cobra"
)

var toolCmd = &cobra.Command{
	Use:   "tool",
	Short: "Show information about a tool",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(toolCmd).Standalone()

	toolCmd.Flags().Bool("active", false, "Only show active versions")
	toolCmd.Flags().Bool("backend", false, "Only show backend field")
	toolCmd.Flags().Bool("config-source", false, "Only show config source")
	toolCmd.Flags().Bool("description", false, "Only show description field")
	toolCmd.Flags().BoolP("help", "h", false, "Print help")
	toolCmd.Flags().Bool("installed", false, "Only show installed versions")
	toolCmd.Flags().BoolP("json", "J", false, "Output in JSON format")
	toolCmd.Flags().Bool("requested", false, "Only show requested versions")
	toolCmd.Flags().Bool("tool-options", false, "Only show tool options")
	toolCmd.Flags().Bool("url", false, "Only show the project URL from the registry")
	rootCmd.AddCommand(toolCmd)

	carapace.Gen(toolCmd).PositionalCompletion(
		mise.ActionTools(),
	)
}
