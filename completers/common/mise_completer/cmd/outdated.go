// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).
// Completions below the flag definitions were wired by hand.

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/mise"
	"github.com/spf13/cobra"
)

var outdatedCmd = &cobra.Command{
	Use:   "outdated",
	Short: "Show outdated tool versions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(outdatedCmd).Standalone()

	outdatedCmd.Flags().BoolP("bump", "b", false, "Compare against the newest release overall, not only the configured range")
	outdatedCmd.Flags().Bool("global", false, "Only check tools defined in the global config")
	outdatedCmd.Flags().BoolP("help", "h", false, "Print help")
	outdatedCmd.Flags().Bool("inactive", false, "Also check installed tools that the current config does not request")
	outdatedCmd.Flags().BoolP("json", "J", false, "Output in JSON format")
	outdatedCmd.Flags().BoolS("l", "l", false, "Deprecated shorthand for --bump")
	outdatedCmd.Flags().Bool("local", false, "Only check tools defined in project config files")
	outdatedCmd.Flags().Bool("monorepo", false, "Placeholder for future monorepo outdated checks; `mise outdated --monorepo` is not implemented yet.")
	outdatedCmd.Flags().Bool("no-header", false, "Do not print the table header")
	outdatedCmd.Flag("l").Hidden = true
	outdatedCmd.Flag("monorepo").Hidden = true
	rootCmd.AddCommand(outdatedCmd)

	carapace.Gen(outdatedCmd).PositionalAnyCompletion(
		mise.ActionToolVersions(),
	)
}
