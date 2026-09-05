package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var unstarCmd = &cobra.Command{
	Use:   "unstar",
	Short: "Unmarks a package as a favorite",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unstarCmd).Standalone()

	unstarCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(unstarCmd)
}
