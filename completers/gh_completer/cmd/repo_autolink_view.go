package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var repo_autolink_viewCmd = &cobra.Command{
	Use:   "view <id>",
	Short: "View an autolink reference",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_autolink_viewCmd).Standalone()

	repo_autolink_viewCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	repo_autolink_viewCmd.Flags().StringSlice("json", nil, "Output JSON with the specified `fields`")
	repo_autolink_viewCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template; see \"gh help formatting\"")
	repo_autolinkCmd.AddCommand(repo_autolink_viewCmd)

	carapace.Gen(repo_autolink_viewCmd).FlagCompletion(carapace.ActionMap{
		"json": gh.ActionAutolinkFields().UniqueList(","),
	})

	carapace.Gen(repo_autolink_viewCmd).PositionalCompletion(
		action.ActionAutolinks(repo_autolink_viewCmd),
	)
}
