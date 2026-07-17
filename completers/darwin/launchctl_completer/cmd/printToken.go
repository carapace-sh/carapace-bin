package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var printTokenCmd = &cobra.Command{
	Use:   "print-token",
	Short: "Print service identifier given an xpc event token",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(printTokenCmd).Standalone()
	rootCmd.AddCommand(printTokenCmd)
}
