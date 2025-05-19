package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var showBranchCmd = &cobra.Command{
	Use:     "show-branch",
	Short:   "Show branches and their commits",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_interrogator].ID,
}

func init() {
	carapace.Gen(showBranchCmd).Standalone()

	showBranchCmd.Flags().BoolP("all", "a", false, "show remote-tracking and local branches")
	showBranchCmd.Flags().String("color", "", "color '*!+-' corresponding to the branch")
	showBranchCmd.Flags().Bool("current", false, "include the current branch")
	showBranchCmd.Flags().Bool("date-order", false, "topologically sort, maintaining date order where possible")
	showBranchCmd.Flags().Bool("independent", false, "show refs unreachable from any other ref")
	showBranchCmd.Flags().Bool("list", false, "synonym to more=-1")
	showBranchCmd.Flags().Bool("merge-base", false, "show possible merge bases")
	showBranchCmd.Flags().String("more", "", "show <n> more commits after the common ancestor")
	showBranchCmd.Flags().Bool("no-name", false, "suppress naming strings")
	showBranchCmd.Flags().StringP("reflog", "g", "", "show <n> most recent ref-log entries starting at base")
	showBranchCmd.Flags().BoolP("remotes", "r", false, "show remote-tracking branches")
	showBranchCmd.Flags().Bool("sha1-name", false, "name commits with their object names")
	showBranchCmd.Flags().Bool("sparse", false, "show merges reachable from only one tip")
	showBranchCmd.Flags().Bool("topics", false, "show only commits not on the first branch")
	showBranchCmd.Flags().Bool("topo-order", false, "show commits in topological order")
	rootCmd.AddCommand(showBranchCmd)

	showBranchCmd.Flag("color").NoOptDefVal = " "
	showBranchCmd.Flag("more").NoOptDefVal = " "
	showBranchCmd.Flag("reflog").NoOptDefVal = " "

	carapace.Gen(showBranchCmd).FlagCompletion(carapace.ActionMap{
		"color": git.ActionColorModes(),
		"reflog": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 1:
				return git.ActionRefs(git.RefOption{}.Default())
			default:
				return carapace.ActionValues()
			}
		}),
	})

	carapace.Gen(showBranchCmd).PositionalCompletion(
		git.ActionRefs(git.RefOption{}.Default()),
	)

}
