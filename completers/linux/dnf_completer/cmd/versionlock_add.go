package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var versionlock_addCmd = &cobra.Command{
	Use:   "add <package-spec>...",
	Short: "add a versionlock",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(versionlock_addCmd).Standalone()

	versionlockCmd.AddCommand(versionlock_addCmd)
}
