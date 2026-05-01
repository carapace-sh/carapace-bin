package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var osCmd = &cobra.Command{
	Use:   "os",
	Short: "NixOS functionality",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(osCmd).Standalone()

	rootCmd.AddCommand(osCmd)
}
