package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_testRiskCmd = &cobra.Command{
	Use:   "test-risk",
	Short: "Do a fake risky action",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_testRiskCmd).Standalone()

	debug_testRiskCmd.Flags().String("accept-risk", "", "accept risk and skip confirmation for risk types")
	debugCmd.AddCommand(debug_testRiskCmd)

	carapace.Gen(debug_testRiskCmd).FlagCompletion(carapace.ActionMap{
		"accept-risk": carapace.ActionValues("lose-ssh", "mac-app-connector", "all"),
	})
}
