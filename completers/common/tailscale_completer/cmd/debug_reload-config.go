package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_reload_configCmd = &cobra.Command{
	Use:   "reload-config",
	Short: "Reload config",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_reload_configCmd).Standalone()

	debugCmd.AddCommand(debug_reload_configCmd)
}
