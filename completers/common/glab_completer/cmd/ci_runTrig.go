package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var ci_runTrigCmd = &cobra.Command{
	Use:     "run-trig [flags]",
	Short:   "Run a CI/CD pipeline trigger.",
	Aliases: []string{"run-trig"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ci_runTrigCmd).Standalone()

	ci_runTrigCmd.Flags().StringP("branch", "b", "", "Create pipeline on branch or reference <string>.")
	ci_runTrigCmd.Flags().StringSliceP("input", "i", nil, "Pass inputs to pipeline in format '<key>:<value>'. Cannot be used for merge request pipelines. See documentation for examples.")
	ci_runTrigCmd.Flags().StringP("token", "t", "", "Pipeline trigger token. Can be omitted only if the `CI_JOB_TOKEN` environment variable is set.")
	ci_runTrigCmd.Flags().StringSlice("variables", nil, "Pass variables to pipeline in the format <key>:<value>.")
	ciCmd.AddCommand(ci_runTrigCmd)

	carapace.Gen(ci_runTrigCmd).FlagCompletion(carapace.ActionMap{
		"branch": action.ActionBranches(ci_runCmd),
	})
}
