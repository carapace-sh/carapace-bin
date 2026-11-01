package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var secret_lsCmd = &cobra.Command{
	Use:     "ls [OPTIONS]",
	Short:   "List secrets",
	Aliases: []string{"list"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secret_lsCmd).Standalone()

	secret_lsCmd.Flags().StringP("filter", "f", "", "Filter output based on conditions provided")
	secret_lsCmd.Flags().String("format", "", "Format output using a custom template:")
	secret_lsCmd.Flags().BoolP("quiet", "q", false, "Only display IDs")
	secretCmd.AddCommand(secret_lsCmd)

	carapace.Gen(secret_lsCmd).FlagCompletion(carapace.ActionMap{
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
					return docker.ActionSecretLabels()
				default:
					return carapace.ActionValues()
				}
			}
		}),
	})
}
