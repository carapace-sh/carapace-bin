package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-jq/pkg/actions/tools/jq"
	"github.com/spf13/cobra"
)

var opentofu_state_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List states.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opentofu_state_listCmd).Standalone()

	opentofu_state_listCmd.Flags().String("jq", "", "Filter JSON output with a jq expression.")
	opentofu_state_listCmd.Flags().StringP("output", "F", "text", "Format output as: text, json.")
	opentofu_stateCmd.AddCommand(opentofu_state_listCmd)

	carapace.Gen(opentofu_state_listCmd).FlagCompletion(carapace.ActionMap{
		"jq":     jq.ActionFilters(),
		"output": carapace.ActionValues("text", "json"),
	})
}
