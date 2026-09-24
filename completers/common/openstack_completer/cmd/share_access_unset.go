package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_access_unsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Unset properties of share access rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_access_unsetCmd).Standalone()

	share_access_unsetCmd.Flags().String("property", "", "Remove property from share access rule.")
	share_accessCmd.AddCommand(share_access_unsetCmd)
}
