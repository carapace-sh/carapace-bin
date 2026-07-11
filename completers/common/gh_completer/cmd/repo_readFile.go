package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/carapace-sh/carapace-jq/pkg/actions/tools/jq"
	"github.com/spf13/cobra"
)

var repo_readFileCmd = &cobra.Command{
	Use:     "read-file <path> [flags]",
	Short:   "Read a file from a repository (preview)",
	GroupID: "Targeted commands",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_readFileCmd).Standalone()

	repo_readFileCmd.Flags().Bool("allow-escape-sequences", false, "Allow printing terminal escape sequences")
	repo_readFileCmd.Flags().Bool("clobber", false, "Overwrite the output path if it already exists")
	repo_readFileCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	repo_readFileCmd.Flags().StringSlice("json", nil, "Output JSON with the specified `fields`")
	repo_readFileCmd.Flags().StringP("output", "o", "", "Write the file to a `path` instead of stdout")
	repo_readFileCmd.Flags().String("ref", "", "The branch, tag, or commit to read from")
	repo_readFileCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository using the `[HOST/]OWNER/REPO` format")
	repo_readFileCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template; see \"gh help formatting\"")
	repoCmd.AddCommand(repo_readFileCmd)

	carapace.Gen(repo_readFileCmd).FlagCompletion(carapace.ActionMap{
		"jq":     jq.ActionFilters(),
		"json":   gh.ActionReadFileFields().UniqueList(","),
		"output": carapace.ActionFiles(),
		"ref":    action.ActionBranches(repo_readFileCmd),
		"repo":   gh.ActionHostOwnerRepositories(),
	})

	carapace.Gen(repo_readFileCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionContents(repo_readFileCmd, "", repo_readFileCmd.Flag("ref").Value.String())
		}),
	)
}
