package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var util_help_mangenCmd = &cobra.Command{
	Use:   "mangen",
	Short: "Print a ROFF (manpage)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(util_help_mangenCmd).Standalone()

	util_helpCmd.AddCommand(util_help_mangenCmd)
}
