package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var generate_debugCmd = &cobra.Command{
	Use:    "debug",
	Short:  "",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(generate_debugCmd).Standalone()

	generateCmd.AddCommand(generate_debugCmd)
}
