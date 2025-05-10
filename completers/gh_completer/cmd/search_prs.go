package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/time"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/carapace-sh/carapace-bin/pkg/styles"
	"github.com/spf13/cobra"
)

var search_prsCmd = &cobra.Command{
	Use:   "prs [<query>]",
	Short: "Search for pull requests",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(search_prsCmd).Standalone()

	search_prsCmd.Flags().String("app", "", "Filter by GitHub App author")
	search_prsCmd.Flags().Bool("archived", false, "Filter based on the repository archived state {true|false}")
	search_prsCmd.Flags().String("assignee", "", "Filter by assignee")
	search_prsCmd.Flags().String("author", "", "Filter by author")
	search_prsCmd.Flags().StringP("base", "B", "", "Filter on base branch name")
	search_prsCmd.Flags().String("checks", "", "Filter based on status of the checks: {pending|success|failure}")
	search_prsCmd.Flags().String("closed", "", "Filter on closed at `date`")
	search_prsCmd.Flags().String("commenter", "", "Filter based on comments by `user`")
	search_prsCmd.Flags().String("comments", "", "Filter on `number` of comments")
	search_prsCmd.Flags().String("created", "", "Filter based on created at `date`")
	search_prsCmd.Flags().Bool("draft", false, "Filter based on draft state")
	search_prsCmd.Flags().StringP("head", "H", "", "Filter on head branch name")
	search_prsCmd.Flags().String("interactions", "", "Filter on `number` of reactions and comments")
	search_prsCmd.Flags().String("involves", "", "Filter based on involvement of `user`")
	search_prsCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	search_prsCmd.Flags().StringSlice("json", nil, "Output JSON with the specified `fields`")
	search_prsCmd.Flags().StringSlice("label", nil, "Filter on label")
	search_prsCmd.Flags().String("language", "", "Filter based on the coding language")
	search_prsCmd.Flags().StringP("limit", "L", "", "Maximum number of results to fetch")
	search_prsCmd.Flags().Bool("locked", false, "Filter on locked conversation status")
	search_prsCmd.Flags().StringSlice("match", nil, "Restrict search to specific field of issue: {title|body|comments}")
	search_prsCmd.Flags().String("mentions", "", "Filter based on `user` mentions")
	search_prsCmd.Flags().Bool("merged", false, "Filter based on merged state")
	search_prsCmd.Flags().String("merged-at", "", "Filter on merged at `date`")
	search_prsCmd.Flags().String("milestone", "", "Filter by milestone `title`")
	search_prsCmd.Flags().Bool("no-assignee", false, "Filter on missing assignee")
	search_prsCmd.Flags().Bool("no-label", false, "Filter on missing label")
	search_prsCmd.Flags().Bool("no-milestone", false, "Filter on missing milestone")
	search_prsCmd.Flags().Bool("no-project", false, "Filter on missing project")
	search_prsCmd.Flags().String("order", "", "Order of results returned, ignored unless '--sort' flag is specified: {asc|desc}")
	search_prsCmd.Flags().StringSlice("owner", nil, "Filter on repository owner")
	search_prsCmd.Flags().String("project", "", "Filter on project board `owner/number`")
	search_prsCmd.Flags().String("reactions", "", "Filter on `number` of reactions")
	search_prsCmd.Flags().StringSliceP("repo", "R", nil, "Filter on repository")
	search_prsCmd.Flags().String("review", "", "Filter based on review status: {none|required|approved|changes_requested}")
	search_prsCmd.Flags().String("review-requested", "", "Filter on `user` or team requested to review")
	search_prsCmd.Flags().String("reviewed-by", "", "Filter on `user` who reviewed")
	search_prsCmd.Flags().String("sort", "", "Sort fetched results: {comments|reactions|reactions-+1|reactions--1|reactions-smile|reactions-thinking_face|reactions-heart|reactions-tada|interactions|created|updated}")
	search_prsCmd.Flags().String("state", "", "Filter based on state: {open|closed}")
	search_prsCmd.Flags().String("team-mentions", "", "Filter based on team mentions")
	search_prsCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template; see \"gh help formatting\"")
	search_prsCmd.Flags().String("updated", "", "Filter on last updated at `date`")
	search_prsCmd.Flags().StringSlice("visibility", nil, "Filter based on repository visibility: {public|private|internal}")
	search_prsCmd.Flags().BoolP("web", "w", false, "Open the search query in the web browser")
	searchCmd.AddCommand(search_prsCmd)

	// TODO app completion, project board number completion, eam-mentions completion
	carapace.Gen(search_prsCmd).FlagCompletion(carapace.ActionMap{
		"assignee": action.ActionSearchMultiRepo(search_prsCmd, func(cmd *cobra.Command) carapace.Action {
			return action.ActionAssignableUsers(cmd)
		}),
		"author": gh.ActionUsers(gh.HostOpts{}),
		"base": action.ActionSearchMultiRepo(search_prsCmd, func(cmd *cobra.Command) carapace.Action {
			return action.ActionBranches(cmd)
		}),
		"checks":    carapace.ActionValues("pending", "success", "failure"),
		"closed":    action.ActionSearchRange(time.ActionDateTime(time.DateTimeOpts{Strict: true})),
		"commenter": gh.ActionUsers(gh.HostOpts{}),
		"created":   action.ActionSearchRange(time.ActionDateTime(time.DateTimeOpts{Strict: true})),
		"head": action.ActionSearchMultiRepo(search_prsCmd, func(cmd *cobra.Command) carapace.Action {
			return action.ActionBranches(cmd)
		}),
		"involves": gh.ActionUsers(gh.HostOpts{}),
		"json":     action.ActionSearchIssueFields().UniqueList(","),
		"label": action.ActionSearchMultiRepo(search_prsCmd, func(cmd *cobra.Command) carapace.Action {
			return action.ActionLabels(cmd).UniqueList(",")
		}),
		"language":  gh.ActionLanguages(),
		"match":     carapace.ActionValues("title", "body", "comments").UniqueList(","),
		"mentions":  gh.ActionUsers(gh.HostOpts{}),
		"merged-at": action.ActionSearchRange(time.ActionDateTime(time.DateTimeOpts{Strict: true})),
		"milestone": action.ActionSearchMultiRepo(search_prsCmd, func(cmd *cobra.Command) carapace.Action {
			return action.ActionMilestones(cmd)
		}),
		"order": carapace.ActionValues("asc", "desc"),
		"owner": gh.ActionOwners(gh.HostOpts{}),
		"repo": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			dummyCmd := &cobra.Command{}
			dummyCmd.Flags().String("repo", c.Value, "fake repo flag")
			return action.ActionOwnerRepositories(dummyCmd).NoSpace()
		}),
		"review": carapace.ActionValues("none", "required", "approved", "changes_requested"),
		"review-requested": action.ActionSearchMultiRepo(search_prsCmd, func(cmd *cobra.Command) carapace.Action {
			return action.ActionAssignableUsers(cmd)
		}),
		"reviewed-by": action.ActionSearchMultiRepo(search_prsCmd, func(cmd *cobra.Command) carapace.Action {
			return action.ActionAssignableUsers(cmd)
		}),
		"sort":       carapace.ActionValues("comments", "created", "interactions", "reactions", "reactions-+1", "reactions--1", "reactions-heart", "reactions-smile", "reactions-tada", "reactions-thinking_face", "updated"),
		"state":      carapace.ActionValues("open", "closed").StyleF(styles.Gh.ForState),
		"updated":    action.ActionSearchRange(time.ActionDateTime(time.DateTimeOpts{Strict: true})),
		"visibility": carapace.ActionValues("public", "private", "internal").UniqueList(","),
	})
}
