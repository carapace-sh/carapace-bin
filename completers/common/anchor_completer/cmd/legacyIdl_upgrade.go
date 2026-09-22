package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var legacyIdl_upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "[DEPRECATED] Upgrade the legacy IDL (write buffer → set buffer → close buffer)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(legacyIdl_upgradeCmd).Standalone()

	legacyIdl_upgradeCmd.Flags().StringP("filepath", "f", "", "")
	legacyIdl_upgradeCmd.Flags().BoolP("help", "h", false, "Print help")
	legacyIdl_upgradeCmd.Flags().String("priority-fee", "", "")
	legacyIdl_upgradeCmd.MarkFlagRequired("filepath")
	legacyIdlCmd.AddCommand(legacyIdl_upgradeCmd)
}
