package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspect_profilesCmd = &cobra.Command{
	Use:   "profiles",
	Short: "Interact with configuration profiles",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspect_profilesCmd).Standalone()

	inspect_profilesCmd.PersistentFlags().StringSliceP("module", "m", nil, "Names of modules to filter target action by.")
	applyCommonFlagCompletion(inspect_profilesCmd)
	inspectCmd.AddCommand(inspect_profilesCmd)
}
