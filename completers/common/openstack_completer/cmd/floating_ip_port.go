package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var floating_ip_portCmd = &cobra.Command{
	Use:   "port",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(floating_ip_portCmd).Standalone()

	floating_ipCmd.AddCommand(floating_ip_portCmd)
}
