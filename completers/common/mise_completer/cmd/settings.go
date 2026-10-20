// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).
// Completions below the flag definitions were wired by hand.

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/mise"
	"github.com/spf13/cobra"
)

var settingsCmd = &cobra.Command{
	Use:   "settings",
	Short: "Manage settings",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(settingsCmd).Standalone()

	settingsCmd.Flags().BoolP("all", "a", false, "List all settings")
	settingsCmd.Flags().Bool("complete", false, "Print all settings with descriptions for shell completions")
	settingsCmd.Flags().BoolP("help", "h", false, "Print help")
	settingsCmd.Flags().BoolP("json", "J", false, "Output in JSON format")
	settingsCmd.Flags().Bool("json-extended", false, "Output in JSON format with sources")
	settingsCmd.PersistentFlags().BoolP("local", "l", false, "Use only the nearest project config")
	settingsCmd.Flags().BoolP("toml", "T", false, "Output in TOML format")
	settingsCmd.Flag("complete").Hidden = true
	rootCmd.AddCommand(settingsCmd)

	carapace.Gen(settingsCmd).PositionalCompletion(
		mise.ActionSettings(),
	)
}
