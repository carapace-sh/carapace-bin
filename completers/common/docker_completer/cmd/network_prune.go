package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var network_pruneCmd = &cobra.Command{
	Use:   "prune [OPTIONS]",
	Short: "Remove all unused networks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_pruneCmd).Standalone()

	network_pruneCmd.Flags().String("filter", "", "Provide filter values (e.g. \"until=<timestamp>\")")
	network_pruneCmd.Flags().BoolP("force", "f", false, "Do not prompt for confirmation")
	networkCmd.AddCommand(network_pruneCmd)

	carapace.Gen(network_pruneCmd).FlagCompletion(carapace.ActionMap{
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
					return docker.ActionNetworkLabels()
				default:
					return carapace.ActionValues()
				}
			}
		}),
	})
}
