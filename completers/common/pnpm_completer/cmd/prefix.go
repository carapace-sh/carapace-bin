package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var prefixCmd = &cobra.Command{
	Use:   "prefix",
	Short: "Print the current package prefix",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(prefixCmd).Standalone()

	prefixCmd.Flags().BoolP("global", "g", false, "Print the global prefix")
	prefixCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(prefixCmd)
}
