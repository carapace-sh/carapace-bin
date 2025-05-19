package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/time"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/carapace-sh/carapace-bin/pkg/styles"
	"github.com/spf13/cobra"
)

var search_issuesCmd = &cobra.Command{
	Use:   "issues [<query>]",
	Short: "Search for issues",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(search_issuesCmd).Standalone()

	search_issuesCmd.Flags().String("app", "", "Filter by GitHub App author")
	search_issuesCmd.Flags().Bool("archived", false, "Filter based on the repository archived state {true|false}")
	search_issuesCmd.Flags().String("assignee", "", "Filter by assignee")
	search_issuesCmd.Flags().String("author", "", "Filter by author")
	search_issuesCmd.Flags().String("closed", "", "Filter on closed at `date`")
	search_issuesCmd.Flags().String("commenter", "", "Filter based on comments by `user`")
	search_issuesCmd.Flags().String("comments", "", "Filter on `number` of comments")
	search_issuesCmd.Flags().String("created", "", "Filter based on created at `date`")
	search_issuesCmd.Flags().Bool("include-prs", false, "Include pull requests in results")
	search_issuesCmd.Flags().String("interactions", "", "Filter on `number` of reactions and comments")
	search_issuesCmd.Flags().String("involves", "", "Filter based on involvement of `user`")
	search_issuesCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	search_issuesCmd.Flags().StringSlice("json", nil, "Output JSON with the specified `fields`")
	search_issuesCmd.Flags().StringSlice("label", nil, "Filter on label")
	search_issuesCmd.Flags().String("language", "", "Filter based on the coding language")
	search_issuesCmd.Flags().StringP("limit", "L", "", "Maximum number of results to fetch")
	search_issuesCmd.Flags().Bool("locked", false, "Filter on locked conversation status")
	search_issuesCmd.Flags().StringSlice("match", nil, "Restrict search to specific field of issue: {title|body|comments}")
	search_issuesCmd.Flags().String("mentions", "", "Filter based on `user` mentions")
	search_issuesCmd.Flags().String("milestone", "", "Filter by milestone `title`")
	search_issuesCmd.Flags().Bool("no-assignee", false, "Filter on missing assignee")
	search_issuesCmd.Flags().Bool("no-label", false, "Filter on missing label")
	search_issuesCmd.Flags().Bool("no-milestone", false, "Filter on missing milestone")
	search_issuesCmd.Flags().Bool("no-project", false, "Filter on missing project")
	search_issuesCmd.Flags().String("order", "", "Order of results returned, ignored unless '--sort' flag is specified: {asc|desc}")
	search_issuesCmd.Flags().StringSlice("owner", nil, "Filter on repository owner")
	search_issuesCmd.Flags().String("project", "", "Filter on project board `owner/number`")
	search_issuesCmd.Flags().String("reactions", "", "Filter on `number` of reactions")
	search_issuesCmd.Flags().StringSliceP("repo", "R", nil, "Filter on repository")
	search_issuesCmd.Flags().String("sort", "", "Sort fetched results: {comments|created|interactions|reactions|reactions-+1|reactions--1|reactions-heart|reactions-smile|reactions-tada|reactions-thinking_face|updated}")
	search_issuesCmd.Flags().String("state", "", "Filter based on state: {open|closed}")
	search_issuesCmd.Flags().String("team-mentions", "", "Filter based on team mentions")
	search_issuesCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template; see \"gh help formatting\"")
	search_issuesCmd.Flags().String("updated", "", "Filter on last updated at `date`")
	search_issuesCmd.Flags().StringSlice("visibility", nil, "Filter based on repository visibility: {public|private|internal}")
	search_issuesCmd.Flags().BoolP("web", "w", false, "Open the search query in the web browser")
	searchCmd.AddCommand(search_issuesCmd)

	// TODO app completion, project board number completion
	carapace.Gen(search_issuesCmd).FlagCompletion(carapace.ActionMap{
		"assignee": action.ActionSearchMultiRepo(search_issuesCmd, func(cmd *cobra.Command) carapace.Action {
			return action.ActionAssignableUsers(cmd)
		}),
		"author":    gh.ActionUsers(gh.HostOpts{}),
		"closed":    action.ActionSearchRange(time.ActionDateTime(time.DateTimeOpts{Strict: true})),
		"commenter": gh.ActionUsers(gh.HostOpts{}),
		"created":   action.ActionSearchRange(time.ActionDateTime(time.DateTimeOpts{Strict: true})),
		"involves":  gh.ActionUsers(gh.HostOpts{}),
		"json":      action.ActionSearchIssueFields().UniqueList(","),
		"label": action.ActionSearchMultiRepo(search_issuesCmd, func(cmd *cobra.Command) carapace.Action {
			return action.ActionLabels(cmd).UniqueList(",")
		}),
		"language": gh.ActionLanguages(),
		"match":    carapace.ActionValues("title", "body", "comments").UniqueList(","),
		"mentions": gh.ActionUsers(gh.HostOpts{}),
		"milestone": action.ActionSearchMultiRepo(search_issuesCmd, func(cmd *cobra.Command) carapace.Action {
			return action.ActionMilestones(cmd)
		}),
		"order": carapace.ActionValues("asc", "desc"),
		"owner": gh.ActionOwners(gh.HostOpts{}),
		"repo": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			dummyCmd := &cobra.Command{}
			dummyCmd.Flags().String("repo", c.Value, "fake repo flag")
			return action.ActionOwnerRepositories(dummyCmd).NoSpace()
		}),
		"sort":       carapace.ActionValues("comments", "created", "interactions", "reactions", "reactions-+1", "reactions--1", "reactions-heart", "reactions-smile", "reactions-tada", "reactions-thinking_face", "updated"),
		"state":      carapace.ActionValues("open", "closed").StyleF(styles.Gh.ForState),
		"updated":    action.ActionSearchRange(time.ActionDateTime(time.DateTimeOpts{Strict: true})),
		"visibility": carapace.ActionValues("public", "private", "internal").UniqueList(","),
	})
}
