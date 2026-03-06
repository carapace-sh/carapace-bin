package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var report_sessionsCmd = &cobra.Command{
	Use:   "sessions",
	Short: "Reports the previous sessions (unstable)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(report_sessionsCmd).Standalone()

	report_sessionsCmd.Flags().BoolP("help", "h", false, "Print help")
	report_sessionsCmd.Flags().String("limit", "", "Limit the number of results")
	report_sessionsCmd.Flags().String("manifest-path", "", "Path to Cargo.toml")
	reportCmd.AddCommand(report_sessionsCmd)

	carapace.Gen(report_sessionsCmd).FlagCompletion(carapace.ActionMap{
		"manifest-path": carapace.ActionDirectories(),
	})
}
