package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/iwctl_completer/cmd/actions"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "iwctl",
	Short: "Interactive wireless daemon client",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error { return rootCmd.Execute() }

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("help", "h", false, "show help")
	rootCmd.Flags().BoolP("version", "v", false, "show version")
	rootCmd.Flags().BoolP("dont-ask", "P", false, "do not ask for missing passwords")
	rootCmd.Flags().BoolP("passphrase", "p", false, "passphrase for network operations")

	rootCmd.AddCommand(
		adapterCmd(),
		deviceCmd(),
		stationCmd(),
		knownNetworksCmd(),
		apCmd(),
	)
}

func adapterCmd() *cobra.Command {
	cmd := command("adapter", "Manage wireless adapters")
	cmd.AddCommand(
		commandWithCompletion("list", "List adapters", nil),
		commandWithCompletion("show", "Show adapter details", actions.ActionAdapters),
	)
	return cmd
}

func deviceCmd() *cobra.Command {
	cmd := command("device", "Manage wireless devices")
	cmd.AddCommand(
		commandWithCompletion("list", "List devices", nil),
		commandWithCompletion("show", "Show device details", actions.ActionDevices),
		commandWithCompletion("set-property", "Set device property", actions.ActionDevices),
	)
	return cmd
}

func stationCmd() *cobra.Command {
	cmd := command("station", "Manage Wi-Fi stations")
	cmd.AddCommand(
		commandWithCompletion("list", "List stations", nil),
		commandWithCompletion("show", "Show station details", actions.ActionStations),
		commandWithCompletion("scan", "Scan for networks", actions.ActionStations),
		commandWithCompletion("get-networks", "List visible networks", actions.ActionStations),
		commandWithCompletion("disconnect", "Disconnect station", actions.ActionStations),
		connectCmd(),
	)
	return cmd
}

func connectCmd() *cobra.Command {
	cmd := command("connect", "Connect station to a network")
	carapace.Gen(cmd).PositionalCompletion(
		actions.ActionStations(),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) == 0 {
				return carapace.ActionValues()
			}
			return actions.ActionNetworks(c.Args[0])
		}),
	)
	return cmd
}

func knownNetworksCmd() *cobra.Command {
	cmd := command("known-networks", "Manage known networks")
	cmd.AddCommand(
		commandWithCompletion("list", "List known networks", nil),
		commandWithCompletion("forget", "Forget a known network", actions.ActionKnownNetworks),
	)
	return cmd
}

func apCmd() *cobra.Command {
	cmd := command("ap", "Manage access point mode")
	cmd.AddCommand(
		commandWithCompletion("start", "Start access point mode", actions.ActionDevices),
		commandWithCompletion("stop", "Stop access point mode", actions.ActionDevices),
	)
	return cmd
}

func command(use, short string) *cobra.Command {
	cmd := &cobra.Command{Use: use, Short: short, Run: func(cmd *cobra.Command, args []string) {}}
	carapace.Gen(cmd).Standalone()
	return cmd
}

func commandWithCompletion(use, short string, completion func() carapace.Action) *cobra.Command {
	cmd := command(use, short)
	if completion != nil {
		carapace.Gen(cmd).PositionalAnyCompletion(completion().FilterArgs())
	}
	return cmd
}
