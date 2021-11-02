package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_imageCmd = &cobra.Command{
	Use:   "image",
	Short: "Display commands to manage images",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_imageCmd).Standalone()
	computeCmd.AddCommand(compute_imageCmd)
}
