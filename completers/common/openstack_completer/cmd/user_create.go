package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var user_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new user",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(user_createCmd).Standalone()

	user_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	user_createCmd.Flags().String("description", "", "User description")
	user_createCmd.Flags().Bool("disable", false, "Disable user")
	user_createCmd.Flags().Bool("disable-lock-password", false, "Enables the ability for a user to change its password through self-service APIs")
	user_createCmd.Flags().Bool("disable-multi-factor-auth", false, "Disables the MFA (Multi Factor Auth)")
	user_createCmd.Flags().String("domain", "", "Default domain (name or ID)")
	user_createCmd.Flags().String("email", "", "Set user email address")
	user_createCmd.Flags().Bool("enable", false, "Enable user (default)")
	user_createCmd.Flags().Bool("enable-lock-password", false, "Disables the ability for a user to change its password through self-service APIs")
	user_createCmd.Flags().Bool("enable-multi-factor-auth", false, "Enables the MFA (Multi Factor Auth)")
	user_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	user_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	user_createCmd.Flags().Bool("ignore-change-password-upon-first-use", false, "Control if a user should be forced to change their password immediately after they log into keystone for the first time.")
	user_createCmd.Flags().Bool("ignore-lockout-failure-attempts", false, "Opt into ignoring the number of times a user has authenticated and locking out the user as a result")
	user_createCmd.Flags().Bool("ignore-password-expiry", false, "Opt into allowing user to continue using passwords that may be expired")
	user_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	user_createCmd.Flags().String("multi-factor-auth-rule", "", "Set multi-factor auth rules.")
	user_createCmd.Flags().Bool("no-ignore-change-password-upon-first-use", false, "Control if a user should be forced to change their password immediately after they log into keystone for the first time.")
	user_createCmd.Flags().Bool("no-ignore-lockout-failure-attempts", false, "Opt out of ignoring the number of times a user has authenticated and locking out the user as a result")
	user_createCmd.Flags().Bool("no-ignore-password-expiry", false, "Opt out of allowing user to continue using passwords that may be expired")
	user_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	user_createCmd.Flags().Bool("or-show", false, "Return existing user")
	user_createCmd.Flags().String("password", "", "Set user password")
	user_createCmd.Flags().Bool("password-prompt", false, "Prompt interactively for password")
	user_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	user_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	user_createCmd.Flags().String("project", "", "Default project (name or ID)")
	user_createCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	user_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	userCmd.AddCommand(user_createCmd)
}
