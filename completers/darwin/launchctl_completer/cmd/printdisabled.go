package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var printDisabledCmd = &cobra.Command{
	Use:   "print-disabled",
	Short: "Print the list of disabled services",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(printDisabledCmd).Standalone()
	rootCmd.AddCommand(printDisabledCmd)
}
