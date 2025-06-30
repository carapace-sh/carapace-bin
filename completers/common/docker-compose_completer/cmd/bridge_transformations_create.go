package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bridge_transformations_createCmd = &cobra.Command{
	Use:   "create [OPTION] PATH",
	Short: "Create a new transformation",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bridge_transformations_createCmd).Standalone()

	bridge_transformations_createCmd.Flags().StringP("from", "f", "", "Existing transformation to copy (default: docker/compose-bridge-kubernetes)")
	bridge_transformationsCmd.AddCommand(bridge_transformations_createCmd)

	// TODO flag completion

	carapace.Gen(bridge_transformations_createCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
