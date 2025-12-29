package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var getsystemdirCommand = &cobra.Command{
	Use:   "getsystemdir",
	Short: "Prints the system xbps.d directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(getsystemdirCommand).Standalone()

	rootCmd.AddCommand(getsystemdirCommand)
}
