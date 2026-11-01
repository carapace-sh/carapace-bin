package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var checkPlatformReqsCmd = &cobra.Command{
	Use:   "check-platform-reqs",
	Short: "Check that platform requirements are satisfied",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(checkPlatformReqsCmd).Standalone()

	checkPlatformReqsCmd.Flags().StringP("format", "f", "", "Format of the output: text or json")
	checkPlatformReqsCmd.Flags().Bool("lock", false, "Checks requirements only from the lock file, not from installed packages.")
	checkPlatformReqsCmd.Flags().Bool("no-dev", false, "Disables checking of require-dev packages requirements.")

	rootCmd.AddCommand(checkPlatformReqsCmd)

	carapace.Gen(checkPlatformReqsCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "text"),
	})
}
