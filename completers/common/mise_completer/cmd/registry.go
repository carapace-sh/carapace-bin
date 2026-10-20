// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).
// Completions below the flag definitions were wired by hand.

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/mise"
	"github.com/spf13/cobra"
)

var registryCmd = &cobra.Command{
	Use:   "registry",
	Short: "List registry shorthand names and their backends",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(registryCmd).Standalone()

	registryCmd.Flags().StringP("backend", "b", "", "Show only tools for this backend")
	registryCmd.Flags().Bool("complete", false, "Print all tools with descriptions for shell completions")
	registryCmd.Flags().BoolP("help", "h", false, "Print help")
	registryCmd.Flags().Bool("hide-aliased", false, "Hide names that are aliases of another registry tool")
	registryCmd.Flags().BoolP("json", "J", false, "Output in JSON format")
	registryCmd.Flags().Bool("security", false, "Include security features for each tool's backends in JSON output")
	registryCmd.Flag("complete").Hidden = true
	rootCmd.AddCommand(registryCmd)

	carapace.Gen(registryCmd).FlagCompletion(carapace.ActionMap{
		"backend": mise.ActionBackends(),
	})
}
