package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var delhostCmd = &cobra.Command{
	Use:   "delhost",
	Short: "delete a host-to-realm mapping",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(delhostCmd).Standalone()
	rootCmd.AddCommand(delhostCmd)
}
