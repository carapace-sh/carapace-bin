package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var publishCmd = &cobra.Command{
	Use:   "publish [OPTIONS] REPOSITORY[:TAG]",
	Short: "Publish compose application",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(publishCmd).Standalone()

	publishCmd.Flags().String("oci-version", "", "OCI image/artifact specification version (automatically determined by default)")
	publishCmd.Flags().Bool("resolve-image-digests", false, "Pin image tags to digests")
	publishCmd.Flags().Bool("with-env", false, "Include environment variables in the published OCI artifact")
	publishCmd.Flags().BoolP("yes", "y", false, "Assume \"yes\" as answer to all prompts")
	rootCmd.AddCommand(publishCmd)
}
