package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var daemonCmd = &cobra.Command{
	Use:     "daemon",
	Short:   "A really simple server for Git repositories",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_synching].ID,
}

func init() {
	carapace.Gen(daemonCmd).Standalone()

	daemonCmd.Flags().String("access-hook", "", "Every time a client connects, first run an external command specified by the <path>")
	daemonCmd.Flags().String("allow-override", "", "Allow overriding the site-wide default with per repository configuration")
	daemonCmd.Flags().String("base-path", "", "Remap all the path requests as relative to the given path")
	daemonCmd.Flags().Bool("base-path-relaxed", false, "Attempt to lookup without prefixing the base path")
	daemonCmd.Flags().Bool("detach", false, "Detach from the shell")
	daemonCmd.Flags().String("disable", "", "Disable the service site-wide per default")
	daemonCmd.Flags().String("enable", "", "Enablethe service site-wide per default")
	daemonCmd.Flags().Bool("export-all", false, "Allow pulling from all directories that look like Git repositories")
	daemonCmd.Flags().String("forbid-override", "", "Forbid overriding the site-wide default with per repository configuration")
	daemonCmd.Flags().String("group", "", "Change daemon’s gid before entering the service loop")
	daemonCmd.Flags().Bool("inetd", false, "Have the server run as an inetd service")
	daemonCmd.Flags().Bool("informative-errors", false, "Report more verbose errors to the client")
	daemonCmd.Flags().String("init-timeout", "", "Timeout (in seconds) between the moment the connection is established and the client request is received")
	daemonCmd.Flags().String("interpolated-path", "", "Use an interpolated path template can be used to dynamically construct alternate paths")
	daemonCmd.Flags().String("listen", "", "Listen on a specific IP address or hostname")
	daemonCmd.Flags().String("log-destination", "", "Send log messages to the specified destination")
	daemonCmd.Flags().String("max-connections", "", "Maximum number of concurrent clients")
	daemonCmd.Flags().Bool("no-informative-errors", false, "Do not report more verbose errors to the client")
	daemonCmd.Flags().String("pid-file", "", "Save the process id in file")
	daemonCmd.Flags().String("port", "", "Listen on an alternative port")
	daemonCmd.Flags().Bool("reuseaddr", false, "Use SO_REUSEADDR when binding the listening socket")
	daemonCmd.Flags().Bool("strict-paths", false, "Match paths exactly and don’t do user-relative paths")
	daemonCmd.Flags().Bool("syslog", false, "Short for --log-destination=syslog")
	daemonCmd.Flags().String("timeout", "", "Timeout (in seconds) for specific client sub-requests")
	daemonCmd.Flags().String("user", "", "Change daemon’s uid before entering the service loop")
	daemonCmd.Flags().String("user-path", "", "Allow ~user notation to be used in requests")
	daemonCmd.Flags().Bool("verbose", false, "Log details about the incoming connections and requested files")
	rootCmd.AddCommand(daemonCmd)

	daemonCmd.Flag("user-path").NoOptDefVal = " "

	carapace.Gen(daemonCmd).FlagCompletion(carapace.ActionMap{
		"access-hook":     bridge.ActionCarapaceBin().Split(),
		"allow-override":  carapace.ActionValues(), // TODO
		"base-path":       carapace.ActionDirectories(),
		"disable":         carapace.ActionValues(), // TODO
		"enable":          carapace.ActionValues(), // TODO
		"forbid-override": carapace.ActionValues(), // TODO
		"group":           os.ActionGroups(),
		"log-destination": carapace.ActionValuesDescribed(
			"stderr", "Write to standard error",
			"syslog", "Write to syslog, using the git-daemon identifier",
			"none", "Disable all logging",
		),
		"pid-file": carapace.ActionFiles(),
		"port":     net.ActionPorts(),
		"user":     os.ActionUsers(),
	})

	carapace.Gen(daemonCmd).PositionalAnyCompletion(
		carapace.ActionDirectories(),
	)
}
