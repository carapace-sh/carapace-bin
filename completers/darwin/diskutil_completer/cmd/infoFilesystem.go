package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var infoFilesystemCmd = &cobra.Command{
	Use:   "infoFilesystem",
	Short: "Get filesystem information on a specific volume",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(infoFilesystemCmd).Standalone()
	rootCmd.AddCommand(infoFilesystemCmd)
	carapace.Gen(infoFilesystemCmd).PositionalCompletion(fs.ActionBlockDevices())
}
