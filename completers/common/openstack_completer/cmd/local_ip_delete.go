package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var local_ip_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete Local IP(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(local_ip_deleteCmd).Standalone()

	local_ipCmd.AddCommand(local_ip_deleteCmd)
}
