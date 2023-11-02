package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var util_mangenCmd = &cobra.Command{
	Use:   "mangen",
	Short: "Print a ROFF (manpage)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(util_mangenCmd).Standalone()

	util_mangenCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	utilCmd.AddCommand(util_mangenCmd)
}
