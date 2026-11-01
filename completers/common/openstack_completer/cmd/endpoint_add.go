package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var endpoint_addCmd = &cobra.Command{
	Use:   "add",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(endpoint_addCmd).Standalone()

	endpointCmd.AddCommand(endpoint_addCmd)
}
