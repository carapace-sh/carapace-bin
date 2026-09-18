package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Runs a package's \"stop\" script, if one was provided",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stopCmd).Standalone()

	stopCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(stopCmd)
}
