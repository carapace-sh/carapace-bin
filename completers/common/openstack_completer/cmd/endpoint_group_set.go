package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var endpoint_group_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set endpoint group properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(endpoint_group_setCmd).Standalone()

	endpoint_group_setCmd.Flags().String("description", "", "New endpoint group description")
	endpoint_group_setCmd.Flags().String("filters", "", "Filename that contains a new set of filters")
	endpoint_group_setCmd.Flags().String("name", "", "New endpoint group name")
	endpoint_groupCmd.AddCommand(endpoint_group_setCmd)
}
