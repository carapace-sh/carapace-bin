package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:     "install",
	GroupID: "configuration",
	Short:   "Configure environment paths and shell settings",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installCmd).Standalone()
	rootCmd.AddCommand(installCmd)
}
