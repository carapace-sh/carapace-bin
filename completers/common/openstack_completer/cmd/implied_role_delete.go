package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var implied_role_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes an association between prior and implied roles",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(implied_role_deleteCmd).Standalone()

	implied_role_deleteCmd.Flags().String("implied-role", "", "<role> (name or ID) implied by another role")
	implied_role_deleteCmd.MarkFlagRequired("implied-role")
	implied_roleCmd.AddCommand(implied_role_deleteCmd)
}
