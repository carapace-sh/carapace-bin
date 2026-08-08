package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-jq/pkg/actions/tools/jq"
	"github.com/spf13/cobra"
)

var deployKey_listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get a list of deploy keys for the current project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deployKey_listCmd).Standalone()

	deployKey_listCmd.Flags().String("jq", "", "Filter JSON output with a jq expression.")
	deployKey_listCmd.Flags().StringP("output", "F", "", "Format output as: text, json.")
	deployKey_listCmd.Flags().StringP("page", "p", "", "Page number.")
	deployKey_listCmd.Flags().StringP("per-page", "P", "", "Number of items to list per page.")
	deployKey_listCmd.Flags().Bool("show-id", false, "Shows IDs of deploy keys.")
	deployKeyCmd.AddCommand(deployKey_listCmd)

	carapace.Gen(deployKey_listCmd).FlagCompletion(carapace.ActionMap{
		"jq":     jq.ActionFilters(),
		"output": carapace.ActionValues("text", "json"),
	})
}
