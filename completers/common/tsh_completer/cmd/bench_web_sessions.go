package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bench_web_sessionsCmd = &cobra.Command{
	Use:    "sessions",
	Short:  "Run session benchmark tests.",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bench_web_sessionsCmd).Standalone()

	bench_web_sessionsCmd.Flags().String("max", "", "The maximum number of sessions to open. If not specified a single session per node will be opened.")
	bench_webCmd.AddCommand(bench_web_sessionsCmd)
}
