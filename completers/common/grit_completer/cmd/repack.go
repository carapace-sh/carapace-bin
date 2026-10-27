package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repackCmd = &cobra.Command{
	Use:   "repack",
	Short: "Pack reachable objects and remove redundant packs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repackCmd).Standalone()

	repackCmd.Flags().Bool("dry-run", false, "Show what would be packed without writing")
	repackCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	repackCmd.Flags().BoolP("quiet", "q", false, "Suppress progress summary on stderr")
	rootCmd.AddCommand(repackCmd)
}
