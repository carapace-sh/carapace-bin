package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var image_importCmd = &cobra.Command{
	Use:   "import [options] PATH [REFERENCE]",
	Short: "Import a tarball to create a filesystem image",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_importCmd).Standalone()

	image_importCmd.Flags().String("arch", "", "Set the architecture of the imported image")
	image_importCmd.Flags().StringSliceP("change", "c", []string{}, "Apply the following possible instructions to the created image (default []): CMD | ENTRYPOINT | ENV | EXPOSE | LABEL | ONBUILD | STOPSIGNAL | USER | VOLUME | WORKDIR")
	image_importCmd.Flags().StringP("message", "m", "", "Set commit message for imported image")
	image_importCmd.Flags().String("os", "", "Set the OS of the imported image")
	image_importCmd.Flags().BoolP("quiet", "q", false, "Suppress output")
	image_importCmd.Flags().String("signature-policy", "", "Path to a signature-policy file")
	image_importCmd.Flags().String("variant", "", "Set the variant of the imported image")
	image_importCmd.Flag("signature-policy").Hidden = true
	imageCmd.AddCommand(image_importCmd)
}
