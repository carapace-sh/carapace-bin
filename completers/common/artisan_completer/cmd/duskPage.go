package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var duskPageCmd = &cobra.Command{
	Use:   "dusk:page <name>",
	Short: "Create a new Dusk page class",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(duskPageCmd).Standalone()

	rootCmd.AddCommand(duskPageCmd)
}
