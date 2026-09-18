package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var node_psCmd = &cobra.Command{
	Use:   "ps [OPTIONS] [NODE...]",
	Short: "List tasks running on one or more nodes, defaults to current node",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(node_psCmd).Standalone()

	node_psCmd.Flags().StringP("filter", "f", "", "Filter output based on conditions provided")
	node_psCmd.Flags().String("format", "", "Pretty-print tasks using a Go template")
	node_psCmd.Flags().Bool("no-resolve", false, "Do not map IDs to Names")
	node_psCmd.Flags().Bool("no-trunc", false, "Do not truncate output")
	node_psCmd.Flags().BoolP("quiet", "q", false, "Only display task IDs")
	nodeCmd.AddCommand(node_psCmd)

	carapace.Gen(node_psCmd).FlagCompletion(carapace.ActionMap{
		"filter": carapace.ActionMultiPartsN("=", 2, func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionValuesDescribed(
					"desired-state", "Filter by desired state of tasks",
					"id", "Filter by ID",
					"label", "Filter by label (key or key=value)",
					"name", "Filter by name",
					"node", "Filter by node",
				).Suffix("=")
			default:
				switch c.Parts[0] {
				case "desired-state":
					return carapace.ActionValues("running", "shutdown", "accepted")
				case "node":
					return docker.ActionNodes()
				default:
					return carapace.ActionValues()
				}
			}
		}),
	})

	carapace.Gen(node_psCmd).PositionalAnyCompletion(docker.ActionNodes())
}
