package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/carapace-sh/carapace-bin/pkg/styles"
	"github.com/carapace-sh/carapace-jq/pkg/actions/tools/jq"
	"github.com/spf13/cobra"
)

var discussion_listCmd = &cobra.Command{
	Use:     "list [flags]",
	Short:   "List discussions in a repository (preview)",
	GroupID: "General commands",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(discussion_listCmd).Standalone()

	discussion_listCmd.Flags().String("after", "", "Cursor for the next page of results")
	discussion_listCmd.Flags().Bool("answered", false, "Filter by answered state")
	discussion_listCmd.Flags().StringP("author", "A", "", "Filter by author")
	discussion_listCmd.Flags().StringP("category", "c", "", "Filter by category name or slug")
	discussion_listCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	discussion_listCmd.Flags().StringSlice("json", nil, "Output JSON with the specified `fields`")
	discussion_listCmd.Flags().StringSliceP("label", "l", nil, "Filter by label")
	discussion_listCmd.Flags().StringP("limit", "L", "", "Maximum number of discussions to fetch")
	discussion_listCmd.Flags().String("order", "", "Order of results: {asc|desc}")
	discussion_listCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository using the `[HOST/]OWNER/REPO` format")
	discussion_listCmd.Flags().StringP("search", "S", "", "Search discussions with `query`")
	discussion_listCmd.Flags().String("sort", "", "Sort by field: {created|updated}")
	discussion_listCmd.Flags().StringP("state", "s", "", "Filter by state: {open|closed|all}")
	discussion_listCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template; see \"gh help formatting\"")
	discussion_listCmd.Flags().BoolP("web", "w", false, "List discussions in the web browser")
	discussionCmd.AddCommand(discussion_listCmd)

	carapace.Gen(discussion_listCmd).FlagCompletion(carapace.ActionMap{
		"author":   gh.ActionUsers(gh.HostOpts{}),
		"category": action.ActionDiscussionCategories(discussion_listCmd),
		"jq":       jq.ActionFilters(),
		"json":     gh.ActionDiscussionListFields().UniqueList(","),
		"label":    action.ActionLabels(discussion_listCmd).UniqueList(","),
		"order":    carapace.ActionValues("asc", "desc"),
		"repo":     gh.ActionHostOwnerRepositories(),
		"sort":     carapace.ActionValues("created", "updated"),
		"state":    carapace.ActionValues("open", "closed", "all").StyleF(styles.Gh.ForState),
	})
}
