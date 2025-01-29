package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var manifest_inspectCmd = &cobra.Command{
	Use:   "inspect [options] IMAGE",
	Short: "Display the contents of a manifest list or image index",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(manifest_inspectCmd).Standalone()

	manifest_inspectCmd.Flags().String("authfile", "", "path of the authentication file. Use REGISTRY_AUTH_FILE environment variable to override")
	manifest_inspectCmd.Flags().Bool("insecure", false, "Purely for Docker compatibility")
	manifest_inspectCmd.Flags().Bool("tls-verify", false, "require HTTPS and verify certificates when accessing the registry")
	manifest_inspectCmd.Flags().BoolP("verbose", "v", false, "Added for Docker compatibility")
	manifest_inspectCmd.Flag("insecure").Hidden = true
	manifest_inspectCmd.Flag("verbose").Hidden = true
	manifestCmd.AddCommand(manifest_inspectCmd)
}
