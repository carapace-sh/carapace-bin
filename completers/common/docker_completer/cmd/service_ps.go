package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var service_psCmd = &cobra.Command{
	Use:   "ps [OPTIONS] SERVICE [SERVICE...]",
	Short: "List the tasks of one or more services",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(service_psCmd).Standalone()

	service_psCmd.Flags().StringP("filter", "f", "", "Filter output based on conditions provided")
	service_psCmd.Flags().String("format", "", "Pretty-print tasks using a Go template")
	service_psCmd.Flags().Bool("no-resolve", false, "Do not map IDs to Names")
	service_psCmd.Flags().Bool("no-trunc", false, "Do not truncate output")
	service_psCmd.Flags().BoolP("quiet", "q", false, "Only display task IDs")
	serviceCmd.AddCommand(service_psCmd)

	carapace.Gen(service_psCmd).FlagCompletion(carapace.ActionMap{
		"filter": carapace.ActionMultiPartsN("=", 2, func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionValuesDescribed(
					"desired-state", "Filter by desired state of tasks",
					"id", "Filter by ID",
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

	carapace.Gen(service_psCmd).PositionalAnyCompletion(docker.ActionServices())
}
