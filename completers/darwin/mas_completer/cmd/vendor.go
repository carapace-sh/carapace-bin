package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vendorCmd = &cobra.Command{
	Use:   "vendor",
	Short: "Open vendor page for an app",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vendorCmd).Standalone()
	rootCmd.AddCommand(vendorCmd)
}
