package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var local_ip_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set Local IP properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(local_ip_setCmd).Standalone()

	local_ip_setCmd.Flags().String("description", "", "Set Local IP description")
	local_ip_setCmd.Flags().String("name", "", "Set local IP name")
	local_ipCmd.AddCommand(local_ip_setCmd)
}
