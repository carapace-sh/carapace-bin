package cmd

import (
	"sort"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "iwctl",
	Short: "Internet wireless control utility",
	Long:  "https://iwd.wiki.kernel.org/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("username", "u", "", "Provide username")
	rootCmd.Flags().StringP("password", "p", "", "Provide password")
	rootCmd.Flags().StringP("passphrase", "P", "", "Provide passphrase")
	rootCmd.Flags().BoolP("dont-ask", "v", false, "Do not ask for missing credentials")
	rootCmd.Flags().BoolP("help", "h", false, "Show help message and exit")

	addCommands()
}

func addCommands() {
	adapterCmd := addCommand(rootCmd, "adapter [adapter] [command]", "Manage adapters")
	carapace.Gen(adapterCmd).PositionalAnyCompletion(actionAdapterArgs())

	deviceCmd := addCommand(rootCmd, "device [device] [command]", "Manage devices")
	carapace.Gen(deviceCmd).PositionalAnyCompletion(actionDeviceArgs())

	stationCmd := addCommand(rootCmd, "station DEVICE COMMAND", "Manage Wi-Fi stations")
	carapace.Gen(stationCmd).PositionalAnyCompletion(actionStationArgs())

	knownNetworksCmd := addCommand(rootCmd, "known-networks [command]", "Manage known networks")
	carapace.Gen(knownNetworksCmd).PositionalAnyCompletion(actionKnownNetworkArgs())

	apCmd := addCommand(rootCmd, "ap DEVICE COMMAND", "Manage access point mode")
	carapace.Gen(apCmd).PositionalAnyCompletion(actionAPArgs())

	wscCmd := addCommand(rootCmd, "wsc DEVICE COMMAND", "Manage Wi-Fi Simple Configuration")
	carapace.Gen(wscCmd).PositionalAnyCompletion(actionWSCArgs())

	helpCmd := addCommand(rootCmd, "help [command]", "Show help")
	carapace.Gen(helpCmd).PositionalCompletion(actionTopLevelCommands())

	addCommand(rootCmd, "version", "Show version")
	addCommand(rootCmd, "quit", "Exit interactive mode")
	addCommand(rootCmd, "exit", "Exit interactive mode")
}

func addCommand(parent *cobra.Command, use string, short string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   use,
		Short: short,
		Run:   func(cmd *cobra.Command, args []string) {},
	}
	carapace.Gen(cmd).Standalone()
	parent.AddCommand(cmd)
	return cmd
}

func actionTopLevelCommands() carapace.Action {
	return carapace.ActionValuesDescribed(
		"adapter", "Manage adapters",
		"device", "Manage devices",
		"station", "Manage Wi-Fi stations",
		"known-networks", "Manage known networks",
		"ap", "Manage access point mode",
		"wsc", "Manage Wi-Fi Simple Configuration",
		"help", "Show help",
		"version", "Show version",
		"quit", "Exit interactive mode",
		"exit", "Exit interactive mode",
	)
}

func actionAdapterArgs() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return carapace.Batch(
				carapace.ActionValuesDescribed("list", "List adapters"),
				actionAdapters(),
			).ToA()
		case 1:
			if c.Parts[0] == "list" {
				return carapace.ActionValues()
			}
			return carapace.ActionValuesDescribed(
				"show", "Show adapter details",
				"set-property", "Set adapter property",
			)
		case 2:
			if c.Parts[1] == "set-property" {
				return carapace.ActionValues("Powered")
			}
		case 3:
			if c.Parts[1] == "set-property" && c.Parts[2] == "Powered" {
				return actionOnOff()
			}
		}
		return carapace.ActionValues()
	})
}

func actionDeviceArgs() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return carapace.Batch(
				carapace.ActionValuesDescribed("list", "List devices"),
				actionDevices(),
			).ToA()
		case 1:
			if c.Parts[0] == "list" {
				return carapace.ActionValues()
			}
			return carapace.ActionValuesDescribed(
				"show", "Show device details",
				"set-property", "Set device property",
			)
		case 2:
			if c.Parts[1] == "set-property" {
				return carapace.ActionValues("Powered", "Mode")
			}
		case 3:
			if c.Parts[1] == "set-property" {
				switch c.Parts[2] {
				case "Powered":
					return actionOnOff()
				case "Mode":
					return carapace.ActionValues("station", "ap", "ad-hoc")
				}
			}
		}
		return carapace.ActionValues()
	})
}

func actionStationArgs() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return actionDevices()
		case 1:
			return carapace.ActionValuesDescribed(
				"scan", "Scan for networks",
				"get-networks", "List visible networks",
				"get-hidden-access-points", "List hidden access points",
				"connect", "Connect to a network",
				"connect-hidden", "Connect to a hidden network",
				"disconnect", "Disconnect",
				"show", "Show station details",
				"set-property", "Set station property",
				"autoconnect", "Enable or disable autoconnect",
			)
		case 2:
			switch c.Parts[1] {
			case "connect":
				return actionNetworks(c.Parts[0])
			case "connect-hidden":
				return carapace.ActionValues()
			case "set-property":
				return carapace.ActionValues("Powered", "AutoConnect", "Scanning")
			case "autoconnect":
				return actionOnOff()
			}
		case 3:
			if c.Parts[1] == "set-property" {
				switch c.Parts[2] {
				case "Powered", "AutoConnect", "Scanning":
					return actionOnOff()
				}
			}
		}
		return carapace.ActionValues()
	})
}

func actionKnownNetworkArgs() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return carapace.ActionValuesDescribed(
				"list", "List known networks",
				"show", "Show known network details",
				"forget", "Forget a known network",
			)
		case 1:
			switch c.Parts[0] {
			case "show", "forget":
				return actionKnownNetworks()
			}
		}
		return carapace.ActionValues()
	})
}

func actionAPArgs() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return actionDevices()
		case 1:
			return carapace.ActionValuesDescribed(
				"start-profile", "Start access point profile",
				"stop", "Stop access point mode",
				"show", "Show access point details",
			)
		case 2:
			if c.Parts[1] == "start-profile" {
				return actionKnownNetworks()
			}
		}
		return carapace.ActionValues()
	})
}

func actionWSCArgs() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return actionDevices()
		case 1:
			return carapace.ActionValuesDescribed(
				"push-button", "Start push-button WPS",
				"pin", "Start PIN WPS",
			)
		}
		return carapace.ActionValues()
	})
}

func actionAdapters() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommandE("iwctl", "adapter", "list")(func(output []byte, err error) carapace.Action {
			if err != nil {
				return carapace.ActionValues()
			}
			return carapace.ActionValuesDescribed(describedFirstColumn(output, "adapter", "name")...)
		})
	})
}

func actionDevices() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommandE("iwctl", "device", "list")(func(output []byte, err error) carapace.Action {
			if err != nil {
				return carapace.ActionValues()
			}
			return carapace.ActionValuesDescribed(describedFirstColumn(output, "device", "name")...)
		})
	})
}

func actionNetworks(device string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommandE("iwctl", "station", device, "get-networks")(func(output []byte, err error) carapace.Action {
			if err != nil {
				return carapace.ActionValues()
			}
			return carapace.ActionValuesDescribed(describedFirstColumn(output, "network", "ssid", "available")...)
		})
	})
}

func actionKnownNetworks() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommandE("iwctl", "known-networks", "list")(func(output []byte, err error) carapace.Action {
			if err != nil {
				return carapace.ActionValues()
			}
			return carapace.ActionValuesDescribed(describedFirstColumn(output, "network", "ssid", "known")...)
		})
	})
}

func actionOnOff() carapace.Action {
	return carapace.ActionValues("on", "off")
}

func describedFirstColumn(output []byte, skip ...string) []string {
	skips := make(map[string]bool)
	for _, value := range skip {
		skips[strings.ToLower(value)] = true
	}

	seen := make(map[string]bool)
	vals := make([]string, 0)
	for _, line := range strings.Split(string(output), "\n") {
		line = strings.TrimSpace(strings.TrimSuffix(line, "\r"))
		if line == "" || strings.HasPrefix(line, "-") || strings.HasPrefix(line, "=") {
			continue
		}
		fields := strings.Fields(line)
		if len(fields) == 0 {
			continue
		}
		value := fields[0]
		if skips[strings.ToLower(value)] || seen[value] {
			continue
		}
		seen[value] = true
		vals = append(vals, value, strings.TrimSpace(strings.TrimPrefix(line, value)))
	}
	sortPairs(vals)
	return vals
}

func sortPairs(vals []string) {
	pairs := make([][2]string, 0, len(vals)/2)
	for i := 0; i+1 < len(vals); i += 2 {
		pairs = append(pairs, [2]string{vals[i], vals[i+1]})
	}
	sort.SliceStable(pairs, func(i int, j int) bool {
		return pairs[i][0] < pairs[j][0]
	})
	for i, pair := range pairs {
		vals[i*2] = pair[0]
		vals[i*2+1] = pair[1]
	}
}
