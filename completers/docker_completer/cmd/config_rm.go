package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var config_rmCmd = &cobra.Command{
	Use:     "rm CONFIG [CONFIG...]",
	Short:   "Remove one or more configs",
	Aliases: []string{"remove"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_rmCmd).Standalone()

	configCmd.AddCommand(config_rmCmd)

	carapace.Gen(config_rmCmd).PositionalAnyCompletion(docker.ActionConfigs())
}
