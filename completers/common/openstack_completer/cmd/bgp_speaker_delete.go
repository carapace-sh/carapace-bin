package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bgp_speaker_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a BGP speaker",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bgp_speaker_deleteCmd).Standalone()

	bgp_speakerCmd.AddCommand(bgp_speaker_deleteCmd)
}
