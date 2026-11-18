package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var trustCmd = &cobra.Command{
	Use:     "trust",
	Short:   "Trust non-official tap formulae, casks or commands so Homebrew may load them when `$HOMEBREW_REQUIRE_TAP_TRUST` is set",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(trustCmd).Standalone()

	trustCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	trustCmd.Flags().BoolP("help", "h", false, "Show this message.")
	trustCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	trustCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
	rootCmd.AddCommand(trustCmd)
}
