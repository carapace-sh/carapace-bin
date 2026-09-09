package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var volume_typeCmd = &cobra.Command{
	Use:   "type",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_typeCmd).Standalone()

	volumeCmd.AddCommand(volume_typeCmd)
}
