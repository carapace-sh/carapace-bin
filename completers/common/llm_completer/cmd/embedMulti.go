package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var embedMultiCmd = &cobra.Command{
	Use:   "embed-multi",
	Short: "Store embeddings for multiple strings at once in the specified collection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(embedMultiCmd).Standalone()

	embedMultiCmd.Flags().String("attach", "", "Additional databases to attach - specify alias and file path")
	embedMultiCmd.Flags().String("batch-size", "", "Batch size to use when running embeddings")
	embedMultiCmd.Flags().Bool("binary", false, "Treat --files as binary data")
	embedMultiCmd.Flags().StringP("database", "d", "", "")
	embedMultiCmd.Flags().String("encoding", "", "Encodings to try when reading --files")
	embedMultiCmd.Flags().String("files", "", "Embed files in this directory - specify directory and glob pattern")
	embedMultiCmd.Flags().String("format", "", "Format of input file - defaults to auto-detect")
	embedMultiCmd.Flags().String("key", "", "API key to use")
	embedMultiCmd.Flags().StringP("model", "m", "", "Embedding model to use")
	embedMultiCmd.Flags().String("prefix", "", "Prefix to add to the IDs")
	embedMultiCmd.Flags().String("prepend", "", "Prepend this string to all content before embedding")
	embedMultiCmd.Flags().String("sql", "", "Read input using this SQL query")
	embedMultiCmd.Flags().Bool("store", false, "Store the text itself in the database")
	rootCmd.AddCommand(embedMultiCmd)
}
