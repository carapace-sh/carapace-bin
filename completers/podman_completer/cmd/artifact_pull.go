package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var artifact_pullCmd = &cobra.Command{
	Use:   "pull [options] ARTIFACT",
	Short: "Pull an OCI artifact",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(artifact_pullCmd).Standalone()

	artifact_pullCmd.Flags().String("authfile", "", "Path of the authentication file. Use REGISTRY_AUTH_FILE environment variable to override")
	artifact_pullCmd.Flags().String("cert-dir", "", "`Pathname` of a directory containing TLS certificates and keys")
	artifact_pullCmd.Flags().String("creds", "", "`Credentials` (USERNAME:PASSWORD) to use for authenticating to a registry")
	artifact_pullCmd.Flags().StringSlice("decryption-key", []string{}, "Key needed to decrypt the image (e.g. /path/to/key.pem)")
	artifact_pullCmd.Flags().BoolP("quiet", "q", false, "Suppress output information when pulling images")
	artifact_pullCmd.Flags().String("retry", "", "number of times to retry in case of failure when performing pull")
	artifact_pullCmd.Flags().String("retry-delay", "", "delay between retries in case of pull failures")
	artifact_pullCmd.Flags().Bool("tls-verify", false, "Require HTTPS and verify certificates when contacting registries")
	artifactCmd.AddCommand(artifact_pullCmd)
}
