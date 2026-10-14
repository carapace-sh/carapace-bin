package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var volume_unsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Unset volume properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_unsetCmd).Standalone()

	volume_unsetCmd.Flags().String("image-property", "", "Remove an image property from volume (repeat option to remove multiple image properties)")
	volume_unsetCmd.Flags().String("property", "", "Remove a property from volume (repeat option to remove multiple properties)")
	volumeCmd.AddCommand(volume_unsetCmd)
}
