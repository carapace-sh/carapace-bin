package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var volume_importCmd = &cobra.Command{
	Use:   "import VOLUME [SOURCE]",
	Short: "Import a tarball contents into a podman volume",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_importCmd).Standalone()

	volumeCmd.AddCommand(volume_importCmd)
}
