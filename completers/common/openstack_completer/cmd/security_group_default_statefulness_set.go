package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var security_group_default_statefulness_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Update a security group default statefulness setting",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(security_group_default_statefulness_setCmd).Standalone()

	security_group_default_statefulness_setCmd.Flags().Bool("stateful", false, "Set default statefulness to stateful")
	security_group_default_statefulness_setCmd.Flags().Bool("stateless", false, "Set default statefulness to stateless")
	security_group_default_statefulnessCmd.AddCommand(security_group_default_statefulness_setCmd)
}
