package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var container_pruneCmd = &cobra.Command{
	Use:   "prune [OPTIONS]",
	Short: "Remove all stopped containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_pruneCmd).Standalone()

	container_pruneCmd.Flags().String("filter", "", "Provide filter values (e.g. \"until=<timestamp>\")")
	container_pruneCmd.Flags().BoolP("force", "f", false, "Do not prompt for confirmation")
	containerCmd.AddCommand(container_pruneCmd)

	carapace.Gen(container_pruneCmd).FlagCompletion(carapace.ActionMap{
		"filter": carapace.ActionMultiPartsN("=", 2, func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionValuesDescribed(
					"label", "Filter by label (key or key=value)",
					"until", "Prune objects created before given duration",
				).Suffix("=")
			default:
				switch c.Parts[0] {
				case "label":
					return docker.ActionContainerLabels()
				default:
					return carapace.ActionValues()
				}
			}
		}),
	})
}
