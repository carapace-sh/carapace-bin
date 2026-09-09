package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var endpoint_group_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete endpoint group(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(endpoint_group_deleteCmd).Standalone()

	endpoint_groupCmd.AddCommand(endpoint_group_deleteCmd)
}
