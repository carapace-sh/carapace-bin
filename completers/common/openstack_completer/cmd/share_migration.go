package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_migrationCmd = &cobra.Command{
	Use:   "migration",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_migrationCmd).Standalone()

	shareCmd.AddCommand(share_migrationCmd)
}
