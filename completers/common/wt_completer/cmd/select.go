package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/wt"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var selectCmd = &cobra.Command{
	Use:    "select",
	Short:  "Deprecated: use wt switch instead",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(selectCmd).Standalone()

	selectCmd.Flags().Bool("branches", false, "Include branches without worktrees")
	selectCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	selectCmd.Flags().Bool("remotes", false, "Include remote branches")
	rootCmd.AddCommand(selectCmd)

	carapace.Gen(selectCmd).PositionalCompletion(
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
						return gh.ActionPullRequests(gh.PullRequestOpts{
							Owner: "{owner}",
							Name:  "{repo}",
						}.Default())
					// TODO case "mr":
					default:
						return carapace.ActionValues()
					}
				}
			}),
		).ToA(),
	)

	carapace.Gen(selectCmd).DashAnyCompletion(
		bridge.ActionCarapaceBin(),
	)
}
