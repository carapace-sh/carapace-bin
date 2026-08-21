package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var registered_limitCmd = &cobra.Command{
	Use:   "limit",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(registered_limitCmd).Standalone()

	registeredCmd.AddCommand(registered_limitCmd)
}
