package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var versionlock_listCmd = &cobra.Command{
	Use:   "list",
	Short: "list current versionlock entries",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(versionlock_listCmd).Standalone()

	versionlockCmd.AddCommand(versionlock_listCmd)
}
