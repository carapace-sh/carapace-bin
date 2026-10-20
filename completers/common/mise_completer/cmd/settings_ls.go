// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).
// Completions below the flag definitions were wired by hand.

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/mise"
	"github.com/spf13/cobra"
)

var settings_lsCmd = &cobra.Command{
	Use:     "ls",
	Short:   "List configured settings and their sources",
	Aliases: []string{"list"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(settings_lsCmd).Standalone()

	settings_lsCmd.Flags().BoolP("all", "a", false, "List all settings")
	settings_lsCmd.Flags().Bool("complete", false, "Print all settings with descriptions for shell completions")
	settings_lsCmd.Flags().BoolP("help", "h", false, "Print help")
	settings_lsCmd.Flags().BoolP("json", "J", false, "Output in JSON format")
	settings_lsCmd.Flags().Bool("json-extended", false, "Output in JSON format with sources")
	settings_lsCmd.PersistentFlags().BoolP("local", "l", false, "Use only the nearest project config")
	settings_lsCmd.Flags().BoolP("toml", "T", false, "Output in TOML format")
	settings_lsCmd.Flag("complete").Hidden = true
	settingsCmd.AddCommand(settings_lsCmd)

	carapace.Gen(settings_lsCmd).PositionalCompletion(
		mise.ActionSettings(),
	)
}
