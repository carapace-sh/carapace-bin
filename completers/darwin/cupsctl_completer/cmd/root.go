package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cupsctl",
	Short: "configure cupsd.conf options",
	Long:  "https://keith.github.io/xcode-manpages/cupsctl.8.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("E", false, "Enable encryption on the connection")
	rootCmd.Flags().String("U", "", "Specify an alternate username for authentication")
	rootCmd.Flags().Bool("debug-logging", false, "Enable debug logging to the error_log file")
	rootCmd.Flags().String("h", "", "Specify the server address")
	rootCmd.Flags().Bool("no-debug-logging", false, "Disable debug logging")
	rootCmd.Flags().Bool("no-remote-admin", false, "Disable remote administration")
	rootCmd.Flags().Bool("no-remote-any", false, "Disable printing from any address")
	rootCmd.Flags().Bool("no-share-printers", false, "Disable sharing of local printers")
	rootCmd.Flags().Bool("no-user-cancel-any", false, "Prevent users from cancelling jobs owned by others")
	rootCmd.Flags().Bool("remote-admin", false, "Enable remote administration")
	rootCmd.Flags().Bool("remote-any", false, "Enable printing from any address")
	rootCmd.Flags().Bool("share-printers", false, "Enable sharing of local printers")
	rootCmd.Flags().Bool("user-cancel-any", false, "Allow users to cancel jobs owned by others")
}
