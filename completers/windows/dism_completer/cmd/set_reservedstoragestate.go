package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var SetReservedStorageStateCmd = &cobra.Command{
	Use:   "Set-ReservedStorageState",
	Short: "set reserved storage to enabled or disabled",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(SetReservedStorageStateCmd).Standalone()
	rootCmd.AddCommand(SetReservedStorageStateCmd)
}
