package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var extension_searchCmd = &cobra.Command{
	Use:   "search [<query>]",
	Short: "Search extensions to the GitHub CLI",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(extension_searchCmd).Standalone()

	extension_searchCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	extension_searchCmd.Flags().StringSlice("json", nil, "Output JSON with the specified `fields`")
	extension_searchCmd.Flags().StringSlice("license", nil, "Filter based on license type")
	extension_searchCmd.Flags().StringP("limit", "L", "", "Maximum number of extensions to fetch")
	extension_searchCmd.Flags().String("order", "", "Order of repositories returned, ignored unless '--sort' flag is specified: {asc|desc}")
	extension_searchCmd.Flags().StringSlice("owner", nil, "Filter on owner")
	extension_searchCmd.Flags().String("sort", "", "Sort fetched repositories: {forks|help-wanted-issues|stars|updated}")
	extension_searchCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template; see \"gh help formatting\"")
	extension_searchCmd.Flags().BoolP("web", "w", false, "Open the search query in the web browser")
	extensionCmd.AddCommand(extension_searchCmd)

	carapace.Gen(extension_searchCmd).FlagCompletion(carapace.ActionMap{
		"json":    action.ActionSearchRepositoryFields().UniqueList(","),
		"license": gh.ActionLicenses(gh.HostOpts{}).UniqueList(","),
		"order":   carapace.ActionValues("asc", "desc").StyleF(style.ForKeyword),
		"owner":   gh.ActionOwners(gh.HostOpts{}),
		"sort":    carapace.ActionValues("forks", "help-wanted-issues", "stars", "updated"),
	})
}
