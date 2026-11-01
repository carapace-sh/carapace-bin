package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var network_lsCmd = &cobra.Command{
	Use:     "ls [OPTIONS]",
	Short:   "List networks",
	Aliases: []string{"list"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_lsCmd).Standalone()

	network_lsCmd.Flags().StringP("filter", "f", "", "Provide filter values (e.g. \"driver=bridge\")")
	network_lsCmd.Flags().String("format", "", "Format output using a custom template:")
	network_lsCmd.Flags().Bool("no-trunc", false, "Do not truncate the output")
	network_lsCmd.Flags().BoolP("quiet", "q", false, "Only display network IDs")
	networkCmd.AddCommand(network_lsCmd)

	carapace.Gen(network_lsCmd).FlagCompletion(carapace.ActionMap{
		"filter": carapace.ActionMultiPartsN("=", 2, func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionValuesDescribed(
					"driver", "Filters networks by their driver",
					"id", "Filters networks by their ID",
					"label", "Filters networks by label",
					"name", "Filters all or part of a network's name",
					"scope", "Filters networks by scope (local, swarm)",
					"type", "Filters networks by type (custom, builtin)",
				).Suffix("=")
			default:
				switch c.Parts[0] {
				case "driver":
					return carapace.ActionValues("bridge", "host", "null", "overlay")
				case "scope":
					return carapace.ActionValues("local", "swarm")
				case "type":
					return carapace.ActionValues("builtin", "custom")
				case "label":
					return docker.ActionNetworkLabels()
				default:
					return carapace.ActionValues()
				}
			}
		}),
	})
}
