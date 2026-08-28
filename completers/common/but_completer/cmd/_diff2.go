package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var _diff2Cmd = &cobra.Command{
	Use:    "_diff2",
	Short:  "Displays the diff of changes in the repo",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(_diff2Cmd).Standalone()

	_diff2Cmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(_diff2Cmd)
}