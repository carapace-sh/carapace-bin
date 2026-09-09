package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var volume_groupCmd = &cobra.Command{
	Use:   "group",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_groupCmd).Standalone()

	volumeCmd.AddCommand(volume_groupCmd)
}
