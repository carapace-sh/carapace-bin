package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "whoami",
	Short: "display user, group, and privileges information",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/whoami",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"upn", "display user principal name",
			"fqdn", "display fully qualified domain name",
			"logonid", "display logon ID",
			"user", "display current user information",
			"groups", "display group information",
			"priv", "display privileges",
			"all", "display all information",
		),
	)
}
