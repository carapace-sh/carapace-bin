package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var image_pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pull images",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_pullCmd).Standalone()

	imageCmd.AddCommand(image_pullCmd)
}
