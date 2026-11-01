package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "arch",
	Short: "print architecture type or run a command with a different architecture",
	Long:  "https://keith.github.io/xcode-manpages/arch.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("32", "32", false, "Add the native 32-bit architecture to the list of architectures")
	rootCmd.Flags().BoolS("64", "64", false, "Add the native 64-bit architecture to the list of architectures")
	rootCmd.Flags().BoolS("c", "c", false, "Clear the environment that will be passed to the command")
	rootCmd.Flags().StringS("d", "d", "", "Delete the named environment variable from the environment")
	rootCmd.Flags().StringS("e", "e", "", "Assign the given value to the named environment variable")
	rootCmd.Flags().BoolS("h", "h", false, "Print a usage message and exit")

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionValues(),
	)
}
