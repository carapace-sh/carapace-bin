package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var runtimeInfoCmd = &cobra.Command{
	Use:    "runtime-info",
	Short:  "Show runtime information",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runtimeInfoCmd).Standalone()

	rootCmd.AddCommand(runtimeInfoCmd)
}
