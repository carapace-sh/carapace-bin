package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var setclassidCmd = &cobra.Command{
	Use:   "setclassid",
	Short: "modify the DHCP class ID for an adapter",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setclassidCmd).Standalone()
	rootCmd.AddCommand(setclassidCmd)
}
