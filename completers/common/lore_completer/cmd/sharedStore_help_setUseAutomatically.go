package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sharedStore_help_setUseAutomaticallyCmd = &cobra.Command{
	Use:   "set-use-automatically",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sharedStore_help_setUseAutomaticallyCmd).Standalone()

	sharedStore_helpCmd.AddCommand(sharedStore_help_setUseAutomaticallyCmd)
}
