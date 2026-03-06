package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var report_rebuildsCmd = &cobra.Command{
	Use:   "rebuilds",
	Short: "Reports rebuild reasons from previous sessions (unstable)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(report_rebuildsCmd).Standalone()

	report_rebuildsCmd.Flags().BoolP("help", "h", false, "Print help")
	report_rebuildsCmd.Flags().String("id", "", "Session ID to report on")
	report_rebuildsCmd.Flags().String("manifest-path", "", "Path to Cargo.toml")
	reportCmd.AddCommand(report_rebuildsCmd)

	carapace.Gen(report_rebuildsCmd).FlagCompletion(carapace.ActionMap{
		"manifest-path": carapace.ActionDirectories(),
	})
}
