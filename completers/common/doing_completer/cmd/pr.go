package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var prCmd = &cobra.Command{
	Use:   "pr [OPTIONS] COMMAND [ARGS]...",
	Short: "Work with pull requests",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(prCmd).Standalone()

	prCmd.Flags().Bool("help", false, "Show this message and exit.")
	rootCmd.AddCommand(prCmd)
}
