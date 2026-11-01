package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var volume_pruneCmd = &cobra.Command{
	Use:   "prune [OPTIONS]",
	Short: "Remove all unused local volumes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_pruneCmd).Standalone()

	volume_pruneCmd.Flags().BoolP("all", "a", false, "Remove all unused volumes, not just anonymous ones")
	volume_pruneCmd.Flags().String("filter", "", "Provide filter values (e.g. \"label=<label>\")")
	volume_pruneCmd.Flags().BoolP("force", "f", false, "Do not prompt for confirmation")
	volumeCmd.AddCommand(volume_pruneCmd)

	carapace.Gen(volume_pruneCmd).FlagCompletion(carapace.ActionMap{
		"filter": carapace.ActionMultiPartsN("=", 2, func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionValuesDescribed(
					"label", "Filter by label (key or key=value)",
				).Suffix("=")
			default:
				switch c.Parts[0] {
				case "label":
					return docker.ActionVolumeLabels()
				default:
					return carapace.ActionValues()
				}
			}
		}),
	})
}
