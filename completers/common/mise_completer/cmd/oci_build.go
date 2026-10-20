package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var oci_buildCmd = &cobra.Command{
	Use:   "build",
	Short: "[experimental] Build an OCI image from the current mise.toml",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(oci_buildCmd).Standalone()

	oci_buildCmd.Flags().StringSlice("copy", nil, "Copy a host file, directory, or symlink into the image (repeatable, HOST:IMAGE)")
	oci_buildCmd.Flags().String("from", "", "Base image reference (overrides [oci].from and the oci.default_from setting)")
	oci_buildCmd.Flags().BoolP("help", "h", false, "Print help")
	oci_buildCmd.Flags().Bool("include-global", false, "Also package tools from the global and system configs")
	oci_buildCmd.Flags().String("mount-point", "", "Where tools install inside the image")
	oci_buildCmd.Flags().Bool("no-cache", false, "Rebuild tool layers without reading or writing the local layer cache")
	oci_buildCmd.Flags().Bool("no-mise", false, "Do not embed the currently-running mise binary at /usr/local/bin/mise")
	oci_buildCmd.Flags().StringP("output", "o", "./mise-oci", "Output directory for the OCI image layout")
	oci_buildCmd.Flags().String("owner", "", "UID[:GID] to assign to every tar entry in generated layers")
	oci_buildCmd.Flags().StringP("tag", "t", "", "Tag to record in the image index (the org.opencontainers.image.ref.name annotation)")
	ociCmd.AddCommand(oci_buildCmd)

	carapace.Gen(oci_buildCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionDirectories(),
	})
}
