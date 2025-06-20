package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "transmission-daemon [flags]",
	Short: "A daemon-based BitTorrent client",
	Long:  "https://transmissionbt.com/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringSliceP("allowed", "a", []string{"127.0.0.1", "::1"}, "Allow RPC access to a comma-delimited whiteliste of IP addresses")
	rootCmd.Flags().BoolP("auth", "t", false, "Require clients to authenticate themselves")
	rootCmd.Flags().StringP("bind-address-ipv4", "i", "0.0.0.0", "Listen for IPv4 connections on a specific address")
	rootCmd.Flags().StringP("bind-address-ipv6", "I", "::/0", "Listen for IPv6 connections on a specific address")
	rootCmd.Flags().BoolP("blocklist", "b", false, "Enable peer blocklists")
	rootCmd.Flags().StringP("config-dir", "g", "", "Directory to look for configuratinon files in")
	rootCmd.Flags().BoolP("dht", "o", false, "Enable distributed hash table")
	rootCmd.Flags().StringP("download-dir", "w", "", "Where to save downloaded data")
	rootCmd.Flags().BoolP("dump", "d", false, "Dump transmission-daemon's settings to stderr")
	rootCmd.Flags().BoolP("encryption-preferred", "ep", false, "Prefer encrypted peer connections")
	rootCmd.Flags().BoolP("encryption-required", "er", false, "Encrypt all peer connections")
	rootCmd.Flags().BoolP("encryption-tolerated", "et", false, "Prefer unencrypted peer connections")
	rootCmd.Flags().BoolP("foreground", "f", false, "Run in the foreground and print errors to stderr")
	rootCmd.Flags().Float64S("global-seedratio", "gsr", 0, "All torrents, unless overriden by a per-torrent setting, should seed until this ratio")
	rootCmd.Flags().BoolP("help", "h", false, "Prints a short usage summary")
	rootCmd.Flags().String("incomplete-dir", "", "When adding new torrents, store their contents in directory until torrent is done")
	rootCmd.Flags().Bool("log-debug", false, "Show error, info, and debug messages")
	rootCmd.Flags().Bool("log-error", false, "Show error messages")
	rootCmd.Flags().Bool("log-info", false, "Show error and info messages")
	rootCmd.Flags().StringP("logfile", "e", "", "Where to send log output")
	rootCmd.Flags().BoolP("no-auth", "T", false, "Don't require authentication from clients")
	rootCmd.Flags().BoolP("no-blocklist", "B", false, "Disable peer blocklists")
	rootCmd.Flags().BoolP("no-dht", "O", false, "Disable distributed hash table")
	rootCmd.Flags().BoolP("no-global-seedratio", "GSR", false, "All torrents, unless overriden by a per-torrent setting, should seed regardless of ratio")
	rootCmd.Flags().Bool("no-incomplete-dir", false, "Don't store incomplete torrents in a different directory")
	rootCmd.Flags().BoolP("no-portmap", "M", true, "Disable portmapping")
	rootCmd.Flags().Bool("no-utp", false, "Disable uTP for peer connections")
	rootCmd.Flags().BoolS("no-watch", "C", false, "Do not watch for new .torrent files")
	rootCmd.Flags().StringP("password", "v", "", "Client authentication password")
	rootCmd.Flags().Bool("paused", false, "Pause all torrents on startup")
	rootCmd.Flags().IntP("peerlimit-global", "L", 240, "Overall peer limit")
	rootCmd.Flags().IntP("peerlimit-torrent", "l", 60, "Peer limit per torrent")
	rootCmd.Flags().IntP("peerport", "P", 51413, "Set the port to listen for incoming peers")
	rootCmd.Flags().StringP("pid-file", "x", "", "Name of the daemon PID file")
	rootCmd.Flags().IntP("port", "p", 9091, "Port to open and listen for RPC requests on")
	rootCmd.Flags().BoolP("portmap", "m", false, "Enable port mapping via NAT-PMP or UPnP")
	rootCmd.Flags().StringP("rpc-bind-address", "r", "0.0.0.0", "Listen for RPC connections on a specific IPv4 or IPv6 address")
	rootCmd.Flags().StringP("username", "u", "", "Client authentication username")
	rootCmd.Flags().Bool("utp", false, "Enable uTP for peer connections")
	rootCmd.Flags().BoolP("version", "V", false, "Show version number and exit")
	rootCmd.Flags().StringS("watch", "c", "", "Directory to watch for new .torrent files to be added")

	rootCmd.MarkFlagsMutuallyExclusive("auth", "no-auth")
	rootCmd.MarkFlagsMutuallyExclusive("blocklist", "no-blocklist")
	rootCmd.MarkFlagsMutuallyExclusive("dht", "no-dht")
	rootCmd.MarkFlagsMutuallyExclusive("encryption-required", "encryption-preferred", "encryption-tolerated")
	rootCmd.MarkFlagsMutuallyExclusive("global-seedratio", "no-global-seedratio")
	rootCmd.MarkFlagsMutuallyExclusive("incomplete-dir", "no-incomplete-dir")
	rootCmd.MarkFlagsMutuallyExclusive("log-error", "log-info", "log-debug")
	rootCmd.MarkFlagsMutuallyExclusive("portmap", "no-portmap")
	rootCmd.MarkFlagsMutuallyExclusive("utp", "no-utp")
	rootCmd.MarkFlagsMutuallyExclusive("watch", "no-watch")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"allowed":           net.ActionIpv4Addresses(),
		"bind-address-ipv4": net.ActionIpv4Addresses(),
		"bind-address-ipv6": carapace.ActionValues("::0/0", "fe80::/10"),
		"config-dir":        carapace.ActionDirectories().Chdir("/"),
		"download-dir":      carapace.ActionDirectories().Chdir("/"),
		"logfile":           carapace.ActionFiles().Chdir("/"),
		"pid-file":          carapace.ActionFiles().Chdir("/"),
		"port":              net.ActionPorts(),
		"rpc-bind-address":  net.ActionIpv4Addresses(),
		"watch":             carapace.ActionDirectories().Chdir("/"),
	})
}
