package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_infoCmd = &cobra.Command{
	Use:   "info",
	Short: "display internal information",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_infoCmd).Standalone()

	debug_infoCmd.Flags().String("format", "", "Format the output using the given Go template")
	debugCmd.AddCommand(debug_infoCmd)
}
