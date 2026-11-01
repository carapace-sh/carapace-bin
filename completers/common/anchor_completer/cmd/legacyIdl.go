package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var legacyIdlCmd = &cobra.Command{
	Use:   "legacy-idl",
	Short: "[DEPRECATED] Manage legacy on-chain IDL accounts. These commands interact with the old Anchor IDL instruction protocol and will be removed in a future release. Migrate to Program Metadata-based IDL management (`anchor idl`)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(legacyIdlCmd).Standalone()

	legacyIdlCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(legacyIdlCmd)
}
