package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var artifact_pushCmd = &cobra.Command{
	Use:   "push [options] ARTIFACT.",
	Short: "Push an OCI artifact",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(artifact_pushCmd).Standalone()

	artifact_pushCmd.Flags().String("authfile", "", "Path of the authentication file. Use REGISTRY_AUTH_FILE environment variable to override")
	artifact_pushCmd.Flags().String("cert-dir", "", "Path to a directory containing TLS certificates and keys")
	artifact_pushCmd.Flags().String("creds", "", "`Credentials` (USERNAME:PASSWORD) to use for authenticating to a registry")
	artifact_pushCmd.Flags().String("digestfile", "", "Write the digest of the pushed image to the specified file")
	artifact_pushCmd.Flags().BoolP("quiet", "q", false, "Suppress output information when pushing images")
	artifact_pushCmd.Flags().String("retry", "", "number of times to retry in case of failure when performing push")
	artifact_pushCmd.Flags().String("retry-delay", "", "delay between retries in case of push failures")
	artifact_pushCmd.Flags().String("sign-by", "", "Add a signature at the destination using the specified key")
	artifact_pushCmd.Flags().String("sign-by-sigstore", "", "Sign the image using a sigstore parameter file at `PATH`")
	artifact_pushCmd.Flags().String("sign-by-sigstore-private-key", "", "Sign the image using a sigstore private key at `PATH`")
	artifact_pushCmd.Flags().String("sign-passphrase-file", "", "Read a passphrase for signing an image from `PATH`")
	artifact_pushCmd.Flags().String("signature-policy", "", "Path to a signature-policy file")
	artifact_pushCmd.Flags().Bool("tls-verify", false, "Require HTTPS and verify certificates when contacting registries")
	artifact_pushCmd.Flag("signature-policy").Hidden = true
	artifactCmd.AddCommand(artifact_pushCmd)
}
