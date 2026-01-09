package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var source_removeCmd = &cobra.Command{
	Use:     "remove",
	Aliases: []string{"rm"},
	Short:   "Remove current sources",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(source_removeCmd).Standalone()

	source_removeCmd.Flags().StringP("name", "n", "", "Name of the source")
	sourceCmd.AddCommand(source_removeCmd)

	// TODO flag completion
}
