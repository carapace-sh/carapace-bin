package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var printDisabledCmd = &cobra.Command{
	Use:   "print-disabled",
	Short: "Prints which services are disabled",
	Run:   func(*cobra.Command, []string) {},
}

func init() {
	carapace.Gen(printDisabledCmd).Standalone()
	rootCmd.AddCommand(printDisabledCmd)
}
