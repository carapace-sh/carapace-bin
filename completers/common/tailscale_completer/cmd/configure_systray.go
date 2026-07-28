package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configure_systrayCmd = &cobra.Command{
	Use:   "systray",
	Short: "[ALPHA] Manage the systray client for Linux",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configure_systrayCmd).Standalone()

	configure_systrayCmd.Flags().Bool("enable-startup", false, "install a startup script for the init system")
	configureCmd.AddCommand(configure_systrayCmd)
}
