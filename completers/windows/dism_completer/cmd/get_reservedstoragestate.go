package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var GetReservedStorageStateCmd = &cobra.Command{
	Use:   "Get-ReservedStorageState",
	Short: "get the current reserved storage state",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(GetReservedStorageStateCmd).Standalone()
	rootCmd.AddCommand(GetReservedStorageStateCmd)
}
