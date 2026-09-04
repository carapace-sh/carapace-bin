package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var security_group_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete security group(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(security_group_deleteCmd).Standalone()

	security_groupCmd.AddCommand(security_group_deleteCmd)
}
