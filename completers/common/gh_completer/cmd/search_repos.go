package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/time"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var search_reposCmd = &cobra.Command{
	Use:   "repos [<query>]",
	Short: "Search for repositories",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(search_reposCmd).Standalone()

	search_reposCmd.Flags().Bool("archived", false, "Filter based on the repository archived state {true|false}")
	search_reposCmd.Flags().String("created", "", "Filter based on created at `date`")
	search_reposCmd.Flags().String("followers", "", "Filter based on `number` of followers")
	search_reposCmd.Flags().String("forks", "", "Filter on `number` of forks")
	search_reposCmd.Flags().String("good-first-issues", "", "Filter on `number` of issues with the 'good first issue' label")
	search_reposCmd.Flags().String("help-wanted-issues", "", "Filter on `number` of issues with the 'help wanted' label")
	search_reposCmd.Flags().String("include-forks", "", "Include forks in fetched repositories: {false|true|only}")
	search_reposCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	search_reposCmd.Flags().StringSlice("json", nil, "Output JSON with the specified `fields`")
	search_reposCmd.Flags().String("language", "", "Filter based on the coding language")
	search_reposCmd.Flags().StringSlice("license", nil, "Filter based on license type")
	search_reposCmd.Flags().StringP("limit", "L", "", "Maximum number of repositories to fetch")
	search_reposCmd.Flags().StringSlice("match", nil, "Restrict search to specific field of repository: {name|description|readme}")
	search_reposCmd.Flags().String("number-topics", "", "Filter on `number` of topics")
	search_reposCmd.Flags().String("order", "", "Order of repositories returned, ignored unless '--sort' flag is specified: {asc|desc}")
	search_reposCmd.Flags().StringSlice("owner", nil, "Filter on owner")
	search_reposCmd.Flags().String("size", "", "Filter on a size range, in kilobytes")
	search_reposCmd.Flags().String("sort", "", "Sort fetched repositories: {forks|help-wanted-issues|stars|updated}")
	search_reposCmd.Flags().String("stars", "", "Filter on `number` of stars")
	search_reposCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template; see \"gh help formatting\"")
	search_reposCmd.Flags().StringSlice("topic", nil, "Filter on topic")
	search_reposCmd.Flags().String("updated", "", "Filter on last updated at `date`")
	search_reposCmd.Flags().StringSlice("visibility", nil, "Filter based on visibility: {public|private|internal}")
	search_reposCmd.Flags().BoolP("web", "w", false, "Open the search query in the web browser")
	searchCmd.AddCommand(search_reposCmd)

	carapace.Gen(search_reposCmd).FlagCompletion(carapace.ActionMap{
		"created":       action.ActionSearchRange(time.ActionDateTime(time.DateTimeOpts{Strict: true})),
		"include-forks": carapace.ActionValues("false", "true", "only"),
		"json":          action.ActionSearchRepositoryFields().UniqueList(","),
		"language":      gh.ActionLanguages(),
		"license":       gh.ActionLicenses(gh.HostOpts{}).UniqueList(","),
		"match":         carapace.ActionValues("name", "description", "readme").UniqueList(","),
		"order":         carapace.ActionValues("asc", "desc"),
		"owner":         gh.ActionOwners(gh.HostOpts{}),
		"sort":          carapace.ActionValues("forks", "help-wanted-issues", "stars", "updated"),
		"topic":         action.ActionTopicSearch(search_reposCmd).UniqueList(","),
		"updated":       action.ActionSearchRange(time.ActionDateTime(time.DateTimeOpts{Strict: true})),
		"visibility":    carapace.ActionValues("public", "private", "internal").UniqueList(","),
	})
}
