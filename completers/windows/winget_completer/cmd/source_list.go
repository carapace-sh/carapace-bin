package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var source_listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List current sources",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(source_listCmd).Standalone()

	source_listCmd.Flags().StringP("name", "n", "", "Name of the source")
	sourceCmd.AddCommand(source_listCmd)

	// TODO flag completion
}
