package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var service_rmCmd = &cobra.Command{
	Use:     "rm SERVICE [SERVICE...]",
	Short:   "Remove one or more services",
	Aliases: []string{"remove"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(service_rmCmd).Standalone()

	serviceCmd.AddCommand(service_rmCmd)

	carapace.Gen(service_rmCmd).PositionalAnyCompletion(docker.ActionServices())
}
