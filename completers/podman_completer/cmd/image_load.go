package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var image_loadCmd = &cobra.Command{
	Use:   "load [options]",
	Short: "Load image(s) from a tar archive",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_loadCmd).Standalone()

	image_loadCmd.Flags().StringP("input", "i", "", "Read from specified archive file (default: stdin)")
	image_loadCmd.Flags().BoolP("quiet", "q", false, "Suppress the output")
	image_loadCmd.Flags().String("signature-policy", "", "Pathname of signature policy file")
	image_loadCmd.Flag("signature-policy").Hidden = true
	imageCmd.AddCommand(image_loadCmd)
}
