package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var repo_autolink_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List autolink references for a GitHub repository",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_autolink_listCmd).Standalone()

	repo_autolink_listCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	repo_autolink_listCmd.Flags().StringSlice("json", []string{}, "Output JSON with the specified `fields`")
	repo_autolink_listCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template; see \"gh help formatting\"")
	repo_autolink_listCmd.Flags().BoolP("web", "w", false, "List autolink references in the web browser")
	repo_autolinkCmd.AddCommand(repo_autolink_listCmd)

	carapace.Gen(repo_autolink_listCmd).FlagCompletion(carapace.ActionMap{
		"json": gh.ActionAutolinkFields().UniqueList(","),
	})
}
