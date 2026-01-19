package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pruneCmd = &cobra.Command{
	Use:   "prune",
	Short: "clean up build cache",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pruneCmd).Standalone()

	pruneCmd.Flags().Bool("all", false, "Include internal/frontend references")
	pruneCmd.Flags().StringP("filter", "f", "", "Filter records")
	pruneCmd.Flags().String("format", "", "Format the output using the given Go template")
	pruneCmd.Flags().String("free-storage", "", "Keep free data below this limit")
	pruneCmd.Flags().String("keep-duration", "", "Keep data newer than this limit")
	pruneCmd.Flags().String("keep-storage", "", "Keep data below this limit")
	pruneCmd.Flags().String("keep-storage-min", "", "Always allow data above this limit")
	pruneCmd.Flags().BoolP("verbose", "v", false, "Verbose output")
	rootCmd.AddCommand(pruneCmd)
}
