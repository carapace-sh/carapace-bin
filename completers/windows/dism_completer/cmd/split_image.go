package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var SplitImageCmd = &cobra.Command{
	Use:   "Split-Image",
	Short: "split an image file into smaller files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(SplitImageCmd).Standalone()
	rootCmd.AddCommand(SplitImageCmd)
}
