package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "scutil",
	Short: "communicate with the configd daemon",
	Long:  "https://keith.github.io/xcode-manpages/scutil.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("allow-new-interfaces", false, "Manage new interface creation with screen locked")
	rootCmd.Flags().BoolP("dns", "d", false, "Display DNS configuration")
	rootCmd.Flags().String("error", "", "Display a descriptive message for the given error code")
	rootCmd.Flags().String("get", "", "Get the specified preference (ComputerName, LocalHostName, HostName)")
	rootCmd.Flags().BoolP("help", "H", false, "Display usage information")
	rootCmd.Flags().BoolP("nc", "n", false, "Show VPN network configuration information")
	rootCmd.Flags().Bool("nwi", false, "Show network information")
	rootCmd.Flags().BoolP("prefs", "P", false, "Interactive access to stored preferences")
	rootCmd.Flags().BoolP("proxy", "p", false, "Display proxy configuration")
	rootCmd.Flags().StringP("reach", "r", "", "Check reachability of node, address, or address pair")
	rootCmd.Flags().String("renew", "", "Re-evaluate network configuration on the interface")
	rootCmd.Flags().String("set", "", "Set the specified preference (ComputerName, LocalHostName, HostName)")
	rootCmd.Flags().IntP("timeout", "t", 0, "Time to wait for key")
	rootCmd.Flags().StringP("wait", "w", "", "Wait for presence of dynamic store key")
	rootCmd.Flags().Bool("watch", false, "Watch reachability changes")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"error": carapace.ActionValues(),
		"get":   carapace.ActionValues("ComputerName", "LocalHostName", "HostName"),
		"reach": carapace.ActionValues(),
		"set":   carapace.ActionValues("ComputerName", "LocalHostName", "HostName"),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"add", "Add a new key-value pair",
			"close", "Close a preference file",
			"get", "Get the value of a key",
			"list", "List available keys",
			"notify", "Register for notification",
			"open", "Open a preference file",
			"quit", "Quit scutil",
			"readFile", "Read a file",
			"remove", "Remove a key",
			"set", "Set the value of a key",
			"show", "Show the value of a key or preference",
			"wait", "Wait for notification",
			"writeFile", "Write to a file",
		),
	)
}
