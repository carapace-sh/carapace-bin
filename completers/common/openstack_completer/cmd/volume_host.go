package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var volume_hostCmd = &cobra.Command{
	Use:   "host",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_hostCmd).Standalone()

	volumeCmd.AddCommand(volume_hostCmd)
}
