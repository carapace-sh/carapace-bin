package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-jq/pkg/actions/tools/jq"
	"github.com/spf13/cobra"
)

var label_getCmd = &cobra.Command{
	Use:   "get <label-id>",
	Short: "Get information about a single label by ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(label_getCmd).Standalone()

	label_getCmd.Flags().String("jq", "", "Filter JSON output with a jq expression.")
	label_getCmd.Flags().StringP("output", "F", "", "Format output as: text, json.")
	labelCmd.AddCommand(label_getCmd)

	carapace.Gen(label_getCmd).FlagCompletion(carapace.ActionMap{
		"jq":     jq.ActionFilters(),
		"output": carapace.ActionValues("text", "json"),
	})
}
