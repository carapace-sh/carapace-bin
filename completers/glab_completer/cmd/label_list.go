package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var label_listCmd = &cobra.Command{
	Use:     "list [flags]",
	Short:   "List labels in the repository.",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(label_listCmd).Standalone()

	label_listCmd.Flags().StringP("group", "g", "", "List labels for a group.")
	label_listCmd.Flags().StringP("output", "F", "", "Format output as: text, json.")
	label_listCmd.Flags().StringP("page", "p", "", "Page number.")
	label_listCmd.Flags().StringP("per-page", "P", "", "Number of items to list per page.")
	labelCmd.AddCommand(label_listCmd)

	// TODO complete group
	carapace.Gen(label_listCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionValues("text", "json"),
	})
}
