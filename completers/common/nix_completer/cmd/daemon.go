package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var daemonCmd = &cobra.Command{
	Use:     "daemon",
	Short:   "daemon to perform store operations on behalf of non-root clients",
	GroupID: "utility",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(daemonCmd).Standalone()

	daemonCmd.Flags().Bool("default-trust", false, "Use Nix's default trust")
	daemonCmd.Flags().Bool("force-trusted", false, "Force the daemon to trust connecting clients")
	daemonCmd.Flags().Bool("force-untrusted", false, "Force the daemon to not trust connecting clients")
	daemonCmd.Flags().Bool("stdio", false, "Attach to standard I/O instead of trying to bind to a Unix socket")

	daemonCmd.MarkFlagsMutuallyExclusive("default-trust", "force-trusted", "force-untrusted")

	addEvaluationFlags(daemonCmd)
	addLoggingFlags(daemonCmd)

	rootCmd.AddCommand(daemonCmd)
}
