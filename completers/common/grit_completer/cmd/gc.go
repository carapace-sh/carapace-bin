package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gcCmd = &cobra.Command{
	Use:   "gc",
	Short: "Pack refs, expire reflogs, repack, prune, and refresh commit-graph",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gcCmd).Standalone()

	gcCmd.Flags().Bool("dry-run", false, "Show what would run without writing")
	gcCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	gcCmd.Flags().BoolP("quiet", "q", false, "Suppress progress summary on stderr")
	rootCmd.AddCommand(gcCmd)
}
