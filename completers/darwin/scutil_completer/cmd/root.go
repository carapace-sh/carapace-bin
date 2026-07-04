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

	rootCmd.Flags().BoolP("dns", "d", false, "Display DNS configuration")
	rootCmd.Flags().BoolP("help", "H", false, "Display usage information")
	rootCmd.Flags().BoolP("host", "h", false, "Display host information")
	rootCmd.Flags().BoolP("local", "l", false, "Use the local preferences")
	rootCmd.Flags().BoolP("net", "n", false, "Display network configuration")
	rootCmd.Flags().BoolP("nw", "w", false, "Display network reachability")
	rootCmd.Flags().BoolP("prefs", "P", false, "Display preferences")
	rootCmd.Flags().BoolP("proxy", "p", false, "Display proxy configuration")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"get", "Get the value of a key",
			"set", "Set the value of a key",
			"show", "Show the value of a key or preference",
			"add", "Add a new key-value pair",
			"remove", "Remove a key",
			"list", "List available keys",
			"open", "Open a preference file",
			"close", "Close a preference file",
			"quit", "Quit scutil",
			"readFile", "Read a file",
			"writeFile", "Write to a file",
			"notify", "Register for notification",
			"wait", "Wait for notification",
		),
	)
}
