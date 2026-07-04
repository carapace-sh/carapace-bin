package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var mountCmd = &cobra.Command{
	Use:   "mount",
	Short: "Mount a volume",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mountCmd).Standalone()
	rootCmd.AddCommand(mountCmd)
	carapace.Gen(mountCmd).PositionalCompletion(fs.ActionBlockDevices())
}
