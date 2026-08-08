package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-jq/pkg/actions/tools/jq"
	"github.com/spf13/cobra"
)

var containerRegistry_tag_viewCmd = &cobra.Command{
	Use:     "view <repository-id> <tag-name> [flags]",
	Short:   "View a container registry tag.",
	Aliases: []string{"show"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(containerRegistry_tag_viewCmd).Standalone()

	containerRegistry_tag_viewCmd.Flags().String("jq", "", "Filter JSON output with a jq expression.")
	containerRegistry_tag_viewCmd.Flags().StringP("output", "F", "", "Format output as: text, json.")
	containerRegistry_tagCmd.AddCommand(containerRegistry_tag_viewCmd)

	carapace.Gen(containerRegistry_tag_viewCmd).FlagCompletion(carapace.ActionMap{
		"jq":     jq.ActionFilters(),
		"output": carapace.ActionValues("text", "json"),
	})
}
