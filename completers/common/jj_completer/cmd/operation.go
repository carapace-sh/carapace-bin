package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var operationCmd = &cobra.Command{
	Use:     "operation",
	Short:   "Commands for working with the operation log",
	Aliases: []string{"op"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(operationCmd).Standalone()

	operationCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(operationCmd)
}
