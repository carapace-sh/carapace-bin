package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var legacyIdl_eraseAuthorityCmd = &cobra.Command{
	Use:   "erase-authority",
	Short: "[DEPRECATED] Remove the ability to modify the legacy IDL account",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(legacyIdl_eraseAuthorityCmd).Standalone()

	legacyIdl_eraseAuthorityCmd.Flags().BoolP("help", "h", false, "Print help")
	legacyIdl_eraseAuthorityCmd.Flags().String("priority-fee", "", "")
	legacyIdl_eraseAuthorityCmd.Flags().StringP("program-id", "p", "", "")
	legacyIdl_eraseAuthorityCmd.MarkFlagRequired("program-id")
	legacyIdlCmd.AddCommand(legacyIdl_eraseAuthorityCmd)
}
