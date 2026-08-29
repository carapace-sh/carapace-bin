package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var _expandCmd = &cobra.Command{
	Use:    "_expand",
	Short:  "Debug command for expanding a CLI ID into any matching resources",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(_expandCmd).Standalone()

	_expandCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(_expandCmd)
}
