package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var security_group_default_statefulness_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete security group default statefulness setting(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(security_group_default_statefulness_deleteCmd).Standalone()

	security_group_default_statefulnessCmd.AddCommand(security_group_default_statefulness_deleteCmd)
}
