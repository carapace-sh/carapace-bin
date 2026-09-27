package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bgp_dragent_add_speakerCmd = &cobra.Command{
	Use:   "speaker",
	Short: "Add a BGP speaker to a dynamic routing agent",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bgp_dragent_add_speakerCmd).Standalone()

	bgp_dragent_addCmd.AddCommand(bgp_dragent_add_speakerCmd)
}
