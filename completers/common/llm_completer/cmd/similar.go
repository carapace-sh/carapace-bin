package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var similarCmd = &cobra.Command{
	Use:   "similar",
	Short: "Return top N similar IDs from a collection using cosine similarity.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(similarCmd).Standalone()

	similarCmd.Flags().Bool("binary", false, "Treat input as binary data")
	similarCmd.Flags().StringP("content", "c", "", "Content to embed for comparison")
	similarCmd.Flags().StringP("database", "d", "", "")
	similarCmd.Flags().StringP("input", "i", "", "File to embed for comparison")
	similarCmd.Flags().StringP("number", "n", "", "Number of results to return")
	similarCmd.Flags().BoolP("plain", "p", false, "Output in plain text format")
	similarCmd.Flags().String("prefix", "", "Just IDs with this prefix")
	rootCmd.AddCommand(similarCmd)
}
