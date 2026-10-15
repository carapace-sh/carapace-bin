package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/tlp_completer/cmd/action"
	"github.com/spf13/cobra"
)

var usbCmd = &cobra.Command{
	Use:   "usb",
	Short: "Enable autosuspend for all USB devices",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(usbCmd).Standalone()

	rootCmd.AddCommand(usbCmd)

	carapace.Gen(usbCmd).DashAnyCompletion(action.ActionConfigParams())
}
