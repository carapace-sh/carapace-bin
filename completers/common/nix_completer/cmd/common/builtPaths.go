package common

import (
	"github.com/spf13/cobra"
)

// AddBuiltPathsFlags adds flags for commands operating on store paths
// (BuiltPathsCommand in nix source).
func AddBuiltPathsFlags(cmd *cobra.Command) {
	cmd.Flags().Bool("all", false, "Apply the operation to every store path")
	cmd.Flags().Bool("derivation", false, "Operate on the store derivation rather than its outputs")
	cmd.Flags().BoolP("recursive", "r", false, "Apply operation to closure of the specified paths")
}
