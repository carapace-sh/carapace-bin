package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-jq/pkg/actions/tools/jq"
	"github.com/spf13/cobra"
)

var containerRegistry_tag_listCmd = &cobra.Command{
	Use:     "list <repository-id> [flags]",
	Short:   "List container registry repository tags.",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(containerRegistry_tag_listCmd).Standalone()

	containerRegistry_tag_listCmd.Flags().Bool("details", false, "Fetch digest, size, and creation time for each tag. Makes one API call per tag.")
	containerRegistry_tag_listCmd.Flags().String("jq", "", "Filter JSON output with a jq expression.")
	containerRegistry_tag_listCmd.Flags().StringP("output", "F", "", "Format output as: text, json.")
	containerRegistry_tag_listCmd.Flags().StringP("page", "p", "", "Page number.")
	containerRegistry_tag_listCmd.Flags().StringP("per-page", "P", "", "Number of items to list per page.")
	containerRegistry_tagCmd.AddCommand(containerRegistry_tag_listCmd)

	carapace.Gen(containerRegistry_tag_listCmd).FlagCompletion(carapace.ActionMap{
		"jq":     jq.ActionFilters(),
		"output": carapace.ActionValues("text", "json"),
	})
}
