package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "build the new configuration without activating it or adding it to the boot menu",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(buildCmd).Standalone()
	rootCmd.AddCommand(buildCmd)
}
