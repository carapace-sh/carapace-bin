package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var image_taskCmd = &cobra.Command{
	Use:   "task",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_taskCmd).Standalone()

	imageCmd.AddCommand(image_taskCmd)
}
