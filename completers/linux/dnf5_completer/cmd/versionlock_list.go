package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var versionlockListCmd = &cobra.Command{
	Use:   "list",
	Short: "list the current versionlock configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(versionlockListCmd).Standalone()

	versionlockCmd.AddCommand(versionlockListCmd)
}
