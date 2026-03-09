package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var versionlock_clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "remove all versionlock entries",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(versionlock_clearCmd).Standalone()

	versionlockCmd.AddCommand(versionlock_clearCmd)
}
