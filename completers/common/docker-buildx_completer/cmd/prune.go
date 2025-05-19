package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pruneCmd = &cobra.Command{
	Use:   "prune",
	Short: "Remove build cache",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pruneCmd).Standalone()
	pruneCmd.Flags().BoolP("all", "a", false, "Include internal/frontend images")
	pruneCmd.Flags().String("filter", "", "Provide filter values (e.g., \"until=24h\")")
	pruneCmd.Flags().BoolP("force", "f", false, "Do not prompt for confirmation")
	pruneCmd.Flags().String("keep-storage", "", "Amount of disk space to keep for cache")
	pruneCmd.Flags().Bool("verbose", false, "Provide a more verbose output")
	rootCmd.AddCommand(pruneCmd)
}
