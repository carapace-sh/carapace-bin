package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_packages_brew_untapCmd = &cobra.Command{
	Use:     "untap",
	Short:   "Remove Homebrew tap URLs from `[bootstrap.brew.taps]`",
	Aliases: []string{"remove", "rm"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_packages_brew_untapCmd).Standalone()

	bootstrap_packages_brew_untapCmd.Flags().BoolP("dry-run", "n", false, "Show the config change without writing it")
	bootstrap_packages_brew_untapCmd.Flags().String("file", "", "Write to this config file or directory")
	bootstrap_packages_brew_untapCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrap_packages_brew_untapCmd.Flags().BoolP("local", "l", false, "Write to the local config instead of the global config")
	bootstrap_packages_brew_untapCmd.Flags().StringP("path", "p", "", "Write to this config file or directory")
	bootstrap_packages_brewCmd.AddCommand(bootstrap_packages_brew_untapCmd)
}
