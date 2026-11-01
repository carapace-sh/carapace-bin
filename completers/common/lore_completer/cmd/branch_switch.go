package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/lore_completer/cmd/action"
	"github.com/spf13/cobra"
)

var branch_switchCmd = &cobra.Command{
	Use:   "switch",
	Short: "Switch to a different branch",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_switchCmd).Standalone()

	branch_switchCmd.Flags().Bool("bare", false, "Only update anchor tracking without modifying or verifying files, useful for bare repositories")
	branch_switchCmd.Flags().Bool("dry-run", false, "Do a dry run sync and only report what changes would be done, do not change anything in the file system")
	branch_switchCmd.Flags().BoolP("help", "h", false, "Print help")
	branch_switchCmd.Flags().Bool("local", false, "Keep last local latest revision, do not sync latest revision from remote (implied by offline mode)")
	branch_switchCmd.Flags().Bool("reset", false, "Reset any local modified files to match the incoming revision")
	branchCmd.AddCommand(branch_switchCmd)

	carapace.Gen(branch_switchCmd).PositionalCompletion(
		action.ActionBranches(branch_switchCmd),
		action.ActionRevisions(branch_switchCmd),
	)
}
