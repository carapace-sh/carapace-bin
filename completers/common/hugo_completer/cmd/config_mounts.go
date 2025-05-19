package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_mountsCmd = &cobra.Command{
	Use:   "mounts",
	Short: "Print the configured file mounts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_mountsCmd).Standalone()
	configCmd.AddCommand(config_mountsCmd)
}
