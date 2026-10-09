package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var volume_backend_capabilityCmd = &cobra.Command{
	Use:   "capability",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_backend_capabilityCmd).Standalone()

	volume_backendCmd.AddCommand(volume_backend_capabilityCmd)
}
