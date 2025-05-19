package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var search_commitsCmd = &cobra.Command{
	Use:   "commits [<query>]",
	Short: "Search for commits",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(search_commitsCmd).Standalone()

	search_commitsCmd.Flags().String("author", "", "Filter by author")
	search_commitsCmd.Flags().String("author-date", "", "Filter based on authored `date`")
	search_commitsCmd.Flags().String("author-email", "", "Filter on author email")
	search_commitsCmd.Flags().String("author-name", "", "Filter on author name")
	search_commitsCmd.Flags().String("committer", "", "Filter by committer")
	search_commitsCmd.Flags().String("committer-date", "", "Filter based on committed `date`")
	search_commitsCmd.Flags().String("committer-email", "", "Filter on committer email")
	search_commitsCmd.Flags().String("committer-name", "", "Filter on committer name")
	search_commitsCmd.Flags().String("hash", "", "Filter by commit hash")
	search_commitsCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	search_commitsCmd.Flags().StringSlice("json", nil, "Output JSON with the specified `fields`")
	search_commitsCmd.Flags().StringP("limit", "L", "", "Maximum number of commits to fetch")
	search_commitsCmd.Flags().Bool("merge", false, "Filter on merge commits")
	search_commitsCmd.Flags().String("order", "", "Order of commits returned, ignored unless '--sort' flag is specified: {asc|desc}")
	search_commitsCmd.Flags().StringSlice("owner", nil, "Filter on repository owner")
	search_commitsCmd.Flags().String("parent", "", "Filter by parent hash")
	search_commitsCmd.Flags().StringSliceP("repo", "R", nil, "Filter on repository")
	search_commitsCmd.Flags().String("sort", "", "Sort fetched commits: {author-date|committer-date}")
	search_commitsCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template; see \"gh help formatting\"")
	search_commitsCmd.Flags().String("tree", "", "Filter by tree hash")
	search_commitsCmd.Flags().StringSlice("visibility", nil, "Filter based on repository visibility: {public|private|internal}")
	search_commitsCmd.Flags().BoolP("web", "w", false, "Open the search query in the web browser")
	searchCmd.AddCommand(search_commitsCmd)

	carapace.Gen(search_commitsCmd).FlagCompletion(carapace.ActionMap{
		"author":         gh.ActionUsers(gh.HostOpts{}),
		"author-date":    gh.ActionDateFields(),
		"committer":      gh.ActionUsers(gh.HostOpts{}),
		"committer-date": gh.ActionDateFields(),
		"json":           gh.ActionCommitFields().UniqueList(","),
		"order":          carapace.ActionValues("asc", "desc"),
		"owner":          gh.ActionOwners(gh.HostOpts{}),
		"repo": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			dummyCmd := &cobra.Command{}
			dummyCmd.Flags().String("repo", c.Value, "fake repo flag")
			return action.ActionOwnerRepositories(dummyCmd).NoSpace()
		}),
		"sort":       carapace.ActionValues("author-date", "committer-date"),
		"visibility": carapace.ActionValues("public", "private", "internal").UniqueList(","),
	})
}
