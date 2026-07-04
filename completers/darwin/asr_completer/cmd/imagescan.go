package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagescanCmd = &cobra.Command{
	Use:   "imagescan",
	Short: "Scan an image for ASR",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagescanCmd).Standalone()
	rootCmd.AddCommand(imagescanCmd)
	carapace.Gen(imagescanCmd).PositionalCompletion(carapace.ActionFiles())
}
