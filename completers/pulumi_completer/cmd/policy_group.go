package cmd

import (
	"github.com/spf13/cobra"
)

var policy_groupCmd = &cobra.Command{
	Use:   "group",
	Short: "Manage policy groups",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	policyCmd.AddCommand(policy_groupCmd)
}
