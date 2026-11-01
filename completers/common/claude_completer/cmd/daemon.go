package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "Manage the background-session supervisor",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(daemonCmd).Standalone()

	daemonCmd.Flags().String("json-path", "", "Config file (default: ~/.claude/daemon.json)")
	daemonCmd.Flags().String("log-file", "", "Log file (default: ~/.claude/daemon.log)")
	rootCmd.AddCommand(daemonCmd)

	carapace.Gen(daemonCmd).FlagCompletion(carapace.ActionMap{
		"json-path": carapace.ActionFiles(),
		"log-file":  carapace.ActionFiles(),
	})
}
