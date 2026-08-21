package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var image_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create/upload an image",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_createCmd).Standalone()

	image_createCmd.Flags().String("checksum", "", "==SUPPRESS==")
	image_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	image_createCmd.Flags().Bool("community", false, "Image is accessible by all users but does not appear in the default image list of any user except the owner (requires --os-image-api-version 2.5 or later)")
	image_createCmd.Flags().String("container-format", "", "Image container format.")
	image_createCmd.Flags().String("copy-from", "", "==SUPPRESS==")
	image_createCmd.Flags().String("disk-format", "", "Image disk format.")
	image_createCmd.Flags().String("file", "", "Upload image from local file")
	image_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	image_createCmd.Flags().Bool("force", false, "Force image creation if volume is in use (only meaningful with --volume)")
	image_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	image_createCmd.Flags().String("id", "", "Image ID to reserve")
	image_createCmd.Flags().Bool("import", false, "Force the use of glance image import instead of direct upload")
	image_createCmd.Flags().String("location", "", "==SUPPRESS==")
	image_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	image_createCmd.Flags().String("min-disk", "", "Minimum disk size needed to boot image, in gigabytes")
	image_createCmd.Flags().String("min-ram", "", "Minimum RAM size needed to boot image, in megabytes")
	image_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	image_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	image_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	image_createCmd.Flags().Bool("private", false, "Image is only accessible by the owner (default until --os-image-api-version 2.5)")
	image_createCmd.Flags().Bool("progress", false, "Show upload progress bar (ignored if passing data via stdin)")
	image_createCmd.Flags().String("project", "", "Set an alternate project on this image (name or ID)")
	image_createCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	image_createCmd.Flags().String("property", "", "Set a property on this image (repeat option to set multiple properties)")
	image_createCmd.Flags().Bool("protected", false, "Prevent image from being deleted")
	image_createCmd.Flags().Bool("public", false, "Image is accessible and visible to all users")
	image_createCmd.Flags().Bool("shared", false, "Image is only accessible by the owner and image members (requires --os-image-api-version 2.5 or later) (default since --os-image-api-version 2.5)")
	image_createCmd.Flags().String("sign-cert-id", "", "The specified certificate UUID is a reference to the certificate in the key manager that corresponds to the public key and is used for signature validation.")
	image_createCmd.Flags().String("sign-key-path", "", "Sign the image using the specified private key.")
	image_createCmd.Flags().String("size", "", "==SUPPRESS==")
	image_createCmd.Flags().String("store", "", "==SUPPRESS==")
	image_createCmd.Flags().String("tag", "", "Set a tag on this image (repeat option to set multiple tags)")
	image_createCmd.Flags().Bool("unprotected", false, "Allow image to be deleted (default)")
	image_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	image_createCmd.Flags().String("volume", "", "Create image from a volume")
	imageCmd.AddCommand(image_createCmd)
}
