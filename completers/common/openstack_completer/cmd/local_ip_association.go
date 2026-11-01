package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var local_ip_associationCmd = &cobra.Command{
	Use:   "association",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(local_ip_associationCmd).Standalone()

	local_ipCmd.AddCommand(local_ip_associationCmd)
}
