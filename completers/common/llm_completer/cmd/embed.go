package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var embedCmd = &cobra.Command{
	Use:   "embed",
	Short: "Embed text and store or return the result",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(embedCmd).Standalone()

	embedCmd.Flags().Bool("binary", false, "Treat input as binary data")
	embedCmd.Flags().StringP("content", "c", "", "Content to embed")
	embedCmd.Flags().StringP("database", "d", "", "")
	embedCmd.Flags().StringP("format", "f", "", "Output format")
	embedCmd.Flags().StringP("input", "i", "", "File to embed")
	embedCmd.Flags().String("key", "", "API key to use")
	embedCmd.Flags().String("metadata", "", "JSON object metadata to store")
	embedCmd.Flags().StringP("model", "m", "", "Embedding model to use")
	embedCmd.Flags().Bool("store", false, "Store the text itself in the database")
	rootCmd.AddCommand(embedCmd)
}
