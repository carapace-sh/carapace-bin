package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var manifest_annotateCmd = &cobra.Command{
	Use:   "annotate [OPTIONS] MANIFEST_LIST MANIFEST",
	Short: "Add additional information to a local image manifest",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(manifest_annotateCmd).Standalone()

	manifest_annotateCmd.Flags().String("arch", "", "Set architecture")
	manifest_annotateCmd.Flags().String("os", "", "Set operating system")
	manifest_annotateCmd.Flags().StringSlice("os-features", nil, "Set operating system feature")
	manifest_annotateCmd.Flags().String("os-version", "", "Set operating system version")
	manifest_annotateCmd.Flags().String("variant", "", "Set architecture variant")
	manifestCmd.AddCommand(manifest_annotateCmd)

	// TODO completion
}
