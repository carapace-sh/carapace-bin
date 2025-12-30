package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var getpkgdepversionCommand = &cobra.Command{
	Use:   "getpkgdepversion <dep...>",
	Short: "Prints version constraint from dependencies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(getpkgdepversionCommand).Standalone()

	rootCmd.AddCommand(getpkgdepversionCommand)
}
