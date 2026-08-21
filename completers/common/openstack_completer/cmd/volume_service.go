package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var volume_serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_serviceCmd).Standalone()

	volumeCmd.AddCommand(volume_serviceCmd)
}
