package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var iteration_listCmd = &cobra.Command{
	Use:     "list [flags]",
	Short:   "List project iterations",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iteration_listCmd).Standalone()

	iteration_listCmd.Flags().StringP("group", "g", "", "List iterations for a group.")
	iteration_listCmd.Flags().StringP("output", "F", "", "Format output as: text, json.")
	iteration_listCmd.Flags().StringP("page", "p", "", "Page number.")
	iteration_listCmd.Flags().StringP("per-page", "P", "", "Number of items to list per page.")
	iterationCmd.AddCommand(iteration_listCmd)

	carapace.Gen(iteration_listCmd).FlagCompletion(carapace.ActionMap{
		"group":  action.ActionGroups(iteration_listCmd),
		"output": carapace.ActionValues("text", "json"),
	})
}
