package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var release_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List releases in a repository",
	GroupID: "General commands",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(release_listCmd).Standalone()

	release_listCmd.Flags().Bool("exclude-drafts", false, "Exclude draft releases")
	release_listCmd.Flags().Bool("exclude-pre-releases", false, "Exclude pre-releases")
	release_listCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	release_listCmd.Flags().StringSlice("json", nil, "Output JSON with the specified `fields`")
	release_listCmd.Flags().StringP("limit", "L", "", "Maximum number of items to fetch")
	release_listCmd.Flags().StringP("order", "O", "", "Order of releases returned: {asc|desc}")
	release_listCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template; see \"gh help formatting\"")
	releaseCmd.AddCommand(release_listCmd)

	carapace.Gen(release_listCmd).FlagCompletion(carapace.ActionMap{
		"json":  gh.ActionReleaseFields().UniqueList(","),
		"order": carapace.ActionValues("asc", "desc").StyleF(style.ForKeyword),
	})
}
