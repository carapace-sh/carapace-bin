package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ksetup",
	Short: "configure Kerberos realm settings",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/ksetup",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"setrealm", "set the Kerberos realm",
			"delrealm", "delete the Kerberos realm",
			"addkdc", "add a KDC address",
			"delkdc", "delete a KDC address",
			"addkpasswd", "add a Kerberos password server",
			"delkpasswd", "delete a Kerberos password server",
			"setencryptiontype", "set encryption type",
			"dumpstate", "display current Kerberos configuration",
			"addhost", "add a host-to-realm mapping",
			"delhost", "delete a host-to-realm mapping",
			"setcomputerpassword", "set the computer account password",
			"mapuser", "map a Kerberos principal to an account",
			"unmapuser", "remove a Kerberos principal mapping",
		),
	)
}
