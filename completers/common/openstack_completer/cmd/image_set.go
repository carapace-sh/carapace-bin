package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var image_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set image properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_setCmd).Standalone()

	image_setCmd.Flags().Bool("accept", false, "Accept the image membership for either the project indicated by '--project', if provided, or the current user's project")
	image_setCmd.Flags().Bool("activate", false, "Activate the image")
	image_setCmd.Flags().String("architecture", "", "Operating system architecture")
	image_setCmd.Flags().Bool("community", false, "Image is accessible by all users but does not appear in the default image list of any user except the owner (requires --os-image-api-version 2.5 or later)")
	image_setCmd.Flags().String("container-format", "", "Image container format.")
	image_setCmd.Flags().Bool("deactivate", false, "Deactivate the image")
	image_setCmd.Flags().String("disk-format", "", "Image disk format.")
	image_setCmd.Flags().Bool("hidden", false, "Hide the image")
	image_setCmd.Flags().String("instance-id", "", "ID of server instance used to create this image")
	image_setCmd.Flags().String("instance-uuid", "", "==SUPPRESS==")
	image_setCmd.Flags().String("kernel-id", "", "ID of kernel image used to boot this disk image")
	image_setCmd.Flags().String("min-disk", "", "Minimum disk size needed to boot image, in gigabytes")
	image_setCmd.Flags().String("min-ram", "", "Minimum RAM size needed to boot image, in megabytes")
	image_setCmd.Flags().String("name", "", "New image name")
	image_setCmd.Flags().String("os-distro", "", "Operating system distribution name")
	image_setCmd.Flags().String("os-version", "", "Operating system distribution version")
	image_setCmd.Flags().Bool("pending", false, "Reset the image membership to 'pending'")
	image_setCmd.Flags().Bool("private", false, "Image is only accessible by the owner (default until --os-image-api-version 2.5)")
	image_setCmd.Flags().String("project", "", "Set an alternate project on this image (name or ID)")
	image_setCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	image_setCmd.Flags().String("property", "", "Set a property on this image (repeat option to set multiple properties)")
	image_setCmd.Flags().Bool("protected", false, "Prevent image from being deleted")
	image_setCmd.Flags().Bool("public", false, "Image is accessible and visible to all users")
	image_setCmd.Flags().String("ramdisk-id", "", "ID of ramdisk image used to boot this disk image")
	image_setCmd.Flags().Bool("reject", false, "Reject the image membership for either the project indicated by '--project', if provided, or the current user's project")
	image_setCmd.Flags().Bool("shared", false, "Image is only accessible by the owner and image members (requires --os-image-api-version 2.5 or later) (default since --os-image-api-version 2.5)")
	image_setCmd.Flags().String("tag", "", "Set a tag on this image (repeat option to set multiple tags)")
	image_setCmd.Flags().Bool("unhidden", false, "Unhide the image")
	image_setCmd.Flags().Bool("unprotected", false, "Allow image to be deleted (default)")
	imageCmd.AddCommand(image_setCmd)
}
