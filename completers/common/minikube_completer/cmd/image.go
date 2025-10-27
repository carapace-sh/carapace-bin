package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imageCmd = &cobra.Command{
	Use:     "image COMMAND",
	Short:   "Manage images",
	GroupID: "images",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imageCmd).Standalone()

	rootCmd.AddCommand(imageCmd)
}
