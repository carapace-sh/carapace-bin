package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var normalizeCmd = &cobra.Command{
	Use:   "normalize",
	Short: "Normalize paths",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(normalizeCmd).Standalone()

	normalizeCmd.Flags().BoolS("Z", "Z", false, "NUL-delimited output")
	normalizeCmd.Flags().Bool("null-in", false, "NUL-delimited input")
	normalizeCmd.Flags().Bool("null-out", false, "NUL-delimited output")
	normalizeCmd.Flags().BoolS("q", "q", false, "suppress output")
	normalizeCmd.Flags().Bool("quiet", false, "suppress output")
	normalizeCmd.Flags().BoolS("z", "z", false, "NUL-delimited input")

	rootCmd.AddCommand(normalizeCmd)
}
