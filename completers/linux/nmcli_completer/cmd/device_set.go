package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/nmcli_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var device_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Modify device properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(device_setCmd).Standalone()

	device_setCmd.Flags().Bool("permanent", false, "persist managed state to disk")
	device_setCmd.Flags().Bool("permanent-only", false, "set only permanent managed state")
	deviceCmd.AddCommand(device_setCmd)

	carapace.Gen(device_setCmd).PositionalCompletion(
		net.ActionDevices(net.AllDevices),
		carapace.ActionValues("autoconnect", "managed"),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 0 && c.Args[len(c.Args)-1] == "autoconnect" {
				return action.ActionYesNo()
			}
			return carapace.ActionValues("yes", "no", "up", "down", "reset")
		}),
		carapace.ActionValues("autoconnect", "managed"),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 0 && c.Args[len(c.Args)-1] == "autoconnect" {
				return action.ActionYesNo()
			}
			return carapace.ActionValues("yes", "no", "up", "down", "reset")
		}),
	)
}
