package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var image_importCmd = &cobra.Command{
	Use:   "import",
	Short: "Initiate the image import process.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_importCmd).Standalone()

	image_importCmd.Flags().Bool("all-stores", false, "Make image available to all stores (either '--store' or '--all-stores' required with the 'copy-image' import method)")
	image_importCmd.Flags().Bool("allow-failure", false, "When uploading to multiple stores, indicate that the import should be continue should any of the uploads fail.")
	image_importCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	image_importCmd.Flags().Bool("disallow-failure", false, "When uploading to multiple stores, indicate that the import should be reverted should any of the uploads fail.")
	image_importCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	image_importCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	image_importCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	image_importCmd.Flags().String("method", "", "Import method used for image import process.")
	image_importCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	image_importCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	image_importCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	image_importCmd.Flags().String("remote-image", "", "The image of remote glance (ID only) to be imported (only valid with the 'glance-download' import method)")
	image_importCmd.Flags().String("remote-region", "", "The remote Glance region to download the image from (only valid with the 'glance-download' import method)")
	image_importCmd.Flags().String("remote-service-interface", "", "The remote Glance service interface to use when importing images (only valid with the 'glance-download' import method)")
	image_importCmd.Flags().String("store", "", "Backend store to upload image to (specify multiple times to upload to multiple stores) (either '--store' or '--all-stores' required with the 'copy-image' import method)")
	image_importCmd.Flags().String("uri", "", "URI to download the external image (only valid with the 'web-download' import method)")
	image_importCmd.Flags().String("variable", "", "==SUPPRESS==")
	image_importCmd.Flags().Bool("wait", false, "Wait for operation to complete")
	imageCmd.AddCommand(image_importCmd)
}
