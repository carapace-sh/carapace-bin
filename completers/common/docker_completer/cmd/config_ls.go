package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var config_lsCmd = &cobra.Command{
	Use:     "ls [OPTIONS]",
	Short:   "List configs",
	Aliases: []string{"list"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_lsCmd).Standalone()

	config_lsCmd.Flags().StringP("filter", "f", "", "Filter output based on conditions provided")
	config_lsCmd.Flags().String("format", "", "Format output using a custom template:")
	config_lsCmd.Flags().BoolP("quiet", "q", false, "Only display IDs")
	configCmd.AddCommand(config_lsCmd)

	carapace.Gen(config_lsCmd).FlagCompletion(carapace.ActionMap{
		"filter": carapace.ActionMultiPartsN("=", 2, func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionValuesDescribed(
					"id", "Filter by ID",
					"label", "Filter by label (key or key=value)",
					"name", "Filter by name",
				).Suffix("=")
			default:
				switch c.Parts[0] {
				case "label":
					return docker.ActionConfigLabels()
				default:
					return carapace.ActionValues()
				}
			}
		}),
	})
}
