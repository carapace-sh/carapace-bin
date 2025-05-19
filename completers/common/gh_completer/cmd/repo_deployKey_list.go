package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var repo_deployKey_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List deploy keys in a GitHub repository",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_deployKey_listCmd).Standalone()

	repo_deployKey_listCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	repo_deployKey_listCmd.Flags().StringSlice("json", nil, "Output JSON with the specified `fields`")
	repo_deployKey_listCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template; see \"gh help formatting\"")
	repo_deployKeyCmd.AddCommand(repo_deployKey_listCmd)

	carapace.Gen(repo_deployKey_listCmd).FlagCompletion(carapace.ActionMap{
		"json": gh.ActionDeployKeyFields().UniqueList(","),
	})
}
