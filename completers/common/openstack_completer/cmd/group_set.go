package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var group_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set group properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(group_setCmd).Standalone()

	group_setCmd.Flags().String("description", "", "New group description")
	group_setCmd.Flags().String("domain", "", "Domain containing <group> (name or ID)")
	group_setCmd.Flags().String("name", "", "New group name")
	groupCmd.AddCommand(group_setCmd)
}
