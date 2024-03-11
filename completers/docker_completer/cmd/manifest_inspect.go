package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var manifest_inspectCmd = &cobra.Command{
	Use:   "inspect [OPTIONS] [MANIFEST_LIST] MANIFEST",
	Short: "Display an image manifest, or manifest list",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(manifest_inspectCmd).Standalone()

	manifest_inspectCmd.Flags().Bool("insecure", false, "Allow communication with an insecure registry")
	manifest_inspectCmd.Flags().BoolP("verbose", "v", false, "Output additional info including layers and platform")
	manifestCmd.AddCommand(manifest_inspectCmd)

	// TODO completion
}
