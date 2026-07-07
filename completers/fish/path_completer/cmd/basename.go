package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var basenameCmd = &cobra.Command{
	Use:   "basename",
	Short: "Return the last path component",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(basenameCmd).Standalone()

	basenameCmd.Flags().BoolS("E", "E", false, "strip the extension")
	basenameCmd.Flags().BoolS("Z", "Z", false, "NUL-delimited output")
	basenameCmd.Flags().Bool("no-extension", false, "strip the extension")
	basenameCmd.Flags().Bool("null-in", false, "NUL-delimited input")
	basenameCmd.Flags().Bool("null-out", false, "NUL-delimited output")
	basenameCmd.Flags().BoolS("q", "q", false, "suppress output")
	basenameCmd.Flags().Bool("quiet", false, "suppress output")
	basenameCmd.Flags().BoolS("z", "z", false, "NUL-delimited input")

	rootCmd.AddCommand(basenameCmd)
}
