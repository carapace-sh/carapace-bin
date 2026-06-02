package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootCmd = &cobra.Command{
	Use:   "boot",
	Short: "Build the new configuration and make it the boot default, but do not activate it",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootCmd).Standalone()
	rootCmd.AddCommand(bootCmd)
}
