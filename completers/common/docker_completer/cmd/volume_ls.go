package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var volume_lsCmd = &cobra.Command{
	Use:     "ls [OPTIONS]",
	Short:   "List volumes",
	Aliases: []string{"list"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_lsCmd).Standalone()

	volume_lsCmd.Flags().Bool("cluster", false, "Display only cluster volumes, and use cluster volume list formatting")
	volume_lsCmd.Flags().StringP("filter", "f", "", "Provide filter values (e.g. \"dangling=true\")")
	volume_lsCmd.Flags().String("format", "", "Format output using a custom template:")
	volume_lsCmd.Flags().BoolP("quiet", "q", false, "Only display volume names")
	volumeCmd.AddCommand(volume_lsCmd)

	carapace.Gen(volume_lsCmd).FlagCompletion(carapace.ActionMap{
		"filter": carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionValuesDescribed(
					"dangling", "Filter by dangling state",
					"driver", "Filter by driver",
					"label", "Filter by label (key or key=value)",
					"name", "Filter by name",
				).Suffix("=")
			default:
				switch c.Parts[0] {
				case "dangling":
					return carapace.ActionValues("true", "false").StyleF(style.ForKeyword)
				case "label":
					return docker.ActionVolumeLabels()
				default:
					return carapace.ActionValues()
				}
			}
		}),
	})
}
