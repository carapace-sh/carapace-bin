package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var report_timingsCmd = &cobra.Command{
	Use:   "timings",
	Short: "Reports the build timings of previous sessions (unstable)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(report_timingsCmd).Standalone()

	report_timingsCmd.Flags().BoolP("help", "h", false, "Print help")
	report_timingsCmd.Flags().String("id", "", "Session ID to report on")
	report_timingsCmd.Flags().String("manifest-path", "", "Path to Cargo.toml")
	report_timingsCmd.Flags().Bool("open", false, "Opens the timing report in a browser")
	reportCmd.AddCommand(report_timingsCmd)

	carapace.Gen(report_timingsCmd).FlagCompletion(carapace.ActionMap{
		"manifest-path": carapace.ActionDirectories(),
	})
}
