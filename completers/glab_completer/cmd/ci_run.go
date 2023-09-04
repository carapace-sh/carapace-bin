package cmd

import (
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var ci_runCmd = &cobra.Command{
	Use:     "run [flags]",
	Short:   "Create or run a new CI/CD pipeline",
	Aliases: []string{"create"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ci_runCmd).Standalone()

	ci_runCmd.Flags().StringP("branch", "b", "", "Create pipeline on branch/ref <string>")
	ci_runCmd.Flags().StringSlice("variables", []string{}, "Pass variables to pipeline in format <key>:<value>")
	ci_runCmd.Flags().StringSlice("variables-env", []string{}, "Pass variables to pipeline in format <key>:<value>")
	ci_runCmd.Flags().StringSlice("variables-file", []string{}, "Pass file contents as a file variable to pipeline in format <key>:<filename>")
	ci_runCmd.Flags().StringP("variables-from", "f", "", "JSON file containing variables for pipeline execution")
	ciCmd.AddCommand(ci_runCmd)

	carapace.Gen(ci_runCmd).FlagCompletion(carapace.ActionMap{
		"branch": action.ActionBranches(ci_runCmd),
		"variables-file": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			if splitted := strings.SplitN(c.Value, ":", 2); len(splitted) == 2 {
				c.Value = splitted[1]
				return carapace.ActionFiles().Invoke(c).Prefix(splitted[0] + ":").ToA().NoSpace()
			}
			return carapace.ActionValues().NoSpace()
		}),
	})
}
