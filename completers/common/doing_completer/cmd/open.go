package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var openCmd = &cobra.Command{
	Use:   "open [OPTIONS] COMMAND [ARGS]...",
	Short: "Quickly open certain links",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(openCmd).Standalone()

	openCmd.Flags().Bool("help", false, "Show this message and exit.")
	rootCmd.AddCommand(openCmd)
}
