package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var ci_getCmd = &cobra.Command{
	Use:     "get [flags]",
	Short:   "Get JSON of a running CI/CD pipeline on the current or other specified branch.",
	Aliases: []string{"stats"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ci_getCmd).Standalone()

	ci_getCmd.Flags().StringP("branch", "b", "", "Check pipeline status for a branch. (default current branch)")
	ci_getCmd.Flags().StringP("output", "F", "", "Format output. Options: text, json.")
	ci_getCmd.Flags().StringP("output-format", "o", "", "Use output.")
	ci_getCmd.Flags().StringP("pipeline-id", "p", "", "Provide pipeline ID.")
	ci_getCmd.Flags().BoolP("with-job-details", "d", false, "Show extended job information.")
	ci_getCmd.Flags().Bool("with-variables", false, "Show variables in pipeline. Requires the Maintainer role.")
	ci_getCmd.Flag("output-format").Hidden = true
	ciCmd.AddCommand(ci_getCmd)

	carapace.Gen(ci_getCmd).FlagCompletion(carapace.ActionMap{
		"branch":        action.ActionBranches(ci_getCmd),
		"output":        carapace.ActionValues("text", "json"),
		"output-format": carapace.ActionValues("text", "json"),
		"pipeline-id":   action.ActionPipelines(ci_getCmd, ""),
	})
}
