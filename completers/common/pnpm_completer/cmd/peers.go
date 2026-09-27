package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var peersCmd = &cobra.Command{
	Use:   "peers",
	Short: "Checks for unmet or missing peer dependency issues",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(peersCmd).Standalone()

	peersCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	peersCmd.Flags().Bool("json", false, "")
	peersCmd.Flags().Bool("lockfile-only", false, "")
	rootCmd.AddCommand(peersCmd)
}
