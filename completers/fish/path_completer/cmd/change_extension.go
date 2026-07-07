package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var changeExtensionCmd = &cobra.Command{
	Use:   "change-extension",
	Short: "Change or strip the file extension",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(changeExtensionCmd).Standalone()

	changeExtensionCmd.Flags().BoolS("Z", "Z", false, "NUL-delimited output")
	changeExtensionCmd.Flags().Bool("null-in", false, "NUL-delimited input")
	changeExtensionCmd.Flags().Bool("null-out", false, "NUL-delimited output")
	changeExtensionCmd.Flags().BoolS("q", "q", false, "suppress output")
	changeExtensionCmd.Flags().Bool("quiet", false, "suppress output")
	changeExtensionCmd.Flags().BoolS("z", "z", false, "NUL-delimited input")

	rootCmd.AddCommand(changeExtensionCmd)
}
