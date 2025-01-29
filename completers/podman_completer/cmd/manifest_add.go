package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var manifest_addCmd = &cobra.Command{
	Use:   "add [options] LIST IMAGEORARTIFACT [IMAGEORARTIFACT...]",
	Short: "Add images or artifacts to a manifest list or image index",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(manifest_addCmd).Standalone()

	manifest_addCmd.Flags().Bool("all", false, "add all of the list's images if the image is a list")
	manifest_addCmd.Flags().StringSlice("annotation", []string{}, "set an `annotation` for the specified image")
	manifest_addCmd.Flags().String("arch", "", "override the `architecture` of the specified image")
	manifest_addCmd.Flags().Bool("artifact", false, "add all arguments as artifact files rather than as images")
	manifest_addCmd.Flags().String("artifact-config", "", "artifact configuration file")
	manifest_addCmd.Flags().String("artifact-config-type", "", "artifact configuration media type")
	manifest_addCmd.Flags().Bool("artifact-exclude-titles", false, "refrain from setting \"org.opencontainers.image.title\" annotations on \"layers\"")
	manifest_addCmd.Flags().String("artifact-layer-type", "", "artifact layer media type")
	manifest_addCmd.Flags().String("artifact-subject", "", "artifact subject reference")
	manifest_addCmd.Flags().String("artifact-type", "", "override the artifactType value")
	manifest_addCmd.Flags().String("authfile", "", "path of the authentication file. Use REGISTRY_AUTH_FILE environment variable to override")
	manifest_addCmd.Flags().String("cert-dir", "", "use certificates at the specified path to access the registry")
	manifest_addCmd.Flags().String("creds", "", "use `[username[:password]]` for accessing the registry")
	manifest_addCmd.Flags().StringSlice("features", []string{}, "override the `features` of the specified image")
	manifest_addCmd.Flags().Bool("insecure", false, "neither require HTTPS nor verify certificates when accessing the registry")
	manifest_addCmd.Flags().String("os", "", "override the `OS` of the specified image")
	manifest_addCmd.Flags().String("os-version", "", "override the OS `version` of the specified image")
	manifest_addCmd.Flags().Bool("tls-verify", false, "require HTTPS and verify certificates when accessing the registry")
	manifest_addCmd.Flags().String("variant", "", "override the `Variant` of the specified image")
	manifest_addCmd.Flag("insecure").Hidden = true
	manifestCmd.AddCommand(manifest_addCmd)
}
