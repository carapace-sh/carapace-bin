package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_canonicalizeFlagsCmd = &cobra.Command{
	Use:   "canonicalize-flags",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_canonicalizeFlagsCmd).Standalone()

	helpCmd.AddCommand(help_canonicalizeFlagsCmd)
}
