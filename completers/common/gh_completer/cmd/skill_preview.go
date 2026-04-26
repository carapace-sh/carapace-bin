package cmd

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var skill_previewCmd = &cobra.Command{
	Use:     "preview <repository> [<skill>]",
	Short:   "Preview a skill from a GitHub repository (preview)",
	Aliases: []string{"show"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(skill_previewCmd).Standalone()

	skillCmd.AddCommand(skill_previewCmd)

	carapace.Gen(skill_previewCmd).PositionalCompletion(
		gh.ActionOwnerRepositories(gh.HostOpts{}),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			owner, repo, ok := strings.Cut(c.Args[0], "/")
			if !ok {
				return carapace.ActionValues()
			}
			return gh.ActionSkills(gh.RepoOpts{Owner: owner, Name: repo})
		}),
	)
}
