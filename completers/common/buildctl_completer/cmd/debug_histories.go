package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_historiesCmd = &cobra.Command{
	Use:   "histories",
	Short: "list build records",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_historiesCmd).Standalone()

	debug_historiesCmd.Flags().String("format", "", "Format the output using the given Go template")
	debugCmd.AddCommand(debug_historiesCmd)
}
