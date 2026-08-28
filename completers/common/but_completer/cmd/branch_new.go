package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/but"
	"github.com/spf13/cobra"
)

var branch_newCmd = &cobra.Command{
	Use:   "new",
	Short: "Creates a new branch in the workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_newCmd).Standalone()

	branch_newCmd.Flags().StringP("above", "A", "", "Place the branch above BRANCH_OR_COMMIT, which must be an applied branch or commit")
	branch_newCmd.Flags().Bool("allow-merged", false, "Allow targeting branches and commits that are already merged upstream")
	branch_newCmd.Flags().StringP("anchor", "a", "", "Deprecated flag that will be removed in a future release. Use --above instead")
	branch_newCmd.Flags().StringP("below", "B", "", "Place the branch below BRANCH_OR_COMMIT, which must be an applied branch or commit")
	branch_newCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	branch_newCmd.Flags().BoolP("switch", "s", false, "Switch to the newly created branch instead of applying it to the GitButler workspace")
	branch_newCmd.Flag("anchor").Hidden = true
	branchCmd.AddCommand(branch_newCmd)

	carapace.Gen(branch_newCmd).FlagCompletion(carapace.ActionMap{
		"anchor": but.ActionTargets(),
	})

	carapace.Gen(branch_newCmd).PositionalCompletion(
		but.ActionLocalBranches(),
	)
}
