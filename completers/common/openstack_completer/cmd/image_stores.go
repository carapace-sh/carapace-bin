package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var image_storesCmd = &cobra.Command{
	Use:   "stores",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_storesCmd).Standalone()

	imageCmd.AddCommand(image_storesCmd)
}
