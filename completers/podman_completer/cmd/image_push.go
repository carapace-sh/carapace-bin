package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var image_pushCmd = &cobra.Command{
	Use:   "push [options] IMAGE [DESTINATION]",
	Short: "Push an image to a specified destination",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_pushCmd).Standalone()

	image_pushCmd.Flags().String("authfile", "", "Path of the authentication file. Use REGISTRY_AUTH_FILE environment variable to override")
	image_pushCmd.Flags().String("cert-dir", "", "Path to a directory containing TLS certificates and keys")
	image_pushCmd.Flags().Bool("compress", false, "Compress tarball image layers when pushing to a directory using the 'dir' transport. (default is same compression type as source)")
	image_pushCmd.Flags().String("compression-format", "", "compression format to use")
	image_pushCmd.Flags().String("compression-level", "", "compression level to use")
	image_pushCmd.Flags().String("creds", "", "`Credentials` (USERNAME:PASSWORD) to use for authenticating to a registry")
	image_pushCmd.Flags().String("digestfile", "", "Write the digest of the pushed image to the specified file")
	image_pushCmd.Flags().Bool("disable-content-trust", false, "This is a Docker specific option and is a NOOP")
	image_pushCmd.Flags().StringSlice("encrypt-layer", []string{}, "Layers to encrypt, 0-indexed layer indices with support for negative indexing (e.g. 0 is the first layer, -1 is the last layer). If not defined, will encrypt all layers if encryption-key flag is specified")
	image_pushCmd.Flags().StringSlice("encryption-key", []string{}, "Key with the encryption protocol to use to encrypt the image (e.g. jwe:/path/to/key.pem)")
	image_pushCmd.Flags().Bool("force-compression", false, "Use the specified compression algorithm even if the destination contains a differently-compressed variant already")
	image_pushCmd.Flags().StringP("format", "f", "", "Manifest type (oci, v2s2, or v2s1) to use in the destination (default is manifest type of source, with fallbacks)")
	image_pushCmd.Flags().BoolP("quiet", "q", false, "Suppress output information when pushing images")
	image_pushCmd.Flags().Bool("remove-signatures", false, "Discard any pre-existing signatures in the image")
	image_pushCmd.Flags().String("retry", "", "number of times to retry in case of failure when performing push")
	image_pushCmd.Flags().String("retry-delay", "", "delay between retries in case of push failures")
	image_pushCmd.Flags().String("sign-by", "", "Add a signature at the destination using the specified key")
	image_pushCmd.Flags().String("sign-by-sigstore", "", "Sign the image using a sigstore parameter file at `PATH`")
	image_pushCmd.Flags().String("sign-by-sigstore-private-key", "", "Sign the image using a sigstore private key at `PATH`")
	image_pushCmd.Flags().String("sign-passphrase-file", "", "Read a passphrase for signing an image from `PATH`")
	image_pushCmd.Flags().String("signature-policy", "", "Path to a signature-policy file")
	image_pushCmd.Flags().Bool("tls-verify", false, "Require HTTPS and verify certificates when contacting registries")
	image_pushCmd.Flag("signature-policy").Hidden = true
	imageCmd.AddCommand(image_pushCmd)
}
