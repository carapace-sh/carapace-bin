package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var image_memberCmd = &cobra.Command{
	Use:   "member",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_memberCmd).Standalone()

	imageCmd.AddCommand(image_memberCmd)
}
