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

	rootCmd.Flags().Bool("getallowpowerbuttontosleep", false, "Get power button sleep status")
	rootCmd.Flags().Bool("getcomputername", false, "Get computer name")
	rootCmd.Flags().Bool("getcomputersleep", false, "Get computer sleep timer")
	rootCmd.Flags().Bool("getdate", false, "Get the current date")
	rootCmd.Flags().Bool("getdisplaysleep", false, "Get display sleep timer")
	rootCmd.Flags().Bool("getharddisksleep", false, "Get hard disk sleep timer")
	rootCmd.Flags().Bool("getlocalsubnetname", false, "Get local subnet name")
	rootCmd.Flags().Bool("getnetworktimeserver", false, "Get network time server")
	rootCmd.Flags().Bool("getremoteappleevents", false, "Get remote Apple events status")
	rootCmd.Flags().Bool("getremotelogin", false, "Get remote login status")
	rootCmd.Flags().Bool("getremotemanagement", false, "Get remote management status")
	rootCmd.Flags().Bool("getrestartfreeze", false, "Get restart on freeze status")
	rootCmd.Flags().Bool("getrestartpowerfailure", false, "Get restart on power failure status")
	rootCmd.Flags().Bool("getsleep", false, "Get sleep settings")
	rootCmd.Flags().Bool("getstartupdisk", false, "Get startup disk")
	rootCmd.Flags().Bool("gettime", false, "Get the current time")
	rootCmd.Flags().Bool("gettimezone", false, "Get the current time zone")
	rootCmd.Flags().Bool("getusingnetworktime", false, "Get network time status")
	rootCmd.Flags().Bool("getwakeonmodem", false, "Get wake on modem status")
	rootCmd.Flags().Bool("getwakeonnetworkaccess", false, "Get wake on network access status")
	rootCmd.Flags().Bool("liststartupdisks", false, "List available startup disks")
	rootCmd.Flags().Bool("listtimezones", false, "List available time zones")
	rootCmd.Flags().String("setallowpowerbuttontosleep", "", "Enable/disable power button sleep (on/off)")
	rootCmd.Flags().String("setcomputername", "", "Set computer name")
	rootCmd.Flags().String("setcomputersleep", "", "Set computer sleep timer (minutes or Never)")
	rootCmd.Flags().String("setdate", "", "Set the date (mm:dd:yy)")
	rootCmd.Flags().String("setdisplaysleep", "", "Set display sleep timer (minutes or Never)")
	rootCmd.Flags().String("setharddisksleep", "", "Set hard disk sleep timer (minutes or Never)")
	rootCmd.Flags().String("setlocalsubnetname", "", "Set local subnet name")
	rootCmd.Flags().String("setnetworktimeserver", "", "Set network time server")
	rootCmd.Flags().String("setremoteappleevents", "", "Enable/disable remote Apple events (on/off)")
	rootCmd.Flags().String("setremotelogin", "", "Enable/disable remote login (on/off)")
	rootCmd.Flags().String("setremotemanagement", "", "Enable/disable remote management (on/off)")
	rootCmd.Flags().String("setrestartfreeze", "", "Enable/disable restart on freeze (on/off)")
	rootCmd.Flags().String("setrestartpowerfailure", "", "Enable/disable restart on power failure (on/off)")
	rootCmd.Flags().String("setsleep", "", "Set sleep timer (minutes or Never)")
	rootCmd.Flags().String("setstartupdisk", "", "Set startup disk path")
	rootCmd.Flags().String("settime", "", "Set the time (hh:mm:ss)")
	rootCmd.Flags().String("settimezone", "", "Set the time zone")
	rootCmd.Flags().String("setusingnetworktime", "", "Enable/disable network time (on/off)")
	rootCmd.Flags().String("setwakeonmodem", "", "Enable/disable wake on modem (on/off)")
	rootCmd.Flags().String("setwakeonnetworkaccess", "", "Enable/disable wake on network access (on/off)")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"setallowpowerbuttontosleep": carapace.ActionValues("on", "off"),
		"setremoteappleevents":       carapace.ActionValues("on", "off"),
		"setremotelogin":             carapace.ActionValues("on", "off"),
		"setremotemanagement":        carapace.ActionValues("on", "off"),
		"setrestartfreeze":           carapace.ActionValues("on", "off"),
		"setrestartpowerfailure":     carapace.ActionValues("on", "off"),
		"setstartupdisk":             carapace.ActionDirectories(),
		"setusingnetworktime":        carapace.ActionValues("on", "off"),
		"setwakeonmodem":             carapace.ActionValues("on", "off"),
		"setwakeonnetworkaccess":     carapace.ActionValues("on", "off"),
	})
}
