package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var group_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete group(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(group_deleteCmd).Standalone()

	group_deleteCmd.Flags().String("domain", "", "Domain containing group(s) (name or ID)")
	groupCmd.AddCommand(group_deleteCmd)
}
