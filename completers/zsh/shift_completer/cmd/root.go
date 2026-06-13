package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/number"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "shift",
	Short: "Shift positional parameters",
	Long:  "https://zsh.sourceforge.io/Doc/Release/Shell-Builtin-Commands.html#index-shift",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalCompletion(
		number.ActionRange(number.RangeOpts{Start: 1, End: 9}).
			Usage("number of positions to shift"),
	)
}
