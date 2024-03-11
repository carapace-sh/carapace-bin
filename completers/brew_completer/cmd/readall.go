package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var readallCmd = &cobra.Command{
	Use:     "readall",
	Short:   "Import all items from the specified <tap>, or from all installed taps if none is provided",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(readallCmd).Standalone()

	readallCmd.Flags().Bool("aliases", false, "Verify any alias symlinks in each tap.")
	readallCmd.Flags().Bool("arch", false, "Read using the given CPU architecture. (Pass `all` to simulate all architectures.)")
	readallCmd.Flags().Bool("debug", false, "Display any debugging information.")
	readallCmd.Flags().Bool("eval-all", false, "Evaluate all available formulae and casks, whether installed or not. Implied if `HOMEBREW_EVAL_ALL` is set.")
	readallCmd.Flags().Bool("help", false, "Show this message.")
	readallCmd.Flags().Bool("no-simulate", false, "Don't simulate other system configurations when checking formulae and casks.")
	readallCmd.Flags().Bool("os", false, "Read using the given operating system. (Pass `all` to simulate all operating systems.)")
	readallCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	readallCmd.Flags().Bool("syntax", false, "Syntax-check all of Homebrew's Ruby files (if no <tap> is passed).")
	readallCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(readallCmd)
}
