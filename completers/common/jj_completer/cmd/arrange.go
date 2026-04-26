package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var arrangeCmd = &cobra.Command{
	Use:   "arrange",
	Short: "Interactively arrange the commit graph",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(arrangeCmd).Standalone()

	arrangeCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(arrangeCmd)
}
