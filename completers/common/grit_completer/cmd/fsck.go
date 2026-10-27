package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fsckCmd = &cobra.Command{
	Use:     "fsck",
	Short:   "Check refs, objects, and pack indexes for problems",
	Aliases: []string{"verify"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fsckCmd).Standalone()

	fsckCmd.Flags().Bool("connectivity-only", false, "Check connectivity from refs only (skip store-wide object validation)")
	fsckCmd.Flags().Bool("dangling", false, "Report unreachable objects (with `--connectivity-only`, off unless this flag is set)")
	fsckCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	fsckCmd.Flags().Bool("no-dangling", false, "Do not list dangling objects (matches Git `--no-dangling`)")
	fsckCmd.Flags().Bool("strict", false, "Treat ref warnings as errors")
	fsckCmd.Flags().BoolP("verbose", "v", false, "Include info-level object diagnostics in human output (matches Git `--verbose`)")
	rootCmd.AddCommand(fsckCmd)
}
