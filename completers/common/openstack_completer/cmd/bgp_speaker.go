package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bgp_speakerCmd = &cobra.Command{
	Use:   "speaker",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bgp_speakerCmd).Standalone()

	bgpCmd.AddCommand(bgp_speakerCmd)
}
