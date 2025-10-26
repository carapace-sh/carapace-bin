package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opentofu_state_downloadCmd = &cobra.Command{
	Use:   "download <state> [<serial>]",
	Short: "Download the given state and output as JSON to stdout.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opentofu_state_downloadCmd).Standalone()

	opentofu_stateCmd.AddCommand(opentofu_state_downloadCmd)
}
