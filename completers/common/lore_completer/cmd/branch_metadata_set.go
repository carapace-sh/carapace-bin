package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/lore_completer/cmd/action"
	"github.com/spf13/cobra"
)

var branch_metadata_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set metadata on the branch",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_metadata_setCmd).Standalone()

	branch_metadata_setCmd.Flags().Bool("binary", false, "Indicator that values are paths to binary files")
	branch_metadata_setCmd.Flags().String("branch", "", "Branch name (uses current branch if not specified)")
	branch_metadata_setCmd.Flags().BoolP("help", "h", false, "Print help")
	branch_metadata_setCmd.Flags().Bool("numeric", false, "Indicator that values are numeric (u64)")
	branch_metadataCmd.AddCommand(branch_metadata_setCmd)

	carapace.Gen(branch_metadata_setCmd).FlagCompletion(carapace.ActionMap{
		"branch": action.ActionBranches(branch_metadata_setCmd),
	})
}
