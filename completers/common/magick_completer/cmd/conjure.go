package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/magick_completer/cmd/action"
	"github.com/spf13/cobra"
)

var conjureCmd = &cobra.Command{
	Use:   "conjure",
	Short: "interpret and execute scripts written in the Magick Scripting Language (MSL).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(conjureCmd).Standalone()

	conjureCmd.Flags().StringSliceS("debug", "debug", nil, "display copious debugging information")
	conjureCmd.Flags().BoolS("help", "help", false, "print program options")
	conjureCmd.Flags().StringSliceS("log", "log", nil, "format of debugging information")
	conjureCmd.Flags().CountS("monitor", "monitor", "monitor progress")
	conjureCmd.Flags().CountS("quiet", "quiet", "suppress all warning messages")
	conjureCmd.Flags().CountS("regard-warnings", "regard-warnings", "pay attention to warning messages.")
	conjureCmd.Flags().StringSliceS("seed", "seed", nil, "seed a new sequence of pseudo-random numbers")
	conjureCmd.Flags().BoolS("verbose", "verbose", false, "print detailed information about the image")
	conjureCmd.Flags().BoolS("version", "version", false, "print version information")
	rootCmd.AddCommand(conjureCmd)

	carapace.Gen(conjureCmd).FlagCompletion(carapace.ActionMap{
		"debug": action.ActionCompleteOption("Debug").UniqueList(","),
	})

	carapace.Gen(conjureCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
