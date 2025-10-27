package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var image_loadCmd = &cobra.Command{
	Use:   "load [OPTIONS]",
	Short: "Load an image from a tar archive or STDIN",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_loadCmd).Standalone()

	image_loadCmd.Flags().StringP("input", "i", "", "Read from tar archive file, instead of STDIN")
	image_loadCmd.Flags().String("platform", "", "Load only the given platform variant. Formatted as \"os[/arch[/variant]]\" (e.g., \"linux/amd64\")")
	image_loadCmd.Flags().BoolP("quiet", "q", false, "Suppress the load output")
	imageCmd.AddCommand(image_loadCmd)

	carapace.Gen(image_loadCmd).FlagCompletion(carapace.ActionMap{
		"input": carapace.ActionFiles(),
	})
}
