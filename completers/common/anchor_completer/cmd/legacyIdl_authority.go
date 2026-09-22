package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var legacyIdl_authorityCmd = &cobra.Command{
	Use:   "authority",
	Short: "[DEPRECATED] Output the authority for the legacy IDL account",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(legacyIdl_authorityCmd).Standalone()

	legacyIdl_authorityCmd.Flags().BoolP("help", "h", false, "Print help")
	legacyIdlCmd.AddCommand(legacyIdl_authorityCmd)
}
