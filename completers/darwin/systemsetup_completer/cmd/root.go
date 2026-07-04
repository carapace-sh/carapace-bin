package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "systemsetup",
	Short: "configuration tool for system settings",
	Long:  "https://keith.github.io/xcode-manpages/systemsetup.8.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"-getdate", "Get the current date",
			"-setdate", "Set the date (mm:dd:yy)",
			"-gettime", "Get the current time",
			"-settime", "Set the time (hh:mm:ss)",
			"-gettimezone", "Get the current time zone",
			"-listtimezones", "List available time zones",
			"-settimezone", "Set the time zone",
			"-getusingnetworktime", "Get network time status",
			"-setusingnetworktime", "Enable/disable network time",
			"-getnetworktimeserver", "Get network time server",
			"-setnetworktimeserver", "Set network time server",
			"-getsleep", "Get sleep settings",
			"-setsleep", "Set sleep timer",
			"-getcomputersleep", "Get computer sleep timer",
			"-setcomputersleep", "Set computer sleep timer",
			"-getdisplaysleep", "Get display sleep timer",
			"-setdisplaysleep", "Set display sleep timer",
			"-getharddisksleep", "Get hard disk sleep timer",
			"-setharddisksleep", "Set hard disk sleep timer",
			"-getwakeonmodem", "Get wake on modem status",
			"-setwakeonmodem", "Enable/disable wake on modem",
			"-getwakeonnetworkaccess", "Get wake on network access status",
			"-setwakeonnetworkaccess", "Enable/disable wake on network access",
			"-getrestartpowerfailure", "Get restart on power failure status",
			"-setrestartpowerfailure", "Enable/disable restart on power failure",
			"-getrestartfreeze", "Get restart on freeze status",
			"-setrestartfreeze", "Enable/disable restart on freeze",
			"-getallowpowerbuttontosleep", "Get power button sleep status",
			"-setallowpowerbuttontosleep", "Enable/disable power button sleep",
			"-getremotemanagement", "Get remote management status",
			"-setremotemanagement", "Enable/disable remote management",
			"-getremoteappleevents", "Get remote Apple events status",
			"-setremoteappleevents", "Enable/disable remote Apple events",
			"-getremotelogin", "Get remote login status",
			"-setremotelogin", "Enable/disable remote login",
			"-getcomputername", "Get computer name",
			"-setcomputername", "Set computer name",
			"-getlocalsubnetname", "Get local subnet name",
			"-setlocalsubnetname", "Set local subnet name",
			"-getstartupdisk", "Get startup disk",
			"-setstartupdisk", "Set startup disk",
			"-liststartupdisks", "List available startup disks",
		),
	)
}
