// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).
// Completions below the flag definitions were wired by hand.

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/mise"
	"github.com/spf13/cobra"
)

var settings_setCmd = &cobra.Command{
	Use:     "set",
	Short:   "Set a setting",
	Aliases: []string{"create"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(settings_setCmd).Standalone()

	settings_setCmd.Flags().BoolP("help", "h", false, "Print help")
	settings_setCmd.Flags().BoolP("local", "l", false, "Write to the nearest project config instead of the global config")
	settings_setCmd.Flags().Bool("override", false, "Write to the global override file, `config.override.toml`")
	settingsCmd.AddCommand(settings_setCmd)

	carapace.Gen(settings_setCmd).PositionalCompletion(
		mise.ActionSettings(),
	)
}
