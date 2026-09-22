package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var idl_setBufferAuthorityCmd = &cobra.Command{
	Use:   "set-buffer-authority",
	Short: "Set a new authority on a buffer account",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(idl_setBufferAuthorityCmd).Standalone()

	idl_setBufferAuthorityCmd.Flags().BoolP("help", "h", false, "Print help")
	idl_setBufferAuthorityCmd.Flags().StringP("new-authority", "n", "", "The new authority")
	idl_setBufferAuthorityCmd.Flags().String("priority-fee", "", "Priority fees in micro-lamports per compute unit")
	idl_setBufferAuthorityCmd.MarkFlagRequired("new-authority")
	idlCmd.AddCommand(idl_setBufferAuthorityCmd)
}
