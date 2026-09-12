package cmd

import (
	"github.com/carapace-sh/carapace"
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
	settings_lsCmd.Flags().BoolP("json", "J", false, "Output in JSON format")
	settings_lsCmd.Flags().BoolP("local", "l", false, "Use local config file")
	settings_lsCmd.Flags().BoolP("toml", "T", false, "Output in TOML format")
	settingsCmd.AddCommand(settings_lsCmd)
}
