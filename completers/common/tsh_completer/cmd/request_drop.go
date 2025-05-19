package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var request_dropCmd = &cobra.Command{
	Use:   "drop",
	Short: "Drop one more access requests from current identity.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(request_dropCmd).Standalone()

	requestCmd.AddCommand(request_dropCmd)
}
