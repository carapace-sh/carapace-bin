package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var unmountCmd = &cobra.Command{
	Use:   "unmount",
	Short: "Unmount a volume",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unmountCmd).Standalone()
	rootCmd.AddCommand(unmountCmd)
	carapace.Gen(unmountCmd).PositionalCompletion(fs.ActionBlockDevices())
}
