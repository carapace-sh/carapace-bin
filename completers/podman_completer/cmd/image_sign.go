package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var image_signCmd = &cobra.Command{
	Use:   "sign [options] IMAGE [IMAGE...]",
	Short: "Sign an image",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_signCmd).Standalone()

	image_signCmd.Flags().BoolP("all", "a", false, "Sign all the manifests of the multi-architecture image")
	image_signCmd.Flags().String("authfile", "", "Path of the authentication file. Use REGISTRY_AUTH_FILE environment variable to override")
	image_signCmd.Flags().String("cert-dir", "", "`Pathname` of a directory containing TLS certificates and keys")
	image_signCmd.Flags().StringP("directory", "d", "", "Define an alternate directory to store signatures")
	image_signCmd.Flags().String("sign-by", "", "Name of the signing key")
	imageCmd.AddCommand(image_signCmd)
}
