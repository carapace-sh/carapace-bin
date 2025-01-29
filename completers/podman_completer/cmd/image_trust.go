package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var image_trustCmd = &cobra.Command{
	Use:   "trust",
	Short: "Manage container image trust policy",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_trustCmd).Standalone()

	imageCmd.AddCommand(image_trustCmd)
}
