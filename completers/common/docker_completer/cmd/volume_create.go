package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var volume_createCmd = &cobra.Command{
	Use:   "create [OPTIONS] [VOLUME]",
	Short: "Create a volume",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_createCmd).Standalone()

	volume_createCmd.Flags().String("availability", "", "Cluster Volume availability (\"active\", \"pause\", \"drain\")")
	volume_createCmd.Flags().StringP("driver", "d", "", "Specify volume driver name")
	volume_createCmd.Flags().String("group", "", "Cluster Volume group (cluster volumes)")
	volume_createCmd.Flags().String("label", "", "Set metadata for a volume")
	volume_createCmd.Flags().String("limit-bytes", "", "Minimum size of the Cluster Volume in bytes")
	volume_createCmd.Flags().String("name", "", "Specify volume name")
	volume_createCmd.Flags().StringP("opt", "o", "", "Set driver specific options")
	volume_createCmd.Flags().String("required-bytes", "", "Maximum size of the Cluster Volume in bytes")
	volume_createCmd.Flags().String("scope", "", "Cluster Volume access scope (\"single\", \"multi\")")
	volume_createCmd.Flags().String("secret", "", "Cluster Volume secrets")
	volume_createCmd.Flags().String("sharing", "", "Cluster Volume access sharing (\"none\", \"readonly\", \"onewriter\", \"all\")")
	volume_createCmd.Flags().String("topology-preferred", "", "A topology that the Cluster Volume would be preferred in")
	volume_createCmd.Flags().String("topology-required", "", "A topology that the Cluster Volume must be accessible from")
	volume_createCmd.Flags().String("type", "", "Cluster Volume access type (\"mount\", \"block\")")
	volume_createCmd.Flag("name").Hidden = true
	volumeCmd.AddCommand(volume_createCmd)

	carapace.Gen(volume_createCmd).FlagCompletion(carapace.ActionMap{
		"availability": carapace.ActionValues("active", "pause", "drain"),
		"driver":       carapace.ActionValues("local"),
		"scope":        carapace.ActionValues("single", "multi"),
		"sharing":      carapace.ActionValues("none", "readonly", "onewriter", "all"),
		"type":         carapace.ActionValues("mount", "block"),
	})

	carapace.Gen(volume_createCmd).PositionalCompletion(docker.ActionVolumes())
}
