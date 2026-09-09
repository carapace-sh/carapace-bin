package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var starCmd = &cobra.Command{
	Use:   "star",
	Short: "Marks a package as a favorite",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(starCmd).Standalone()

	starCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(starCmd)
}
