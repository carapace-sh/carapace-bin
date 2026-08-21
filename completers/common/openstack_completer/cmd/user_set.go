package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var user_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set user properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(user_setCmd).Standalone()

	user_setCmd.Flags().String("description", "", "Set user description")
	user_setCmd.Flags().Bool("disable", false, "Disable user")
	user_setCmd.Flags().Bool("disable-lock-password", false, "Enables the ability for a user to change its password through self-service APIs")
	user_setCmd.Flags().Bool("disable-multi-factor-auth", false, "Disables the MFA (Multi Factor Auth)")
	user_setCmd.Flags().String("domain", "", "Domain the user belongs to (name or ID).")
	user_setCmd.Flags().String("email", "", "Set user email address")
	user_setCmd.Flags().Bool("enable", false, "Enable user (default)")
	user_setCmd.Flags().Bool("enable-lock-password", false, "Disables the ability for a user to change its password through self-service APIs")
	user_setCmd.Flags().Bool("enable-multi-factor-auth", false, "Enables the MFA (Multi Factor Auth)")
	user_setCmd.Flags().Bool("ignore-change-password-upon-first-use", false, "Control if a user should be forced to change their password immediately after they log into keystone for the first time.")
	user_setCmd.Flags().Bool("ignore-lockout-failure-attempts", false, "Opt into ignoring the number of times a user has authenticated and locking out the user as a result")
	user_setCmd.Flags().Bool("ignore-password-expiry", false, "Opt into allowing user to continue using passwords that may be expired")
	user_setCmd.Flags().String("multi-factor-auth-rule", "", "Set multi-factor auth rules.")
	user_setCmd.Flags().String("name", "", "Set user name")
	user_setCmd.Flags().Bool("no-ignore-change-password-upon-first-use", false, "Control if a user should be forced to change their password immediately after they log into keystone for the first time.")
	user_setCmd.Flags().Bool("no-ignore-lockout-failure-attempts", false, "Opt out of ignoring the number of times a user has authenticated and locking out the user as a result")
	user_setCmd.Flags().Bool("no-ignore-password-expiry", false, "Opt out of allowing user to continue using passwords that may be expired")
	user_setCmd.Flags().String("password", "", "Set user password")
	user_setCmd.Flags().Bool("password-prompt", false, "Prompt interactively for password")
	user_setCmd.Flags().String("project", "", "Set default project (name or ID)")
	user_setCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	userCmd.AddCommand(user_setCmd)
}
