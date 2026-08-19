package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var showOptionCmd = &cobra.Command{
	Use:   "show-option",
	Short: "print the value of a configuration option",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(showOptionCmd).Standalone()
	rootCmd.AddCommand(showOptionCmd)
}
