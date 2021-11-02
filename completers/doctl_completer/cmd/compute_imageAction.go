package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_imageActionCmd = &cobra.Command{
	Use:   "image-action",
	Short: "Display commands to perform actions on images",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_imageActionCmd).Standalone()
	computeCmd.AddCommand(compute_imageActionCmd)
}
