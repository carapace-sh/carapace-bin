package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
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

	rootCmd.Flags().String("GID", "", "Group ID for new user")
	rootCmd.Flags().String("UID", "", "User ID for new user")
	rootCmd.Flags().String("adminPassword", "", "Administrator password (use '-' for interactive prompt)")
	rootCmd.Flags().String("adminUser", "", "Administrator user name")
	rootCmd.Flags().String("afpGuestAccess", "", "AFP guest access: on, off, or status")
	rootCmd.Flags().String("automaticTime", "", "Automatic time: on, off, or status")
	rootCmd.Flags().BoolP("deleteUser", "D", false, "Delete a user")
	rootCmd.Flags().String("filesystem", "", "Filesystem status")
	rootCmd.Flags().String("fullName", "", "Full name for new user")
	rootCmd.Flags().String("guestAccount", "", "Guest account: on, off, or status")
	rootCmd.Flags().String("hint", "", "Password hint")
	rootCmd.Flags().String("home", "", "Full path to home directory")
	rootCmd.Flags().String("newPassword", "", "New password")
	rootCmd.Flags().String("oldPassword", "", "Old password")
	rootCmd.Flags().String("password", "", "User password")
	rootCmd.Flags().String("passwordHint", "", "Password hint")
	rootCmd.Flags().String("picture", "", "Full path to user image")
	rootCmd.Flags().String("resetPasswordFor", "", "Reset password for local user")
	rootCmd.Flags().String("roleAccount", "", "Role account")
	rootCmd.Flags().String("screenLock", "", "Screen lock: status, immediate, off, or seconds")
	rootCmd.Flags().String("secureTokenOff", "", "Disable secure token for user")
	rootCmd.Flags().String("secureTokenOn", "", "Enable secure token for user")
	rootCmd.Flags().String("secureTokenStatus", "", "Check secure token status for user")
	rootCmd.Flags().String("shell", "", "Path to shell")
	rootCmd.Flags().String("smbGuestAccess", "", "SMB guest access: on, off, or status")
	rootCmd.Flags().String("use12HourClockForLoginWindow", "", "12-hour clock: on, off, or status")
	rootCmd.Flags().String("userName", "", "User name")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"afpGuestAccess":               carapace.ActionValues("on", "off", "status"),
		"automaticTime":                carapace.ActionValues("on", "off", "status"),
		"deleteUser":                   os.ActionUsers(),
		"guestAccount":                 carapace.ActionValues("on", "off", "status"),
		"secureTokenStatus":            os.ActionUsers(),
		"smbGuestAccess":               carapace.ActionValues("on", "off", "status"),
		"use12HourClockForLoginWindow": carapace.ActionValues("on", "off", "status"),
	})
}
