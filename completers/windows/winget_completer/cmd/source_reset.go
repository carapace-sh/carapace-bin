package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var source_resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset sources",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(source_resetCmd).Standalone()

	source_resetCmd.Flags().StringP("name", "n", "", "Name of the source")
	sourceCmd.AddCommand(source_resetCmd)

	// TODO flag completion
}
