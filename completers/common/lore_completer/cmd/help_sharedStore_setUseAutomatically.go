package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_sharedStore_setUseAutomaticallyCmd = &cobra.Command{
	Use:   "set-use-automatically",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_sharedStore_setUseAutomaticallyCmd).Standalone()

	help_sharedStoreCmd.AddCommand(help_sharedStore_setUseAutomaticallyCmd)
}
