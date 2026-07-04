package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/launchctl"
	"github.com/spf13/cobra"
)

var bootoutCmd = &cobra.Command{
	Use:   "bootout",
	Short: "Unbootstrap a service from the specified domain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootoutCmd).Standalone()
	rootCmd.AddCommand(bootoutCmd)
	carapace.Gen(bootoutCmd).PositionalAnyCompletion(launchctl.ActionServices())
}
