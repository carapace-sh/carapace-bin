package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var GetImageInfoCmd = &cobra.Command{
	Use:   "Get-ImageInfo",
	Short: "display information about images in a file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(GetImageInfoCmd).Standalone()
	rootCmd.AddCommand(GetImageInfoCmd)
}
