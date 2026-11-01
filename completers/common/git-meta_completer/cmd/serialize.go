package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var serializeCmd = &cobra.Command{
	Use:   "serialize",
	Short: "Serialize metadata to Git ref",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serializeCmd).Standalone()

	serializeCmd.Flags().Bool("force-full", false, "Rebuild serialized refs from the full SQLite state")
	serializeCmd.Flags().BoolP("help", "h", false, "Print help")
	serializeCmd.Flags().BoolP("verbose", "v", false, "Show detailed information about serialization decisions")
	rootCmd.AddCommand(serializeCmd)
}
