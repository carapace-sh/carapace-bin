package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var volume_group_typeCmd = &cobra.Command{
	Use:   "type",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_group_typeCmd).Standalone()

	volume_groupCmd.AddCommand(volume_group_typeCmd)
}
