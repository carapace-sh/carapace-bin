package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sysadminctl",
	Short: "system administrator control utility",
	Long:  "https://developer.apple.com/library/archive/documentation/Darwin/Reference/ManPages/man8/sysadminctl.8.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"-deleteUser", "Delete a user",
			"-newUser", "Create a new user",
			"-addAdmin", "Add an admin user",
			"-addGroups", "Add groups for a user",
			"-passwordPolicy", "Set password policy",
			"-mobileAccount", "Create mobile account",
			"-deleteMobileAccount", "Delete mobile account",
			"-secureTokenStatus", "Check secure token status",
			"-admin", "Admin operations",
			"-guestUser", "Guest user operations",
		),
	)
}
