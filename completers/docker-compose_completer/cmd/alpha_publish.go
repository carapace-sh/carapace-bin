package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var alpha_publishCmd = &cobra.Command{
	Use:   "publish [OPTIONS] REPOSITORY[:TAG]",
	Short: "Publish compose application",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(alpha_publishCmd).Standalone()

	alpha_publishCmd.Flags().String("oci-version", "", "OCI Image/Artifact specification version (automatically determined by default)")
	alpha_publishCmd.Flags().Bool("resolve-image-digests", false, "Pin image tags to digests")
	alphaCmd.AddCommand(alpha_publishCmd)
}
