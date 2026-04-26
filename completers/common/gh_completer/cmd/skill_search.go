package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var skill_searchCmd = &cobra.Command{
	Use:   "search <query> [flags]",
	Short: "Search for skills across GitHub (preview)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(skill_searchCmd).Standalone()

	skill_searchCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	skill_searchCmd.Flags().StringSlice("json", nil, "Output JSON with the specified `fields`")
	skill_searchCmd.Flags().StringP("limit", "L", "", "Maximum number of results per page")
	skill_searchCmd.Flags().String("owner", "", "Filter results to a specific GitHub user or organization")
	skill_searchCmd.Flags().String("page", "", "Page number of results to fetch")
	skill_searchCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template; see \"gh help formatting\"")
	skillCmd.AddCommand(skill_searchCmd)

	carapace.Gen(skill_searchCmd).FlagCompletion(carapace.ActionMap{
		"json":  carapace.ActionValues(), // TODO
		"owner": gh.ActionOwners(gh.HostOpts{}),
	})
}
