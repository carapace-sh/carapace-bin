package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var legacyIdl_closeCmd = &cobra.Command{
	Use:   "close",
	Short: "[DEPRECATED] Close the legacy IDL account and recover rent",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(legacyIdl_closeCmd).Standalone()

	legacyIdl_closeCmd.Flags().BoolP("help", "h", false, "Print help")
	legacyIdl_closeCmd.Flags().String("idl-address", "", "The IDL account to close. If none is given, the IDL account derived from program_id is used")
	legacyIdl_closeCmd.Flags().Bool("print-only", false, "Print the instruction in base64 without executing it. Useful for multisig execution when the local wallet keypair is not available")
	legacyIdl_closeCmd.Flags().String("priority-fee", "", "")
	legacyIdlCmd.AddCommand(legacyIdl_closeCmd)
}
