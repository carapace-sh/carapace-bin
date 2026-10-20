package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var settings_unsetCmd = &cobra.Command{
	Use:     "unset",
	Short:   "Remove a setting from the global or project config",
	Aliases: []string{"rm", "remove", "delete", "del"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(settings_unsetCmd).Standalone()

	settings_unsetCmd.Flags().BoolP("help", "h", false, "Print help")
	settings_unsetCmd.Flags().BoolP("local", "l", false, "Write to the nearest project config instead of the global config")
	settings_unsetCmd.Flags().Bool("override", false, "Write to the global override file, `config.override.toml`")
	settingsCmd.AddCommand(settings_unsetCmd)
}
