package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ci_cancelCmd = &cobra.Command{
	Use:   "cancel <command>",
	Short: "Cancel a running pipeline or job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ci_cancelCmd).Standalone()

	ciCmd.AddCommand(ci_cancelCmd)
}
