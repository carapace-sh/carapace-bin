package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var file_obliterateCmd = &cobra.Command{
	Use:   "obliterate",
	Short: "Obliterate a file or fragment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(file_obliterateCmd).Standalone()

	file_obliterateCmd.Flags().String("address", "", "Address of a blob")
	file_obliterateCmd.Flags().BoolP("help", "h", false, "Print help")
	file_obliterateCmd.Flags().String("path", "", "Path to a file")
	fileCmd.AddCommand(file_obliterateCmd)

	carapace.Gen(file_obliterateCmd).FlagCompletion(carapace.ActionMap{
		"address": carapace.ActionValues(),
		"path":    carapace.ActionFiles(),
	})
}
