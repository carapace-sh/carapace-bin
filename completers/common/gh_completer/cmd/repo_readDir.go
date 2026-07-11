package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/carapace-sh/carapace-jq/pkg/actions/tools/jq"
	"github.com/spf13/cobra"
)

var repo_readDirCmd = &cobra.Command{
	Use:     "read-dir [<path>] [flags]",
	Short:   "List a directory in a repository (preview)",
	GroupID: "Targeted commands",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_readDirCmd).Standalone()

	repo_readDirCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	repo_readDirCmd.Flags().StringSlice("json", nil, "Output JSON with the specified `fields`")
	repo_readDirCmd.Flags().String("ref", "", "The branch, tag, or commit to list from")
	repo_readDirCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository using the `[HOST/]OWNER/REPO` format")
	repo_readDirCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template; see \"gh help formatting\"")
	repoCmd.AddCommand(repo_readDirCmd)

	carapace.Gen(repo_readDirCmd).FlagCompletion(carapace.ActionMap{
		"jq":   jq.ActionFilters(),
		"json": gh.ActionReadDirFields().UniqueList(","),
		"ref":  action.ActionBranches(repo_readDirCmd),
		"repo": gh.ActionHostOwnerRepositories(),
	})

	carapace.Gen(repo_readDirCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionContents(repo_readDirCmd, "", repo_readDirCmd.Flag("ref").Value.String())
		}),
	)
}
