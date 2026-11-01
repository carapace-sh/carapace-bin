package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resetPermissionsCmd = &cobra.Command{
	Use:   "reset-permissions",
	Short: "Revert all runtime permissions to their default state",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resetPermissionsCmd).Standalone()

	rootCmd.AddCommand(resetPermissionsCmd)

}
