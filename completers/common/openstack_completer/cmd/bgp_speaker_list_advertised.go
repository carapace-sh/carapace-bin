package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bgp_speaker_list_advertisedCmd = &cobra.Command{
	Use:   "advertised",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bgp_speaker_list_advertisedCmd).Standalone()

	bgp_speaker_listCmd.AddCommand(bgp_speaker_list_advertisedCmd)
}
