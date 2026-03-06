package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/wt"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var switchCmd = &cobra.Command{
	Use:   "switch",
	Short: "Switch to a worktree; create if needed",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(switchCmd).Standalone()

	switchCmd.Flags().StringP("base", "b", "", "Base branch")
	switchCmd.Flags().Bool("branches", false, "Include branches without worktrees (interactive picker)")
	switchCmd.Flags().Bool("clobber", false, "Remove stale paths at target")
	switchCmd.Flags().BoolP("create", "c", false, "Create a new branch")
	switchCmd.Flags().StringP("execute", "x", "", "Command to run after switch")
	switchCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	switchCmd.Flags().Bool("no-cd", false, "Skip directory change after switching")
	switchCmd.Flags().Bool("no-verify", false, "Skip hooks")
	switchCmd.Flags().Bool("remotes", false, "Include remote branches (interactive picker)")
	switchCmd.Flags().BoolP("yes", "y", false, "Skip approval prompts")
	rootCmd.AddCommand(switchCmd)

	carapace.Gen(switchCmd).FlagCompletion(carapace.ActionMap{
		"base": wt.ActionBranches(),
		"execute": carapace.Batch(
			carapace.ActionExecutables(),
			carapace.ActionFiles(),
		).ToA(),
	})

	carapace.Gen(switchCmd).PositionalCompletion(
		carapace.Batch(
			carapace.ActionValuesDescribed(
				"^", "default branch",
				"@", "current branch/worktree",
				"-", "previous worktree",
			),
			wt.ActionWorktrees(),
			carapace.ActionMultiPartsN(":", 2, func(c carapace.Context) carapace.Action {
				switch len(c.Parts) {
				case 0:
					return carapace.ActionValuesDescribed(
						"pr", "pull requests",
						"mr", "merge requests",
					).Suffix(":")
				default:
					switch c.Parts[0] {
					case "pr":
						return gh.ActionPullRequests(gh.PullRequestOpts{}.Default())
					// TODO case "mr":
					default:
						return carapace.ActionValues()
					}
				}
			}),
		).ToA(),
	)

	carapace.Gen(switchCmd).DashAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return bridge.ActionCarapaceBin(switchCmd.Flag("execute").Value.String())
		}),
	)
}
