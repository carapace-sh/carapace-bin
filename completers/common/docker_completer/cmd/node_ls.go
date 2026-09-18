package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var node_lsCmd = &cobra.Command{
	Use:     "ls [OPTIONS]",
	Short:   "List nodes in the swarm",
	Aliases: []string{"list"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(node_lsCmd).Standalone()

	node_lsCmd.Flags().StringP("filter", "f", "", "Filter output based on conditions provided")
	node_lsCmd.Flags().String("format", "", "Format output using a custom template:")
	node_lsCmd.Flags().BoolP("quiet", "q", false, "Only display IDs")
	nodeCmd.AddCommand(node_lsCmd)

	carapace.Gen(node_lsCmd).FlagCompletion(carapace.ActionMap{
		"filter": carapace.ActionMultiPartsN("=", 2, func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionValuesDescribed(
					"id", "Filter by ID",
					"label", "Filter by label (key or key=value)",
					"membership", "Filter by swarm membership",
					"name", "Filter by name",
					"role", "Filter by swarm role",
				).Suffix("=")
			default:
				switch c.Parts[0] {
				case "membership":
					return carapace.ActionValues("pending", "accepted", "rejected")
				case "role":
					return carapace.ActionValues("manager", "worker")
				default:
					return carapace.ActionValues()
				}
			}
		}),
	})
}
