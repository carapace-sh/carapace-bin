package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/newrelic"
	"github.com/spf13/cobra"
)

var diagnose_runCmd = &cobra.Command{
	Use:   "run",
	Short: "Troubleshoot your New Relic-instrumented application",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(diagnose_runCmd).Standalone()
	diagnose_runCmd.Flags().String("attachment-key", "", "Attachment key for automatic upload to a support ticket (get key from an existing ticket).")
	diagnose_runCmd.Flags().Bool("list-suites", false, "List the task suites available for the --suites argument.")
	diagnose_runCmd.Flags().String("suites", "", "The task suite or comma-separated list of suites to run. Use --list-suites for a list of available suites.")
	diagnose_runCmd.Flags().Bool("verbose", false, "Display verbose logging during task execution.")
	diagnoseCmd.AddCommand(diagnose_runCmd)

	carapace.Gen(diagnose_runCmd).FlagCompletion(carapace.ActionMap{
		"suites": newrelic.ActionDiagnosticSuites().UniqueList(","),
	})
}
