package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var help_util_mangenCmd = &cobra.Command{
	Use:   "mangen",
	Short: "Print a ROFF (manpage)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_util_mangenCmd).Standalone()

	help_utilCmd.AddCommand(help_util_mangenCmd)
}
