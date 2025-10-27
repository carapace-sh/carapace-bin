package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var image_pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push images",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_pushCmd).Standalone()

	imageCmd.AddCommand(image_pushCmd)
}
