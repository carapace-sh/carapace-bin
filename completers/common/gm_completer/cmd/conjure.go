package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var conjureCmd = &cobra.Command{
	Use:   "conjure",
	Short: "execute a Magick Scripting Language (MSL) XML script",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(conjureCmd).Standalone()

	conjureCmd.Flags().StringS("debug", "debug", "", "display copious debugging information")
	conjureCmd.Flags().BoolS("help", "help", false, "print program options")
	conjureCmd.Flags().StringS("log", "log", "", "format of debugging information")
	conjureCmd.Flags().BoolS("verbose", "verbose", false, "print detailed information about the image")
	conjureCmd.Flags().BoolS("version", "version", false, "print version information")
	rootCmd.AddCommand(conjureCmd)

	carapace.Gen(conjureCmd).PositionalCompletion(carapace.ActionFiles())
}
