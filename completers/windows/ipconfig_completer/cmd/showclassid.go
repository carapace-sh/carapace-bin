package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var showclassidCmd = &cobra.Command{
	Use:   "showclassid",
	Short: "display all DHCP class IDs for an adapter",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(showclassidCmd).Standalone()
	rootCmd.AddCommand(showclassidCmd)
}
