package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var legacyIdl_setBufferCmd = &cobra.Command{
	Use:   "set-buffer",
	Short: "[DEPRECATED] Set a new IDL buffer for the program",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(legacyIdl_setBufferCmd).Standalone()

	legacyIdl_setBufferCmd.Flags().StringP("buffer", "b", "", "Address of the buffer account to set as the IDL on the program")
	legacyIdl_setBufferCmd.Flags().BoolP("help", "h", false, "Print help")
	legacyIdl_setBufferCmd.Flags().Bool("print-only", false, "Print the instruction in base64 without executing it")
	legacyIdl_setBufferCmd.Flags().String("priority-fee", "", "")
	legacyIdl_setBufferCmd.MarkFlagRequired("buffer")
	legacyIdlCmd.AddCommand(legacyIdl_setBufferCmd)
}
