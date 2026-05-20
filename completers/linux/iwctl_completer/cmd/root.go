package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/iwctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

type familyCommand struct {
	name        string
	description string
	entity      bool
	args        func([]string) carapace.Action
}

type family struct {
	name        string
	description string
	entity      carapace.Action
	commands    []familyCommand
}

var rootCmd = &cobra.Command{
	Use:   "iwctl",
	Short: "Internet wireless control utility",
	Long:  "https://man.archlinux.org/man/iwctl.1",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("dont-ask", "v", false, "Don't ask for missing credentials")
	rootCmd.Flags().BoolP("help", "h", false, "Show help message and exit")
	rootCmd.Flags().StringP("passphrase", "P", "", "Provide passphrase")
	rootCmd.Flags().StringP("password", "p", "", "Provide password")
	rootCmd.Flags().StringP("username", "u", "", "Provide username")

	for _, f := range families() {
		addFamily(f)
	}
}

func addFamily(f family) {
	cmd := &cobra.Command{
		Use:   f.name,
		Short: f.description,
		Run:   func(cmd *cobra.Command, args []string) {},
	}
	rootCmd.AddCommand(cmd)

	carapace.Gen(cmd).PositionalAnyCompletion(carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return completeFamily(f, c.Args)
	}))
}

func completeFamily(f family, args []string) carapace.Action {
	switch len(args) {
	case 0:
		return carapace.Batch(
			actionNonEntityCommands(f),
			f.entity,
		).ToA()
	case 1:
		if isNonEntityCommand(f, args[0]) {
			return carapace.ActionValues()
		}
		return actionEntityCommands(f)
	default:
		if cmd := findCommand(f, args[1]); cmd != nil && cmd.args != nil {
			return cmd.args(args)
		}
	}
	return carapace.ActionValues()
}

func actionNonEntityCommands(f family) carapace.Action {
	vals := make([]string, 0)
	for _, cmd := range f.commands {
		if !cmd.entity {
			vals = append(vals, cmd.name, cmd.description)
		}
	}
	return carapace.ActionValuesDescribed(vals...)
}

func actionEntityCommands(f family) carapace.Action {
	vals := make([]string, 0)
	for _, cmd := range f.commands {
		if cmd.entity {
			vals = append(vals, cmd.name, cmd.description)
		}
	}
	return carapace.ActionValuesDescribed(vals...)
}

func isNonEntityCommand(f family, name string) bool {
	for _, cmd := range f.commands {
		if !cmd.entity && cmd.name == name {
			return true
		}
	}
	return false
}

func findCommand(f family, name string) *familyCommand {
	for index, cmd := range f.commands {
		if cmd.name == name {
			return &f.commands[index]
		}
	}
	return nil
}

func families() []family {
	return []family{
		{
			name:        "adapter",
			description: "Wireless adapters",
			entity:      action.ActionAdapters(),
			commands: []familyCommand{
				{name: "list", description: "List adapters"},
				{name: "show", description: "Show adapter info", entity: true},
				{name: "set-property", description: "Set property", entity: true, args: adapterPropertyArgs},
			},
		},
		{
			name:        "device",
			description: "Wireless devices",
			entity:      action.ActionDevices(),
			commands: []familyCommand{
				{name: "list", description: "List devices"},
				{name: "show", description: "Show device info", entity: true},
				{name: "set-property", description: "Set property", entity: true, args: devicePropertyArgs},
			},
		},
		{
			name:        "station",
			description: "Station mode",
			entity:      action.ActionDevices(),
			commands: []familyCommand{
				{name: "list", description: "List devices in Station mode"},
				{name: "connect", description: "Connect to network", entity: true, args: stationConnectArgs},
				{name: "connect-hidden", description: "Connect to hidden network", entity: true},
				{name: "disconnect", description: "Disconnect", entity: true},
				{name: "get-networks", description: "Get networks", entity: true, args: signalFormatArgs},
				{name: "get-hidden-access-points", description: "Get hidden APs", entity: true, args: signalFormatArgs},
				{name: "scan", description: "Scan for networks", entity: true},
				{name: "show", description: "Show station info", entity: true},
				{name: "get-bsses", description: "Get BSS's for a network", entity: true, args: stationConnectArgs},
			},
		},
		{
			name:        "known-networks",
			description: "Known networks",
			entity:      action.ActionKnownNetworks(),
			commands: []familyCommand{
				{name: "list", description: "List known networks"},
				{name: "forget", description: "Forget known network", entity: true},
				{name: "show", description: "Show known network", entity: true},
				{name: "set-property", description: "Set property", entity: true, args: knownNetworkPropertyArgs},
			},
		},
		{
			name:        "ap",
			description: "Access point mode",
			entity:      action.ActionDevices(),
			commands: []familyCommand{
				{name: "list", description: "List devices in AP mode"},
				{name: "start", description: "Start an access point", entity: true},
				{name: "start-profile", description: "Start an access point based on a disk profile", entity: true},
				{name: "stop", description: "Stop a started access point", entity: true},
				{name: "show", description: "Show AP info", entity: true},
				{name: "scan", description: "Start an AP scan", entity: true},
				{name: "get-networks", description: "Get network list after scanning", entity: true},
			},
		},
		{
			name:        "ad-hoc",
			description: "Ad-hoc mode",
			entity:      action.ActionDevices(),
			commands: []familyCommand{
				{name: "list", description: "List devices in Ad-hoc mode"},
				{name: "start", description: "Start or join an Ad-Hoc network", entity: true},
				{name: "start_open", description: "Start or join an open Ad-Hoc network", entity: true},
				{name: "stop", description: "Leave an Ad-Hoc network", entity: true},
			},
		},
		{
			name:        "wsc",
			description: "Wi-Fi Simple Configuration",
			entity:      action.ActionDevices(),
			commands: []familyCommand{
				{name: "list", description: "List WSC-capable devices"},
				{name: "push-button", description: "PushButton mode", entity: true},
				{name: "start-user-pin", description: "PIN mode", entity: true},
				{name: "start-pin", description: "PIN mode with generated 8 digit PIN", entity: true},
				{name: "cancel", description: "Abort WSC operations", entity: true},
			},
		},
		{
			name:        "dpp",
			description: "Device Provisioning Protocol",
			entity:      action.ActionDevices(),
			commands: []familyCommand{
				{name: "list", description: "List DPP-capable devices"},
				{name: "start-enrollee", description: "Starts a DPP Enrollee", entity: true},
				{name: "start-configurator", description: "Starts a DPP Configurator", entity: true},
				{name: "stop", description: "Abort DPP operations", entity: true},
				{name: "show", description: "Shows the DPP state", entity: true},
			},
		},
		{
			name:        "pkex",
			description: "DPP PKEX shared code",
			entity:      action.ActionDevices(),
			commands: []familyCommand{
				{name: "list", description: "List shared code capable devices"},
				{name: "stop", description: "Abort shared code operations", entity: true},
				{name: "show", description: "Shows the shared code state", entity: true},
				{name: "enroll", description: "Start a shared code enrollee", entity: true},
				{name: "configure", description: "Start a shared code configurator", entity: true},
			},
		},
		{
			name:        "debug",
			description: "Station debugging",
			entity:      action.ActionDevices(),
			commands: []familyCommand{
				{name: "connect", description: "Connect to a specific BSS", entity: true},
				{name: "roam", description: "Roam to a BSS", entity: true},
				{name: "get-networks", description: "Get networks", entity: true},
				{name: "autoconnect", description: "Set AutoConnect property", entity: true, args: onOffArgs},
			},
		},
	}
}

func adapterPropertyArgs(args []string) carapace.Action {
	return propertyArgs(carapace.ActionValues("Powered"), args)
}

func devicePropertyArgs(args []string) carapace.Action {
	return propertyArgs(carapace.ActionValues("Mode", "Powered"), args)
}

func knownNetworkPropertyArgs(args []string) carapace.Action {
	return propertyArgs(carapace.ActionValues("AutoConnect"), args)
}

func propertyArgs(properties carapace.Action, args []string) carapace.Action {
	switch len(args) {
	case 2:
		return properties
	case 3:
		switch args[2] {
		case "AutoConnect", "Powered":
			return carapace.ActionValues("on", "off")
		case "Mode":
			return carapace.ActionValues("ad-hoc", "ap", "station")
		}
	}
	return carapace.ActionValues()
}

func stationConnectArgs(args []string) carapace.Action {
	switch len(args) {
	case 2:
		return action.ActionNetworks(args[0])
	case 3:
		return carapace.ActionValues("open", "psk", "8021x", "wep")
	default:
		return carapace.ActionValues()
	}
}

func signalFormatArgs(args []string) carapace.Action {
	if len(args) == 2 {
		return carapace.ActionValues("rssi-dbms", "rssi-bars")
	}
	return carapace.ActionValues()
}

func onOffArgs(args []string) carapace.Action {
	if len(args) == 2 {
		return carapace.ActionValues("on", "off")
	}
	return carapace.ActionValues()
}
