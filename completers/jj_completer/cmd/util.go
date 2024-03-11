package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var utilCmd = &cobra.Command{
	Use:   "util",
	Short: "Infrequently used commands such as for generating shell completions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(utilCmd).Standalone()

	utilCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(utilCmd)
}
