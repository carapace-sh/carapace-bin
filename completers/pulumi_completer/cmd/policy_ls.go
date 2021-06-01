package cmd

import (
	"github.com/spf13/cobra"
)

var policy_lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all Policy Packs for a Pulumi organization",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	policy_lsCmd.PersistentFlags().BoolP("json", "j", false, "Emit output as JSON")
	policyCmd.AddCommand(policy_lsCmd)
}
