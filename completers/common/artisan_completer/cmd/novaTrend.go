package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var novaTrendCmd = &cobra.Command{
	Use:   "nova:trend <name>",
	Short: "Create a new metric (trend) class",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(novaTrendCmd).Standalone()

	rootCmd.AddCommand(novaTrendCmd)
}
