package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var legacyIdl_setAuthorityCmd = &cobra.Command{
	Use:   "set-authority",
	Short: "[DEPRECATED] Set a new authority on the legacy IDL account",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(legacyIdl_setAuthorityCmd).Standalone()

	legacyIdl_setAuthorityCmd.Flags().BoolP("help", "h", false, "Print help")
	legacyIdl_setAuthorityCmd.Flags().StringP("new-authority", "n", "", "New authority of the IDL account")
	legacyIdl_setAuthorityCmd.Flags().Bool("print-only", false, "Print the instruction in base64 without executing it")
	legacyIdl_setAuthorityCmd.Flags().String("priority-fee", "", "")
	legacyIdl_setAuthorityCmd.Flags().StringP("program-id", "p", "", "Program to change the IDL authority")
	legacyIdl_setAuthorityCmd.MarkFlagRequired("new-authority")
	legacyIdl_setAuthorityCmd.MarkFlagRequired("program-id")
	legacyIdlCmd.AddCommand(legacyIdl_setAuthorityCmd)
}
