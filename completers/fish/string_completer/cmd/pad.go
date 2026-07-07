package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var padCmd = &cobra.Command{
	Use:   "pad",
	Short: "Pad strings to a visible width",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(padCmd).Standalone()

	padCmd.Flags().BoolS("C", "C", false, "center the string")
	padCmd.Flags().Bool("center", false, "center the string")
	padCmd.Flags().StringP("char", "c", "", "character to pad with")
	padCmd.Flags().BoolP("right", "r", false, "add padding after the string")
	padCmd.Flags().StringP("width", "w", "", "minimum width to pad to")

	rootCmd.AddCommand(padCmd)
}
