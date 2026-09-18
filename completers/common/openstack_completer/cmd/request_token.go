package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var request_tokenCmd = &cobra.Command{
	Use:   "token",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(request_tokenCmd).Standalone()

	requestCmd.AddCommand(request_tokenCmd)
}
