package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/terraform_completer/cmd/action"
	"github.com/spf13/cobra"
)

var state_rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove instances from the state",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(state_rmCmd).Standalone()

	state_rmCmd.Flags().String("backup", "", "Path where Terraform should write the backup state.")
	state_rmCmd.Flags().Bool("dry-run", false, "Only print out what would've been removed")
	state_rmCmd.Flags().Bool("ignore-remote-version", false, "Continue even if remote and local Terraform versions are incompatible.")
	state_rmCmd.Flags().String("lock", "", "Don't hold a state lock during the operation.")
	state_rmCmd.Flags().String("lock-timeout", "", "Duration to retry a state lock.")
	state_rmCmd.Flags().String("state", "", "Path to the state file to update.")
	stateCmd.AddCommand(state_rmCmd)

	state_rmCmd.Flag("backup").NoOptDefVal = " "
	state_rmCmd.Flag("lock").NoOptDefVal = " "
	state_rmCmd.Flag("lock-timeout").NoOptDefVal = " "
	state_rmCmd.Flag("state").NoOptDefVal = " "

	carapace.Gen(state_rmCmd).FlagCompletion(carapace.ActionMap{
		"backup": carapace.ActionFiles(),
		"lock":   action.ActionBool(),
		"state":  carapace.ActionFiles(),
	})

	// TODO positional completion
}
