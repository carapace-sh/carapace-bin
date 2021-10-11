package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dms",
	Short: "A UPnP DLNA Digital Media Server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	carapace.Override(carapace.Opts{
		LongShorthand: true,
	})
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("allowedIps", "", "allowed ip of clients, separated by comma")
	rootCmd.Flags().String("config", "", "json configuration file")
	rootCmd.Flags().String("deviceIcon", "", "device icon")
	rootCmd.Flags().String("fFprobeCachePath", "", "path to FFprobe cache file (default \"~/.dms-ffprobe-cache\")")
	rootCmd.Flags().String("forceTranscodeTo", "", "force transcoding to certain format, supported: 'chromecast', 'vp8', 'web'")
	rootCmd.Flags().String("friendlyName", "", "server friendly name")
	rootCmd.Flags().String("http", "", "http server port (default \":1338\")")
	rootCmd.Flags().String("ifname", "", "specific SSDP network interface")
	rootCmd.Flags().Bool("ignoreHidden", false, "ignore hidden files and directories")
	rootCmd.Flags().Bool("ignoreUnreadable", false, "ignore unreadable files and directories")
	rootCmd.Flags().Bool("logHeaders", false, "log HTTP headers")
	rootCmd.Flags().Bool("noProbe", false, "disable media probing with ffprobe")
	rootCmd.Flags().Bool("noTranscode", false, "disable transcoding")
	rootCmd.Flags().String("notifyInterval", "", "interval between SSPD announces (default 30s)")
	rootCmd.Flags().String("path", "", "browse root path")
	rootCmd.Flags().Bool("stallEventSubscribe", false, "workaround for some bad event subscribers")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"config":           carapace.ActionFiles(".json"),
		"fFprobeCachePath": carapace.ActionFiles(),
		"forceTranscodeTo": carapace.ActionValues("chromecast", "vp8", "web"),
		"ifname":           net.ActionDevices(net.IncludedDevices{Ethernet: true, Wifi: true}),
		"path":             carapace.ActionDirectories(),
	})
}
