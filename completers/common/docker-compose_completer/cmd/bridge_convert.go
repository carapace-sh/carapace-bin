package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bridge_convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert compose files to Kubernetes manifests, Helm charts, or another model",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bridge_convertCmd).Standalone()

	bridge_convertCmd.Flags().StringP("output", "o", "out", "The output directory for the Kubernetes resources")
	bridge_convertCmd.Flags().String("templates", "", "Directory containing transformation templates")
	bridge_convertCmd.Flags().StringSliceP("transformation", "t", nil, "Transformation to apply to compose model (default: docker/compose-bridge-kubernetes)")
	bridge_convertCmd.Flags().BoolP("yes", "y", false, "Assume \"yes\" to the output directory overwrite prompt. For scripts/CI, where no interactive confirmation is possible")
	bridgeCmd.AddCommand(bridge_convertCmd)

	carapace.Gen(bridge_convertCmd).FlagCompletion(carapace.ActionMap{
		"output":    carapace.ActionDirectories(),
		"templates": carapace.ActionDirectories(),
		// TODO complete transformations
	})
}
