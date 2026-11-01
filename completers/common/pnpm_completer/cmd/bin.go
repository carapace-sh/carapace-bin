package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var binCmd = &cobra.Command{
	Use:   "bin",
	Short: "Print the directory where pnpm will install executables",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(binCmd).Standalone()

	binCmd.Flags().BoolP("global", "g", false, "Print the global executables directory")
	binCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(binCmd)
}
