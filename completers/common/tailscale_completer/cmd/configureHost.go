package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configureHostCmd = &cobra.Command{
	Use:   "configure-host",
	Short: "Compatibility alias for configure synology",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configureHostCmd).Standalone()

	rootCmd.AddCommand(configureHostCmd)
}
