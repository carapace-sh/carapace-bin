package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "blueutil",
	Short: "Bluetooth control utility",
	Long:  "https://github.com/toy/blueutil",
	Run:   func(*cobra.Command, []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("add-favorite", "", "add device to favourites")
	rootCmd.Flags().String("add-favourite", "", "add device to favourites")
	rootCmd.Flags().String("connect", "", "create a connection to a device")
	rootCmd.Flags().Bool("connected", false, "list connected devices")
	rootCmd.Flags().StringP("discoverable", "d", "", "output or set discoverable state")
	rootCmd.Flags().Bool("favorites", false, "list favourite devices")
	rootCmd.Flags().Bool("favourites", false, "list favourite devices")
	rootCmd.Flags().String("format", "", "change output format of info and listing commands")
	rootCmd.Flags().BoolP("help", "h", false, "show help")
	rootCmd.Flags().String("info", "", "show information about a device")
	rootCmd.Flags().String("inquiry", "", "inquiry devices in range for T seconds")
	rootCmd.Flags().String("is-connected", "", "output connected state of device as 1 or 0")
	rootCmd.Flags().String("pair", "", "pair with a device (optional PIN)")
	rootCmd.Flags().Bool("paired", false, "list paired devices")
	rootCmd.Flags().StringP("power", "p", "", "output or set power state")
	rootCmd.Flags().Bool("recent", false, "list recently used devices")
	rootCmd.Flags().String("remove-favorite", "", "remove device from favourites")
	rootCmd.Flags().String("remove-favourite", "", "remove device from favourites")
	rootCmd.Flags().String("unpair", "", "unpair a device")
	rootCmd.Flags().BoolP("version", "v", false, "show version")
	rootCmd.Flags().String("wait-connect", "", "wait for device to connect")
	rootCmd.Flags().String("wait-disconnect", "", "wait for device to disconnect")
	rootCmd.Flags().String("wait-rssi", "", "wait for device RSSI value")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"discoverable": carapace.ActionValues("1", "0", "on", "off", "toggle"),
		"format":       carapace.ActionValues("default", "new-default", "json", "json-pretty"),
		"power":        carapace.ActionValues("1", "0", "on", "off", "toggle"),
	})
}
