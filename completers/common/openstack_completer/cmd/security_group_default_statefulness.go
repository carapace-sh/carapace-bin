package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var security_group_default_statefulnessCmd = &cobra.Command{
	Use:   "statefulness",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(security_group_default_statefulnessCmd).Standalone()

	security_group_defaultCmd.AddCommand(security_group_default_statefulnessCmd)
}
