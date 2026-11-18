package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/brew_completer/cmd/action"
	"github.com/spf13/cobra"
)

var pinCmd = &cobra.Command{
	Use:     "pin",
	Short:   "Pin the specified <formula>, preventing them from being upgraded when issuing the `brew upgrade` <formula> command",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinCmd).Standalone()

	pinCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	pinCmd.Flags().BoolP("help", "h", false, "Show this message.")
	pinCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	pinCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
	rootCmd.AddCommand(pinCmd)

	carapace.Gen(pinCmd).PositionalAnyCompletion(
		action.ActionList(pinCmd).FilterArgs(),
	)
}
