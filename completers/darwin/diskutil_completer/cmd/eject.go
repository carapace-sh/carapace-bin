package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var ejectCmd = &cobra.Command{
	Use:   "eject",
	Short: "Eject a disk",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ejectCmd).Standalone()
	rootCmd.AddCommand(ejectCmd)
	carapace.Gen(ejectCmd).PositionalCompletion(fs.ActionBlockDevices())
}
