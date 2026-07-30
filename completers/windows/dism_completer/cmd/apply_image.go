package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ApplyImageCmd = &cobra.Command{
	Use:   "Apply-Image",
	Short: "apply an image to a directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ApplyImageCmd).Standalone()
	rootCmd.AddCommand(ApplyImageCmd)
}
