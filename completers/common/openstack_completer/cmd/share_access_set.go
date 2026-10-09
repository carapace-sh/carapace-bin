package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_access_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set properties to share access rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_access_setCmd).Standalone()

	share_access_setCmd.Flags().String("access-level", "", "Share access level (\"rw\" and \"ro\" access levels are supported) to set.")
	share_access_setCmd.Flags().String("property", "", "Set a property to this share access rule.")
	share_accessCmd.AddCommand(share_access_setCmd)
}
