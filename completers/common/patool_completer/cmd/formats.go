package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var formatsCmd = &cobra.Command{
	Use:   "formats",
	Short: "show supported archive formats",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(formatsCmd).Standalone()

	rootCmd.AddCommand(formatsCmd)
}
