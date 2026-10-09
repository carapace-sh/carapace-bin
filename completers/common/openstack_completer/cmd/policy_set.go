package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var policy_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set policy properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(policy_setCmd).Standalone()

	policy_setCmd.Flags().String("rules", "", "New serialized policy rules file")
	policy_setCmd.Flags().String("type", "", "New MIME type of the policy rules file")
	policyCmd.AddCommand(policy_setCmd)
}
