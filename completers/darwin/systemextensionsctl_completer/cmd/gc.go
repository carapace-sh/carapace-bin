package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gcCmd = &cobra.Command{
	Use:   "gc",
	Short: "Garbage collect orphaned system extensions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gcCmd).Standalone()
	rootCmd.AddCommand(gcCmd)
}
