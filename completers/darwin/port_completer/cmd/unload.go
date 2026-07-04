package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var unloadCmd = &cobra.Command{
	Use:   "unload",
	Short: "Unload a port's daemon via launchctl",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unloadCmd).Standalone()
	rootCmd.AddCommand(unloadCmd)
}
