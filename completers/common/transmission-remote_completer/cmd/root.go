package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/transmission"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "transmission-remote [host[:port]] [flags]",
	Short: "A remote control utility for transmission-daemon and transmission",
	Long:  "https://transmissionbt.com/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringArrayP("add", "a", nil, "Add torrents")
	rootCmd.Flags().BoolP("alt-speed", "as", false, "Use the alternate limits")
	rootCmd.Flags().StringSlice("alt-speed-days", nil, "Set the numbers days on which to enable the alternate speed scheduler (e.g. \"2\", \"2,4-6\")")
	rootCmd.Flags().IntP("alt-speed-downlimit", "asd", 0, "Limit the alternate download speed in KB/s")
	rootCmd.Flags().BoolP("alt-speed-scheduler", "asc", false, "Use the alternate scheduled on/off times")
	rootCmd.Flags().Int("alt-speed-time-begin", 0, "Time to start using the alt speed limits (hhmm)")
	rootCmd.Flags().Int("alt-speed-time-end", 0, "Time to stop using the alt speed limits (hhmm)")
	rootCmd.Flags().IntP("alt-speed-uplimit", "asu", 0, "Limit the alternate upload speed in KB/s")
	rootCmd.Flags().StringP("auth", "n", "", "Set the username:password for authentication")
	rootCmd.Flags().BoolP("authenv", "ne", false, "Use the TR_AUTH environment variable for authentication")
	rootCmd.Flags().BoolP("bandwidth-high", "Bh", false, "Give this torrent first chance at available bandwith")
	rootCmd.Flags().BoolP("bandwidth-low", "Bl", false, "Give this torrent the bandwith leftover after high and normal priority torrents")
	rootCmd.Flags().BoolP("bandwidth-normal", "Bn", false, "Give this torrent the bandwidth leftover after high priority torrents")
	rootCmd.Flags().Bool("blocklist-update", false, "Update blocklist from URL in remote client's settings")
	rootCmd.Flags().IntP("cache", "e", 0, "Set the session's maximum memory cache size in MiB")
	rootCmd.Flags().BoolP("debug", "b", false, "Enable debugging mode")
	rootCmd.Flags().BoolP("dht", "o", false, "Enable distributed hash table")
	rootCmd.Flags().IntP("downlimit", "d", 0, "Limit the maximum download speed in KB/s")
	rootCmd.Flags().StringP("download-dir", "w", "", "Where to save downloaded data")
	rootCmd.Flags().BoolP("encryption-preferred", "ep", false, "Prefer encrypted peer connections")
	rootCmd.Flags().BoolP("encryption-required", "er", false, "Encrypt all peer connections")
	rootCmd.Flags().BoolP("encryption-tolerated", "et", false, "Prefer unencrypted peer connections")
	rootCmd.Flags().Bool("exit", false, "Initiate a shutdown")
	rootCmd.Flags().BoolP("files", "f", false, "Get a file list for the current torrent(s)")
	rootCmd.Flags().StringArrayP("filter", "F", nil, "Filter selected torrents (may be specified more than once)")
	rootCmd.Flags().String("find", "", "Tell transmission where to look for the current torrent's data")
	rootCmd.Flags().StringSliceP("get", "g", nil, "Mark file(s) for download. (e.g. \"all\", \"1\", \"1,3-5\")")
	rootCmd.Flags().Float64S("global-seedratio", "gsr", 0, "All torrents, unless overriden by a per-torrent setting, should seed until this ratio")
	rootCmd.Flags().BoolP("help", "h", false, "Prints a short usage summary")
	rootCmd.Flags().BoolP("honor-session", "hl", false, "Make the current torrent(s) honor the session limits")
	rootCmd.Flags().IntP("peers", "pr", 0, "Set the maximum number of peers globally, or on current torrents if selected")
	rootCmd.Flags().String("incomplete-dir", "c", "When adding new torrents, store their contents in directory until torrent is done")
	rootCmd.Flags().BoolP("info", "i", false, "Show details of current torrent(s)")
	rootCmd.Flags().BoolP("info-files", "if", false, "List the specified torrent's files")
	rootCmd.Flags().BoolP("info-peers", "ip", false, "List the specified torrent's peers")
	rootCmd.Flags().BoolP("info-pieces", "ic", false, "List the specified torrent's pieces")
	rootCmd.Flags().BoolP("info-trackers", "it", false, "List the specified torrent's trackers")
	rootCmd.Flags().BoolP("json", "j", false, "Return the RPC response as JSON")
	rootCmd.Flags().StringArrayP("labels", "L", nil, "Set the specified torrent's labels")
	rootCmd.Flags().BoolP("lds", "y", false, "Enable local peer discovery (LPD)")
	rootCmd.Flags().BoolP("list", "l", false, "List all torrents")
	rootCmd.Flags().String("move", "", "Move the current torrent's data to the specified directory")
	rootCmd.Flags().StringP("netrc", "N", "", "Use the specified netrc file for authentication")
	rootCmd.Flags().BoolP("no-alt-speed", "AS", false, "Don't use the alternate limits")
	rootCmd.Flags().BoolP("no-alt-speed-scheduler", "ASC", false, "Don't use the alternate scheduled on/off times")
	rootCmd.Flags().BoolP("no-dht", "O", false, "Disable distributed hash table")
	rootCmd.Flags().BoolP("no-downlimit", "D", false, "Don't limit the maximum download speed")
	rootCmd.Flags().StringSliceP("no-get", "G", nil, "Marke file(s) for not downloading (e.g. \"all\", \"1\", \"1,3-5\")")
	rootCmd.Flags().BoolS("no-global-seedratio", "GSR", false, "All torrents, unless overriden by a per-torrent setting, should seed regardless of ratio")
	rootCmd.Flags().BoolP("no-honor-session", "HL", false, "Make the current torrent(s) not honor the session limits")
	rootCmd.Flags().BoolP("no-incomplete-dir", "C", false, "Don't store incomplete torrents in a different directory")
	rootCmd.Flags().BoolP("no-lds", "Y", false, "Disable local peer discovery (LPD)")
	rootCmd.Flags().BoolP("no-pex", "X", false, "Disable peer exchange (PEX)")
	rootCmd.Flags().BoolP("no-portmap", "M", true, "Disable portmapping")
	rootCmd.Flags().BoolP("no-seedratio", "SR", false, "Let the current torrent(s) seed regardless of ratio")
	rootCmd.Flags().Bool("no-torrent-done-script", false, "Don't run any script when a torrent finishes")
	rootCmd.Flags().Bool("no-trash-torrent", false, "Don't delete torrent files after adding")
	rootCmd.Flags().BoolP("no-uplimit", "U", true, "Don't limit the upload speed")
	rootCmd.Flags().Bool("no-utp", false, "Disable uTP for peer connections")
	rootCmd.Flags().String("path", "", "Provide original path for the rename command")
	rootCmd.Flags().BoolP("peer-info", "pi", false, "List the current torrent's connected peers")
	rootCmd.Flags().BoolP("pex", "x", false, "Enable peer exchange (PEX)")
	rootCmd.Flags().BoolP("print-ids", "ids", false, "Print a list of the specified torrent's ids in a format suitable for -t")
	rootCmd.Flags().StringSliceP("priority-high", "ph", nil, "Try to download the specified file(s) first (e.g. \"all\", \"1\", \"1,3-5\")")
	rootCmd.Flags().StringSliceP("priority-low", "pl", nil, "Try to download the specified file(s) last (e.g. \"all\", \"1\", \"1,3-5\")")
	rootCmd.Flags().StringSliceP("priority-normal", "pn", nil, "Try to download the specified file(s) normally (e.g. \"all\", \"1\", \"1,3-5\")")
	rootCmd.Flags().IntP("port", "p", 51413, "Set the port to listen for incoming peers")
	rootCmd.Flags().BoolP("portmap", "m", false, "Enable port mapping via NAT-PMP or UPnP")
	rootCmd.Flags().Bool("reannounce", false, "Reannounce the current torrent(s)")
	rootCmd.Flags().BoolP("remove", "r", false, "Remove the current torrent(s) without deleting their downloaded data")
	rootCmd.Flags().BoolP("remove-and-delete", "rad", false, "Remove the current torrent(s) and delete their downloaded data")
	rootCmd.Flags().String("rename", "", "Rename files or root folder of a torrent")
	rootCmd.Flags().Float64P("seedratio", "sr", 0, "Let the current torrent(s) seed until a specific ratio")
	rootCmd.Flags().Bool("no-start-paused", false, "Start added torrents unpaused")
	rootCmd.Flags().BoolP("seedratio-default", "srd", false, "Let the current torrent(s) seed until the global seedratio")
	rootCmd.Flags().BoolP("session-info", "si", false, "List sessions from the server")
	rootCmd.Flags().BoolP("session-stats", "st", false, "List statistical information from the server")
	rootCmd.Flags().BoolP("start", "s", false, "Start the current torrent(s)")
	rootCmd.Flags().Bool("start-paused", false, "Start added torrents paused")
	rootCmd.Flags().BoolP("stop", "S", false, "Stop the current torrent(s) from downloading or seeding")
	rootCmd.Flags().StringSliceP("torrent", "t", nil, "Set the current torrent(s) for use by subsequent options")
	rootCmd.Flags().String("torrent-done-script", "", "Specify a file to run each time a torrent finishes")
	rootCmd.Flags().Bool("trash-torrent", false, "Delete torrent files after adding")
	rootCmd.Flags().IntP("uplimit", "u", 0, "Set the maximum upload speed in KB/s")
	rootCmd.Flags().Bool("utp", false, "Enable uTP for peer connections")
	rootCmd.Flags().BoolP("verify", "v", true, "Verify the torrent's downloaded data")
	rootCmd.Flags().BoolP("version", "V", false, "Show version number and exit")

	rootCmd.MarkFlagsRequiredTogether("path", "rename")

	rootCmd.MarkFlagsMutuallyExclusive("alt-speed", "no-alt-speed")
	rootCmd.MarkFlagsMutuallyExclusive("alt-speed-scheduler", "no-alt-speed-scheduler")
	rootCmd.MarkFlagsMutuallyExclusive("auth", "authenv", "netrc")
	rootCmd.MarkFlagsMutuallyExclusive("bandwidth-high", "bandwidth-low", "bandwidth-normal")
	rootCmd.MarkFlagsMutuallyExclusive("dht", "no-dht")
	rootCmd.MarkFlagsMutuallyExclusive("downlimit", "no-downlimit")
	rootCmd.MarkFlagsMutuallyExclusive("encryption-required", "encryption-preferred", "encryption-tolerated")
	rootCmd.MarkFlagsMutuallyExclusive("global-seedratio", "no-global-seedratio")
	rootCmd.MarkFlagsMutuallyExclusive("incomplete-dir", "no-incomplete-dir")
	rootCmd.MarkFlagsMutuallyExclusive("info", "info-files", "info-peers", "info-pieces", "info-trackers")
	rootCmd.MarkFlagsMutuallyExclusive("lds", "no-lds")
	rootCmd.MarkFlagsMutuallyExclusive("pex", "no-pex")
	rootCmd.MarkFlagsMutuallyExclusive("portmap", "no-portmap")
	rootCmd.MarkFlagsMutuallyExclusive("priority-high", "priority-low", "priority-normal")
	rootCmd.MarkFlagsMutuallyExclusive("seedratio", "no-seedratio", "seedratio-default")
	rootCmd.MarkFlagsMutuallyExclusive(
		"honor-session", "no-honor-session",
		"move",
		"peers",
		"reannounce",
		"remove", "remove-and-delete",
		"rename",
		"start",
		"stop",
		"uplimit", "no-uplimit",
	)
	rootCmd.MarkFlagsMutuallyExclusive("start-paused", "no-start-paused")
	rootCmd.MarkFlagsMutuallyExclusive("torrent-done-script", "no-torrent-done-script")
	rootCmd.MarkFlagsMutuallyExclusive("trash-torrent", "no-trash-torrent")
	rootCmd.MarkFlagsMutuallyExclusive("utp", "no-utp")

	markFlagsNoCurrentTorret(
		"add",
		"alt-speed",
		"alt-speed-days",
		"alt-speed-downlimit",
		"alt-speed-scheduler",
		"alt-speed-time-begin",
		"alt-speed-time-end",
		"alt-speed-uplimit",
		"blocklist-update",
		"cache",
		"dht",
		"download-dir",
		"exit",
		"encryption-preferred",
		"encryption-required",
		"encryption-tolerated",
		"global-seedratio",
		"pex",
		"incomplete-dir",
		"lds",
		"no-alt-speed",
		"no-alt-speed-scheduler",
		"no-dht",
		"no-global-seedratio",
		"no-incomplete-dir",
		"no-lds",
		"no-pex",
		"no-portmap",
		"no-torrent-done-script",
		"no-trash-torrent",
		"no-utp",
		"port",
		"portmap",
		"session-info",
		"session-stats",
		"torrent-done-script",
		"trash-torrent",
		"utp",
		"verify",
		"version",
	)

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"add":            carapace.ActionFiles(".torrent", ".magnet"),
		"filter":         transmission.ActionFilters(),
		"find":           carapace.ActionDirectories().Chdir("/"),
		"get":            carapace.ActionValuesDescribed("all", "Get all files").StyleF(style.ForKeyword),
		"incomplete-dir": carapace.ActionDirectories().Chdir("/"),
		"move":           carapace.ActionDirectories().Chdir("/"),
		"netrc":          carapace.ActionFiles(),
		"no-get":         carapace.ActionValuesDescribed("all", "Get all files").StyleF(style.ForKeyword),
		"port":           net.ActionKnownPorts(),
		"torrent": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			filters, _ := rootCmd.Flags().GetStringArray("filter")
			return transmission.ActionIds(filters)
		}),
		"torrent-done-script": carapace.ActionDirectories().Chdir("/"),
	})
	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionMultiPartsN(":", 2, func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return net.ActionHosts()
			default:
				return net.ActionKnownPorts()
			}
		}),
	)
}

// Marks the flags so they're exclusive with --torrent
func markFlagsNoCurrentTorret(flags ...string) {
	for _, flag := range flags {
		rootCmd.MarkFlagsMutuallyExclusive(flag, "torrent")
	}
}
