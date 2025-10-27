package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var volume_updateCmd = &cobra.Command{
	Use:   "update [OPTIONS] [VOLUME]",
	Short: "Update a volume (cluster volumes only)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_updateCmd).Standalone()

	volume_updateCmd.Flags().String("availability", "", "Cluster Volume availability (\"active\", \"pause\", \"drain\")")
	volumeCmd.AddCommand(volume_updateCmd)

	carapace.Gen(volume_updateCmd).FlagCompletion(carapace.ActionMap{
		"availability": carapace.ActionValues("active", "pause", "drain"),
	})

	carapace.Gen(volume_updateCmd).PositionalCompletion(
		docker.ActionVolumes(),
	)
}
