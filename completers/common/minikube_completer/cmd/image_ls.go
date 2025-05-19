package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var image_lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List images",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_lsCmd).Standalone()
	imageCmd.AddCommand(image_lsCmd)
}
