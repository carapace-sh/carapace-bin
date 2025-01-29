package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var image_inspectCmd = &cobra.Command{
	Use:   "inspect [options] IMAGE [IMAGE...]",
	Short: "Display the configuration of an image",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_inspectCmd).Standalone()

	image_inspectCmd.Flags().StringP("format", "f", "", "Format the output to a Go template or json")
	imageCmd.AddCommand(image_inspectCmd)
}
