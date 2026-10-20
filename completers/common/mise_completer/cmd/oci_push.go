package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var oci_pushCmd = &cobra.Command{
	Use:   "push",
	Short: "[experimental] Build an OCI image and push it to a registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(oci_pushCmd).Standalone()

	oci_pushCmd.Flags().String("cache-from", "", "Reuse unchanged tool layers from this image instead of the destination's current tag")
	oci_pushCmd.Flags().String("from", "", "Base image for the build")
	oci_pushCmd.Flags().BoolP("help", "h", false, "Print help")
	oci_pushCmd.Flags().String("image-dir", "", "Push an existing OCI image layout instead of building one")
	oci_pushCmd.Flags().Bool("include-global", false, "Also package tools from the global and system configs")
	oci_pushCmd.Flags().String("mount-point", "", "Where tools install inside the image")
	oci_pushCmd.Flags().Bool("no-cache", false, "Rebuild tool layers without using the remote or local layer cache")
	oci_pushCmd.Flags().Bool("no-mise", false, "Do not embed the running mise binary at /usr/local/bin/mise")
	oci_pushCmd.Flags().String("owner", "", "UID[:GID] to assign to every tar entry when building")
	oci_pushCmd.Flags().Bool("update-index", false, "Maintain the tag as a multi-arch image index")
	ociCmd.AddCommand(oci_pushCmd)

	carapace.Gen(oci_pushCmd).FlagCompletion(carapace.ActionMap{
		"image-dir": carapace.ActionDirectories(),
	})
}
