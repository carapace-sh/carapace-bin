package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Open the project in GitButler",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(openCmd).Standalone()

	openCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	openCmd.Flags().Bool("print", false, "Print the link instead of opening it")
	rootCmd.AddCommand(openCmd)
}
