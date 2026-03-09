package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var versionlock_excludeCmd = &cobra.Command{
	Use:   "exclude <package-spec>...",
	Short: "add an exclude within versionlock",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(versionlock_excludeCmd).Standalone()

	versionlockCmd.AddCommand(versionlock_excludeCmd)
}
