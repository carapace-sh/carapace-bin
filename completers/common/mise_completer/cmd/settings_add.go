// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).
// Completions below the flag definitions were wired by hand.

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/mise"
	"github.com/spf13/cobra"
)

var settings_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Append a value to an array setting",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(settings_addCmd).Standalone()

	settings_addCmd.Flags().BoolP("help", "h", false, "Print help")
	settings_addCmd.Flags().BoolP("local", "l", false, "Write to the nearest project config instead of the global config")
	settings_addCmd.Flags().Bool("override", false, "Write to the global override file, `config.override.toml`")
	settingsCmd.AddCommand(settings_addCmd)

	carapace.Gen(settings_addCmd).PositionalCompletion(
		mise.ActionSettings(),
	)
}
