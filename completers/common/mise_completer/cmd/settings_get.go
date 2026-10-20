// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).
// Completions below the flag definitions were wired by hand.

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/mise"
	"github.com/spf13/cobra"
)

var settings_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Show the effective value of a setting",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(settings_getCmd).Standalone()

	settings_getCmd.Flags().BoolP("help", "h", false, "Print help")
	settings_getCmd.Flags().BoolP("local", "l", false, "Read only the explicit settings in the nearest project config")
	settingsCmd.AddCommand(settings_getCmd)

	carapace.Gen(settings_getCmd).PositionalCompletion(
		mise.ActionSettings(),
	)
}
