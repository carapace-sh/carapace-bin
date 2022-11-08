package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var extension_searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search extensions to the GitHub CLI",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(extension_searchCmd).Standalone()
	extension_searchCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	extension_searchCmd.Flags().StringSlice("json", []string{}, "Output JSON with the specified `fields`")
	extension_searchCmd.Flags().StringSlice("license", []string{}, "Filter based on license type")
	extension_searchCmd.Flags().IntP("limit", "L", 30, "Maximum number of extensions to fetch")
	extension_searchCmd.Flags().String("order", "desc", "Order of repositories returned, ignored unless '--sort' flag is specified: {asc|desc}")
	extension_searchCmd.Flags().String("owner", "", "Filter on owner")
	extension_searchCmd.Flags().String("sort", "best-match", "Sort fetched repositories: {forks|help-wanted-issues|stars|updated}")
	extension_searchCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template; see \"gh help formatting\"")
	extension_searchCmd.Flags().BoolP("web", "w", false, "Open the search query in the web browser")
	extensionCmd.AddCommand(extension_searchCmd)

	carapace.Gen(extension_searchCmd).FlagCompletion(carapace.ActionMap{
		"json": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionSearchRepositoryFields().Invoke(c).Filter(c.Parts).ToA()
		}),
		"license": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return gh.ActionLicenses(gh.HostOpts{}).Invoke(c).Filter(c.Parts).ToA()
		}),
		"order": carapace.ActionValues("asc", "desc"),
		"owner": gh.ActionOwners(gh.HostOpts{}),
		"sort":  carapace.ActionValues("forks", "help-wanted-issues", "stars", "updated"),
	})
}
