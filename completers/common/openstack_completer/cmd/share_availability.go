package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_availabilityCmd = &cobra.Command{
	Use:   "availability",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_availabilityCmd).Standalone()

	shareCmd.AddCommand(share_availabilityCmd)
}
