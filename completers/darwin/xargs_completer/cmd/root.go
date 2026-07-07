package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "xargs",
	Short: "construct argument list(s) and execute utility",
	Long:  "https://keith.github.io/xcode-manpages/xargs.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.Flags().SetInterspersed(false)

	rootCmd.Flags().BoolS("0", "0", false, "Input items are terminated by NUL")
	rootCmd.Flags().StringS("E", "E", "", "Set EOF string")
	rootCmd.Flags().StringS("I", "I", "", "Replace replstr in arguments")
	rootCmd.Flags().StringS("J", "J", "", "Replace replstr in initial arguments")
	rootCmd.Flags().StringS("L", "L", "", "Use at most number lines per command")
	rootCmd.Flags().StringS("P", "P", "", "Run up to maxprocs in parallel")
	rootCmd.Flags().StringS("R", "R", "", "Number of replacements")
	rootCmd.Flags().StringS("S", "S", "", "Replacement size")
	rootCmd.Flags().StringS("n", "n", "", "Use at most number arguments per command")
	rootCmd.Flags().BoolS("o", "o", false, "Reopen stdin as /dev/tty")
	rootCmd.Flags().BoolS("p", "p", false, "Prompt before each command")
	rootCmd.Flags().BoolS("r", "r", false, "Do not run if no input")
	rootCmd.Flags().StringS("s", "s", "", "Maximum command line size")
	rootCmd.Flags().BoolS("t", "t", false, "Print command before executing")

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin(),
	)
}
