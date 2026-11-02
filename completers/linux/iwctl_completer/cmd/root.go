package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/iwctl_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "iwctl",
	Short: "Internet Wireless Daemon control utility",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

var families = []string{"ad-hoc", "adapter", "ap", "debug", "device", "dpp", "known-networks", "pkex", "station", "wsc"}

var miscCommands = []action.Command{
	{Name: "exit", Description: "Quit program"},
	{Name: "help", Description: "Display help"},
	{Name: "quit", Description: "Quit program"},
	{Name: "version", Description: "Display version"},
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("dont-ask", "d", false, "Don't ask for missing credentials")
	rootCmd.Flags().BoolP("help", "h", false, "Display help")
	rootCmd.Flags().StringP("passphrase", "P", "", "Provide passphrase")
	rootCmd.Flags().StringP("password", "p", "", "Provide password")
	rootCmd.Flags().StringP("username", "u", "", "Provide username")

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return actionCommands(c.Args)
		}),
	)
}

func actionCommands(args []string) carapace.Action {
	if len(args) == 0 {
		vals := make([]string, 0)
		for _, family := range families {
			vals = append(vals, family, action.FamilyDescriptions[family])
		}
		for _, command := range miscCommands {
			vals = append(vals, command.Name, command.Description)
		}
		return carapace.ActionValuesDescribed(vals...)
	}

	commands, ok := action.Commands[args[0]]
	if !ok {
		return carapace.ActionValues()
	}

	entityless := make([]string, 0)
	withEntity := make([]string, 0)
	for _, command := range commands {
		if command.Entity {
			withEntity = append(withEntity, command.Name)
		} else {
			entityless = append(entityless, command.Name)
		}
	}

	switch len(args) {
	case 1:
		entity := entityAction(args[0])
		if len(entityless) == 0 {
			return entity
		}
		return carapace.Batch(
			entity,
			carapace.ActionValues(entityless...),
		).ToA()
	case 2:
		if isEntityless(commands, args[1]) {
			return carapace.ActionValues()
		}
		return carapace.ActionValues(withEntity...)
	default:
		if isEntityless(commands, args[1]) {
			return carapace.ActionValues()
		}
		return actionArguments(args[0], args[2], args[3:])
	}
}

func isEntityless(commands []action.Command, token string) bool {
	for _, command := range commands {
		if command.Name == token && !command.Entity {
			return true
		}
	}
	return false
}

func entityAction(family string) carapace.Action {
	switch family {
	case "adapter":
		return action.ActionAdapters()
	case "known-networks":
		// known networks are only available via D-Bus
		return carapace.ActionValues()
	default:
		return action.ActionDevices()
	}
}

func actionArguments(family, command string, args []string) carapace.Action {
	index := len(args)

	switch family {
	case "adapter", "device":
		if command == "set-property" {
			switch index {
			case 0:
				if family == "adapter" {
					return carapace.ActionValues("Powered")
				}
				return carapace.ActionValues("Mode", "Powered")
			case 1:
				return propertyValues(family, args[0])
			}
		}
	case "known-networks":
		if command == "set-property" {
			switch index {
			case 0:
				return carapace.ActionValues("AutoConnect")
			case 1:
				return propertyValues(family, args[0])
			}
		}
	case "debug":
		switch command {
		case "autoconnect":
			if index == 0 {
				return carapace.ActionValues("off", "on")
			}
		case "connect", "roam":
			if index == 0 {
				return net.ActionBssids()
			}
		}
	case "station":
		switch command {
		case "connect", "get-bsses":
			switch index {
			case 0:
				return net.ActionSsids()
			case 1:
				return carapace.ActionValues("8021x", "open", "psk")
			}
		case "get-networks":
			if index == 0 {
				return carapace.ActionValues("rssi-bars", "rssi-dbms")
			}
		case "get-hidden-access-points":
			if index == 0 {
				return carapace.ActionValues("rssi-dbms")
			}
		}
	}

	return carapace.ActionValues()
}

func propertyValues(family, property string) carapace.Action {
	switch family {
	case "known-networks":
		if property == "AutoConnect" {
			return carapace.ActionValues("no", "yes")
		}
	case "device":
		switch property {
		case "Mode":
			return carapace.ActionValues("ad-hoc", "ap", "station")
		case "Powered":
			return carapace.ActionValues("off", "on")
		}
	case "adapter":
		if property == "Powered" {
			return carapace.ActionValues("off", "on")
		}
	}
	return carapace.ActionValues()
}
