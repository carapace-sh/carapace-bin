package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/iwctl"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "iwctl [flags] [commands]",
	Short: "Internet wireless control utility",
	Long:  "https://archive.kernel.org/oldwiki/iwd.wiki.kernel.org/",
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

	rootCmd.AddCommand(adapterCmd)
	rootCmd.AddCommand(adHocCmd)
	rootCmd.AddCommand(apCmd)
	rootCmd.AddCommand(debugCmd)
	rootCmd.AddCommand(deviceCmd)
	rootCmd.AddCommand(dppCmd)
	rootCmd.AddCommand(exitCmd)
	rootCmd.AddCommand(helpCmd)
	rootCmd.AddCommand(knownNetworksCmd)
	rootCmd.AddCommand(pkexCmd)
	rootCmd.AddCommand(quitCmd)
	rootCmd.AddCommand(stationCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(wscCmd)

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"passphrase": carapace.ActionValues(),
		"password":   carapace.ActionValues(),
		"username":   carapace.ActionValues(),
	})
}

var adapterCmd = &cobra.Command{
	Use:   "adapter",
	Short: "manage adapters",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var adHocCmd = &cobra.Command{
	Use:   "ad-hoc",
	Short: "manage ad-hoc networks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var apCmd = &cobra.Command{
	Use:   "ap",
	Short: "manage access points",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var debugCmd = &cobra.Command{
	Use:   "debug",
	Short: "manage station debug functions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var deviceCmd = &cobra.Command{
	Use:   "device",
	Short: "manage devices",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var dppCmd = &cobra.Command{
	Use:   "dpp",
	Short: "manage device provisioning",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var exitCmd = &cobra.Command{
	Use:   "exit",
	Short: "exit interactive mode",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "display help",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var knownNetworksCmd = &cobra.Command{
	Use:   "known-networks",
	Short: "manage known networks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var pkexCmd = &cobra.Command{
	Use:   "pkex",
	Short: "manage shared code device provisioning",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var quitCmd = &cobra.Command{
	Use:   "quit",
	Short: "quit interactive mode",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var stationCmd = &cobra.Command{
	Use:   "station",
	Short: "manage stations",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "display version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var wscCmd = &cobra.Command{
	Use:   "wsc",
	Short: "manage Wi-Fi Simple Configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(adapterCmd).Standalone()
	carapace.Gen(adapterCmd).PositionalCompletion(
		withList(iwctl.ActionAdapters(), "list adapters"),
		afterEntity(commands(
			"show", "show adapter info",
			"set-property", "set property",
		)),
		propertyAfter("set-property", iwctl.ActionAdapterProperties()),
		propertyValueAfter("set-property", 2),
	)

	carapace.Gen(deviceCmd).Standalone()
	carapace.Gen(deviceCmd).PositionalCompletion(
		withList(iwctl.ActionDevices(), "list devices"),
		afterEntity(commands(
			"show", "show device info",
			"set-property", "set property",
		)),
		propertyAfter("set-property", iwctl.ActionDeviceProperties()),
		propertyValueAfter("set-property", 2),
	)

	carapace.Gen(knownNetworksCmd).Standalone()
	carapace.Gen(knownNetworksCmd).PositionalCompletion(
		withList(iwctl.ActionKnownNetworks(), "list known networks"),
		afterEntity(commands(
			"forget", "forget known network",
			"show", "show known network",
			"set-property", "set property",
		)),
		propertyAfter("set-property", iwctl.ActionKnownNetworkProperties()),
		propertyValueAfter("set-property", 2),
	)

	carapace.Gen(stationCmd).Standalone()
	carapace.Gen(stationCmd).PositionalCompletion(
		withList(iwctl.ActionDevices(), "list devices in station mode"),
		afterEntity(commands(
			"connect", "connect to network",
			"connect-hidden", "connect to hidden network",
			"disconnect", "disconnect",
			"get-bsses", "get BSSes for a network",
			"get-hidden-access-points", "get hidden access points",
			"get-networks", "get networks",
			"scan", "scan for networks",
			"show", "show station info",
		)),
		stationArgument(),
		stationSecurityArgument(),
	)

	carapace.Gen(apCmd).Standalone()
	carapace.Gen(apCmd).PositionalCompletion(
		withList(iwctl.ActionDevices(), "list devices in AP mode"),
		afterEntity(commands(
			"get-networks", "get network list after scanning",
			"scan", "start an AP scan",
			"show", "show AP info",
			"start", "start an access point",
			"start-profile", "start an access point from a profile",
			"stop", "stop a started access point",
		)),
	)

	carapace.Gen(adHocCmd).Standalone()
	carapace.Gen(adHocCmd).PositionalCompletion(
		withList(iwctl.ActionDevices(), "list devices in ad-hoc mode"),
		afterEntity(commands(
			"start", "start or join an ad-hoc network",
			"start_open", "start or join an open ad-hoc network",
			"stop", "leave an ad-hoc network",
		)),
	)

	carapace.Gen(dppCmd).Standalone()
	carapace.Gen(dppCmd).PositionalCompletion(
		withList(iwctl.ActionDevices(), "list DPP-capable devices"),
		afterEntity(commands(
			"show", "show DPP state",
			"start-configurator", "start a DPP configurator",
			"start-enrollee", "start a DPP enrollee",
			"stop", "abort DPP operations",
		)),
	)

	carapace.Gen(pkexCmd).Standalone()
	carapace.Gen(pkexCmd).PositionalCompletion(
		withList(iwctl.ActionDevices(), "list shared code capable devices"),
		afterEntity(commands(
			"configure", "start a shared code configurator",
			"enroll", "start a shared code enrollee",
			"show", "show shared code state",
			"stop", "abort shared code operations",
		)),
	)

	carapace.Gen(wscCmd).Standalone()
	carapace.Gen(wscCmd).PositionalCompletion(
		withList(iwctl.ActionDevices(), "list WSC-capable devices"),
		afterEntity(commands(
			"cancel", "abort WSC operations",
			"push-button", "push button mode",
			"start-pin", "PIN mode with generated PIN",
			"start-user-pin", "PIN mode",
		)),
	)

	carapace.Gen(debugCmd).Standalone()
	carapace.Gen(debugCmd).PositionalCompletion(
		iwctl.ActionDevices(),
		commands(
			"autoconnect", "set AutoConnect property",
			"connect", "connect to a specific BSS",
			"get-networks", "get networks",
			"roam", "roam to a BSS",
		),
		debugArgument(),
	)
}

func commands(values ...string) carapace.Action {
	return carapace.ActionValuesDescribed(values...)
}

func withList(action carapace.Action, description string) carapace.Action {
	return carapace.Batch(
		carapace.ActionValuesDescribed("list", description),
		action,
	).ToA()
}

func afterEntity(action carapace.Action) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if len(c.Args) > 0 && c.Args[0] == "list" {
			return carapace.ActionValues()
		}
		return action
	})
}

func propertyAfter(command string, action carapace.Action) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if len(c.Args) > 1 && c.Args[1] == command {
			return action
		}
		return carapace.ActionValues()
	})
}

func propertyValueAfter(command string, propertyIndex int) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if len(c.Args) > propertyIndex && c.Args[1] == command {
			return iwctl.ActionPropertyValues(c.Args[propertyIndex])
		}
		return carapace.ActionValues()
	})
}

func stationArgument() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if len(c.Args) < 2 {
			return carapace.ActionValues()
		}

		switch c.Args[1] {
		case "connect", "get-bsses":
			return iwctl.ActionNetworks(c.Args[0])
		case "get-hidden-access-points":
			return carapace.ActionValues("rssi-dbms")
		case "get-networks":
			return carapace.ActionValues("rssi-dbms", "rssi-bars")
		default:
			return carapace.ActionValues()
		}
	})
}

func stationSecurityArgument() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if len(c.Args) < 2 {
			return carapace.ActionValues()
		}

		switch c.Args[1] {
		case "connect", "get-bsses":
			return iwctl.ActionSecurityTypes()
		default:
			return carapace.ActionValues()
		}
	})
}

func debugArgument() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if len(c.Args) < 2 {
			return carapace.ActionValues()
		}

		switch c.Args[1] {
		case "autoconnect":
			return carapace.ActionValues("on", "off")
		case "connect", "roam":
			return iwctl.ActionBSSes(c.Args[0])
		default:
			return carapace.ActionValues()
		}
	})
}
