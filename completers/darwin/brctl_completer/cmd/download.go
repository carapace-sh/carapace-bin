package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(downloadCmd)
}

var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "download a local copy of the document",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(downloadCmd).Standalone()
}