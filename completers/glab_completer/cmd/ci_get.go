package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var ci_getCmd = &cobra.Command{
	Use:     "get [flags]",
	Short:   "Get JSON of a running CI/CD pipeline on the current or other specified branch",
	Aliases: []string{"stats"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ci_getCmd).Standalone()

	ci_getCmd.Flags().StringP("branch", "b", "", "Check pipeline status for a branch. (Default is current branch)")
	ci_getCmd.Flags().StringP("output-format", "o", "", "Format output as: text, json")
	ci_getCmd.Flags().StringP("pipeline-id", "p", "", "Provide pipeline ID")
	ci_getCmd.Flags().Bool("with-variables", false, "Show variables in pipeline (maintainer role required)")
	ciCmd.AddCommand(ci_getCmd)

	carapace.Gen(ci_getCmd).FlagCompletion(carapace.ActionMap{
		"branch":        action.ActionBranches(ci_getCmd),
		"output-format": carapace.ActionValues("text", "json"),
		"pipeline-id":   action.ActionPipelines(ci_getCmd, "running"),
	})
}
