package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var image_saveCmd = &cobra.Command{
	Use:   "save [options] IMAGE [IMAGE...]",
	Short: "Save image(s) to an archive",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_saveCmd).Standalone()

	image_saveCmd.Flags().Bool("compress", false, "Compress tarball image layers when saving to a directory using the 'dir' transport. (default is same compression type as source)")
	image_saveCmd.Flags().String("format", "", "Save image to oci-archive, oci-dir (directory with oci manifest type), docker-archive, docker-dir (directory with v2s2 manifest type)")
	image_saveCmd.Flags().BoolP("multi-image-archive", "m", false, "Interpret additional arguments as images not tags and create a multi-image-archive (only for docker-archive)")
	image_saveCmd.Flags().StringP("output", "o", "", "Write to a specified file (default: stdout, which must be redirected)")
	image_saveCmd.Flags().BoolP("quiet", "q", false, "Suppress the output")
	image_saveCmd.Flags().String("signature-policy", "", "Path to a signature-policy file")
	image_saveCmd.Flags().Bool("uncompressed", false, "Accept uncompressed layers when copying OCI images")
	image_saveCmd.Flag("signature-policy").Hidden = true
	imageCmd.AddCommand(image_saveCmd)
}
