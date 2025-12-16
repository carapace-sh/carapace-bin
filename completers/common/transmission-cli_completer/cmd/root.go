package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/transmission"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "transmission-cli [flags] torrent-file",
	Short: "A fast and easy BitTorrent client",
	Long:  "https://transmissionbt.com/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute(opts ...func(cmd *cobra.Command)) error {
	for _, opt := range opts {
		opt(rootCmd)
	}
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("blocklist", "b", false, "Enable peer blocklists")
	rootCmd.Flags().StringP("config-dir", "g", "", "Directory to look for configuratinon files in")
	rootCmd.Flags().IntP("downlimit", "d", 0, "Set the maximum download speed in KB/s")
	rootCmd.Flags().StringP("download-dir", "w", "", "Where to save downloaded data")
	rootCmd.Flags().BoolP("encryption-preferred", "ep", false, "Prefer encrypted peer connections")
	rootCmd.Flags().BoolP("encryption-required", "er", false, "Encrypt all peer connections")
	rootCmd.Flags().BoolP("encryption-tolerated", "et", false, "Prefer unencrypted peer connections")
	rootCmd.Flags().StringP("finish", "f", "", "Set a script to run when the torrent finishes")
	rootCmd.Flags().BoolP("help", "h", false, "Prints a short usage summary")
	rootCmd.Flags().BoolP("no-blocklist", "B", false, "Disable peer blocklists")
	rootCmd.Flags().BoolP("no-downlimit", "D", true, "Don't limit the download speed")
	rootCmd.Flags().BoolP("no-portmap", "M", true, "Disable portmapping")
	rootCmd.Flags().BoolP("no-uplimit", "U", true, "Don't limit the upload speed")
	rootCmd.Flags().IntP("port", "p", 51413, "Set the port to listen for incoming peers")
	rootCmd.Flags().BoolP("portmap", "m", false, "Enable port mapping via NAT-PMP or UPnP")
	rootCmd.Flags().String("tos", "", "Use a ToS or DSCP name or value from 0-255 to set the peer socket service type")
	rootCmd.Flags().IntP("uplimit", "u", 0, "Set the maximum upload speed in KB/s")
	rootCmd.Flags().BoolP("verify", "v", true, "Verify the torrent's downloaded data")
	rootCmd.Flags().BoolP("version", "V", false, "Show version number and exit")

	rootCmd.MarkFlagsMutuallyExclusive("blocklist", "no-blocklist")
	rootCmd.MarkFlagsMutuallyExclusive("downlimit", "no-downlimit")
	rootCmd.MarkFlagsMutuallyExclusive("encryption-required", "encryption-preferred", "encryption-tolerated")
	rootCmd.MarkFlagsMutuallyExclusive("portmap", "no-portmap")
	rootCmd.MarkFlagsMutuallyExclusive("uplimit", "no-uplimit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"config-dir":   carapace.ActionDirectories().Chdir("/"),
		"download-dir": carapace.ActionDirectories().Chdir("/"),
		"finish":       carapace.ActionFiles().Chdir("/"),
		"port":         net.ActionPorts(),
		"tos":          transmission.ActionTOS(),
	})
	carapace.Gen(rootCmd).PositionalCompletion(carapace.ActionFiles(".torrent", ".magnet"))
}
