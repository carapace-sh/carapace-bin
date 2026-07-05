package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootCmd = &cobra.Command{
	Use:   "boot",
	Short: "indicate whether the last boot should be saved as last-known-good",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootCmd).Standalone()
	rootCmd.AddCommand(bootCmd)
}
