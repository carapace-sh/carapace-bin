package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var policy_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete policy(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(policy_deleteCmd).Standalone()

	policyCmd.AddCommand(policy_deleteCmd)
}
