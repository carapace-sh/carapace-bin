package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var ci_getCmd = &cobra.Command{
	Use:     "get [flags]",
	Short:   "Get the details of a CI/CD pipeline.",
	Aliases: []string{"stats"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ci_getCmd).Standalone()

	ci_getCmd.Flags().StringP("branch", "b", "", "Get the pipeline for a branch. Defaults to the current branch.")
	ci_getCmd.Flags().String("jq", "", "Filter JSON output with a jq expression.")
	ci_getCmd.Flags().String("merge-request", "", "Show the pipeline for the given merge request <iid>.")
	ci_getCmd.Flags().StringP("output", "F", "", "Format output. Options: text, json.")
	ci_getCmd.Flags().StringP("output-format", "o", "", "Use output.")
	ci_getCmd.Flags().StringP("pipeline-id", "p", "", "Get the pipeline with the given <id>.")
	ci_getCmd.Flags().StringP("status", "s", "", "Show only jobs in the given state. Passed through to the API's scope parameter.")
	ci_getCmd.Flags().BoolP("with-job-details", "d", false, "Show extended job information.")
	ci_getCmd.Flags().Bool("with-variables", false, "Show variables in pipeline. Requires the Maintainer role.")
	ci_getCmd.Flag("output-format").Hidden = true
	ci_getCmd.Flag("output-format").Hidden = true
	ciCmd.AddCommand(ci_getCmd)

	carapace.Gen(ci_getCmd).FlagCompletion(carapace.ActionMap{
		"branch":        action.ActionBranches(ci_getCmd),
		"output":        carapace.ActionValues("text", "json"),
		"output-format": carapace.ActionValues("text", "json"),
		"pipeline-id":   action.ActionPipelines(ci_getCmd, ""),
	})
}
