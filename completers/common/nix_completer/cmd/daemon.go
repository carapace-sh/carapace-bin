package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/nix_completer/cmd/common"
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
	daemonCmd.Flags().Bool("process-ops", false, "Forces the daemon to process received commands itself rather than forwarding the commands straight to the remote store")
	daemonCmd.Flags().String("socket-path", "", "Path to the daemon's UNIX socket")
	daemonCmd.Flags().Bool("stdio", false, "Attach to standard I/O, instead of using UNIX socket(s)")

	daemonCmd.MarkFlagsMutuallyExclusive("default-trust", "force-trusted", "force-untrusted")
	daemonCmd.MarkFlagsMutuallyExclusive("socket-path", "stdio")

	common.AddEvaluationFlags(daemonCmd)
	common.AddLoggingFlags(daemonCmd)

	carapace.Gen(daemonCmd).FlagCompletion(carapace.ActionMap{
		"socket-path": carapace.ActionFiles(),
	})

	rootCmd.AddCommand(daemonCmd)
}
