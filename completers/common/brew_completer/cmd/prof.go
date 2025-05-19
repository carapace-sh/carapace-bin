package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var profCmd = &cobra.Command{
	Use:     "prof",
	Short:   "Run Homebrew with a Ruby profiler",
	GroupID: "developer",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(profCmd).Standalone()

	profCmd.Flags().Bool("debug", false, "Display any debugging information.")
	profCmd.Flags().Bool("help", false, "Show this message.")
	profCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	profCmd.Flags().Bool("stackprof", false, "Use `stackprof` instead of `ruby-prof` (the default).")
	profCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(profCmd)
}
