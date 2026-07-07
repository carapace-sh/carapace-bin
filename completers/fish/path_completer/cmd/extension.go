package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var extensionCmd = &cobra.Command{
	Use:   "extension",
	Short: "Return the file extension",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(extensionCmd).Standalone()

	extensionCmd.Flags().BoolS("Z", "Z", false, "NUL-delimited output")
	extensionCmd.Flags().Bool("null-in", false, "NUL-delimited input")
	extensionCmd.Flags().Bool("null-out", false, "NUL-delimited output")
	extensionCmd.Flags().BoolS("q", "q", false, "suppress output")
	extensionCmd.Flags().Bool("quiet", false, "suppress output")
	extensionCmd.Flags().BoolS("z", "z", false, "NUL-delimited input")

	rootCmd.AddCommand(extensionCmd)
}
