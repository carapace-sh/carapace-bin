package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var idlCmd = &cobra.Command{
	Use:   "idl",
	Short: "Commands for interacting with interface definitions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(idlCmd).Standalone()

	idlCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(idlCmd)
}
