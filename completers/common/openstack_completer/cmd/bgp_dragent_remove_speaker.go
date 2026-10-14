package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bgp_dragent_remove_speakerCmd = &cobra.Command{
	Use:   "speaker",
	Short: "Removes a BGP speaker from a dynamic routing agent",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bgp_dragent_remove_speakerCmd).Standalone()

	bgp_dragent_removeCmd.AddCommand(bgp_dragent_remove_speakerCmd)
}
