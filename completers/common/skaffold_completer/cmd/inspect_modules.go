package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspect_modulesCmd = &cobra.Command{
	Use:   "modules",
	Short: "Interact with configuration modules",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspect_modulesCmd).Standalone()

	applyCommonFlagCompletion(inspect_modulesCmd)
	inspectCmd.AddCommand(inspect_modulesCmd)
}
