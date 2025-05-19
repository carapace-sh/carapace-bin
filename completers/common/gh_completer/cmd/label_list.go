package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var label_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List labels in a repository",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(label_listCmd).Standalone()

	label_listCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	label_listCmd.Flags().StringSlice("json", nil, "Output JSON with the specified `fields`")
	label_listCmd.Flags().StringP("limit", "L", "", "Maximum number of labels to fetch")
	label_listCmd.Flags().String("order", "", "Order of labels returned: {asc|desc}")
	label_listCmd.Flags().StringP("search", "S", "", "Search label names and descriptions")
	label_listCmd.Flags().String("sort", "", "Sort fetched labels: {created|name}")
	label_listCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template; see \"gh help formatting\"")
	label_listCmd.Flags().BoolP("web", "w", false, "List labels in the web browser")
	labelCmd.AddCommand(label_listCmd)

	carapace.Gen(label_listCmd).FlagCompletion(carapace.ActionMap{
		"json":  gh.ActionLabelFields().UniqueList(","),
		"order": carapace.ActionValues("asc", "desc"),
		"sort":  carapace.ActionValues("created", "name"),
	})
}
