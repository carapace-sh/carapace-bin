package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mallocCmd = &cobra.Command{
	Use:   "malloc",
	Short: "Dump malloc stats of a D-Bus service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mallocCmd).Standalone()

	rootCmd.AddCommand(mallocCmd)
}
