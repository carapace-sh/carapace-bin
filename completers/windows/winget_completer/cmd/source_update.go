package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var source_updateCmd = &cobra.Command{
	Use:     "update",
	Aliases: []string{"refresh"},
	Short:   "Update current sources",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(source_updateCmd).Standalone()

	source_updateCmd.Flags().StringP("name", "n", "", "Name of the source")
	sourceCmd.AddCommand(source_updateCmd)

	// TODO flag completion
}
