package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var versionlockClearCmd = &cobra.Command{
	Use:   "clear",
	Short: "remove all entries from versionlock configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(versionlockClearCmd).Standalone()

	versionlockCmd.AddCommand(versionlockClearCmd)
}
