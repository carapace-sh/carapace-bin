package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var idl_writeBufferCmd = &cobra.Command{
	Use:   "write-buffer",
	Short: "Write metadata using a buffer account",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(idl_writeBufferCmd).Standalone()

	idl_writeBufferCmd.Flags().StringP("buffer", "b", "", "The buffer account address")
	idl_writeBufferCmd.Flags().Bool("close-buffer", false, "Close the buffer after writing")
	idl_writeBufferCmd.Flags().BoolP("help", "h", false, "Print help")
	idl_writeBufferCmd.Flags().String("priority-fee", "", "Priority fees in micro-lamports per compute unit")
	idl_writeBufferCmd.Flags().String("seed", "idl", "The seed to use for the metadata account (default: \"idl\")")
	idl_writeBufferCmd.MarkFlagRequired("buffer")
	idlCmd.AddCommand(idl_writeBufferCmd)
}
