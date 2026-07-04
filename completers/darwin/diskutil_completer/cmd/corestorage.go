package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var coreStorageCmd = &cobra.Command{
	Use:   "coreStorage",
	Short: "CoreStorage logical volume operations",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(coreStorageCmd).Standalone()
	rootCmd.AddCommand(coreStorageCmd)
}
