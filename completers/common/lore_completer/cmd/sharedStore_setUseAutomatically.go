package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sharedStore_setUseAutomaticallyCmd = &cobra.Command{
	Use:   "set-use-automatically",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sharedStore_setUseAutomaticallyCmd).Standalone()

	sharedStore_setUseAutomaticallyCmd.Flags().BoolP("help", "h", false, "Print help")
	sharedStoreCmd.AddCommand(sharedStore_setUseAutomaticallyCmd)

	carapace.Gen(sharedStore_setUseAutomaticallyCmd).PositionalCompletion(
		carapace.ActionValues("true", "false"),
	)
}
