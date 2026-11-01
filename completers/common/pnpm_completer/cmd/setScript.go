package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var setScriptCmd = &cobra.Command{
	Use:     "set-script",
	Short:   "Set a script in package.json",
	Aliases: []string{"ss"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setScriptCmd).Standalone()

	setScriptCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(setScriptCmd)
}
