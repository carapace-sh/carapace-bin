package cmd

import (
	"fmt"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var skill_installCmd = &cobra.Command{
	Use:     "install <repository> [<skill[@version]>] [flags]",
	Short:   "Install agent skills from a GitHub repository (preview)",
	Aliases: []string{"add"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(skill_installCmd).Standalone()

	skill_installCmd.Flags().String("agent", "", "Target agent (see supported values above)")
	skill_installCmd.Flags().Bool("allow-hidden-dirs", false, "Include skills in hidden directories (e.g. .claude/skills/, .agents/skills/)")
	skill_installCmd.Flags().String("dir", "", "Install to a custom directory (overrides --agent and --scope)")
	skill_installCmd.Flags().BoolP("force", "f", false, "Overwrite existing skills without prompting")
	skill_installCmd.Flags().Bool("from-local", false, "Treat the argument as a local directory path instead of a repository")
	skill_installCmd.Flags().String("pin", "", "Pin to a specific git tag or commit SHA")
	skill_installCmd.Flags().String("scope", "", "Installation scope: {project|user}")
	skill_installCmd.Flags().Bool("upstream", false, "Install from the upstream source when a re-published skill is detected")
	skillCmd.AddCommand(skill_installCmd)

	carapace.Gen(skill_installCmd).FlagCompletion(carapace.ActionMap{
		"agent": gh.ActionAgents(),
		"dir":   carapace.ActionDirectories(),
		"pin": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) == 0 {
				return carapace.ActionValues()
			}
			return git.ActionLsRemoteRefs(git.LsRemoteRefOption{Url: fmt.Sprintf("https://github.com/%v", c.Args[0]), Tags: true})
		}),
		"scope": carapace.ActionValues("project", "user"),
	})

	carapace.Gen(skill_installCmd).PositionalCompletion(
		gh.ActionOwnerRepositories(gh.HostOpts{}),
		carapace.ActionMultiPartsN("@", 2, func(c carapace.Context) carapace.Action {
			owner, repo, ok := strings.Cut(c.Args[0], "/")
			if !ok {
				return carapace.ActionValues()
			}
			switch len(c.Parts) {
			case 0:
				return gh.ActionSkills(gh.RepoOpts{Owner: owner, Name: repo})
			default:
				return git.ActionLsRemoteRefs(git.LsRemoteRefOption{Url: fmt.Sprintf("https://github.com/%v/%v", owner, repo), Tags: true})
			}
		}),
	)
}
