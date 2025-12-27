package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspectElfCmd = &cobra.Command{
	Use:   "inspect-elf",
	Short: "Parse and print ELF package metadata",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspectElfCmd).Standalone()

	rootCmd.AddCommand(inspectElfCmd)

	carapace.Gen(inspectElfCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
