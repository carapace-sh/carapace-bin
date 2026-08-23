package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "smbutil",
	Short: "SMB/CIFS utility",
	Long:  "https://keith.github.io/xcode-manpages/smbutil.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("help", "h", false, "Print a short help message")
	rootCmd.Flags().BoolP("verbose", "v", false, "Verbose output")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"help", "Display help on specified subcommand",
			"lookup", "Resolve NetBIOS name to IP address",
			"lookup_ns", "Resolve NetBIOS name via NBNS",
			"status", "List the current SMB connections",
			"view", "List the available SMB resources on a server",
			"login", "Log into a SMB server",
			"logout", "Log out of a SMB server",
			"identity", "Display the current SMB identity",
			"statshares", "Display SMB share statistics",
			"workgroup", "Display the current SMB workgroup",
			"lanmanager", "Show LAN Manager compatibility information",
		),
	)
}