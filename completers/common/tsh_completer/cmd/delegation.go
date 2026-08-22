package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var delegationCmd = &cobra.Command{
	Use:   "delegation",
	Short: "Manage delegation sessions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(delegationCmd).Standalone()

	rootCmd.AddCommand(delegationCmd)
}
