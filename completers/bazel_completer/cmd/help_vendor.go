package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_vendorCmd = &cobra.Command{
	Use:   "vendor",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_vendorCmd).Standalone()

	helpCmd.AddCommand(help_vendorCmd)
}
