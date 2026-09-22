package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var idl_closeCmd = &cobra.Command{
	Use:   "close",
	Short: "Close a metadata account and recover rent",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(idl_closeCmd).Standalone()

	idl_closeCmd.Flags().BoolP("help", "h", false, "Print help")
	idl_closeCmd.Flags().String("priority-fee", "", "Priority fees in micro-lamports per compute unit")
	idl_closeCmd.Flags().String("seed", "idl", "The seed used for the metadata account (default: \"idl\")")
	idlCmd.AddCommand(idl_closeCmd)
}
