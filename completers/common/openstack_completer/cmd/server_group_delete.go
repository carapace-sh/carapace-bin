package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var server_group_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete existing server group(s).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_group_deleteCmd).Standalone()

	server_groupCmd.AddCommand(server_group_deleteCmd)
}
