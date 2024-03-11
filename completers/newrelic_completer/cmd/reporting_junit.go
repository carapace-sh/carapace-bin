package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var reporting_junitCmd = &cobra.Command{
	Use:   "junit",
	Short: "Send JUnit test run results to New Relic",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reporting_junitCmd).Standalone()
	reporting_junitCmd.Flags().Bool("dryRun", false, "suppress posting custom events to NRDB")
	reporting_junitCmd.Flags().BoolP("output", "o", false, "output generated custom events to stdout")
	reporting_junitCmd.Flags().StringP("path", "p", "", "the path to a JUnit-formatted test results file")
	reportingCmd.AddCommand(reporting_junitCmd)

	carapace.Gen(reporting_junitCmd).FlagCompletion(carapace.ActionMap{
		"path": carapace.ActionFiles(),
	})
}
