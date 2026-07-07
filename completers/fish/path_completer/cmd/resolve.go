package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resolveCmd = &cobra.Command{
	Use:   "resolve",
	Short: "Resolve paths to absolute physical paths",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resolveCmd).Standalone()

	resolveCmd.Flags().BoolS("Z", "Z", false, "NUL-delimited output")
	resolveCmd.Flags().Bool("null-in", false, "NUL-delimited input")
	resolveCmd.Flags().Bool("null-out", false, "NUL-delimited output")
	resolveCmd.Flags().BoolS("q", "q", false, "suppress output")
	resolveCmd.Flags().Bool("quiet", false, "suppress output")
	resolveCmd.Flags().BoolS("z", "z", false, "NUL-delimited input")

	rootCmd.AddCommand(resolveCmd)
}
