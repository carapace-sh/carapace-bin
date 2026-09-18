package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:     "search",
	Short:   "Search the Docker Hub for images",
	GroupID: "common",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(searchCmd).Standalone()

	searchCmd.Flags().StringP("filter", "f", "", "Filter output based on conditions provided")
	searchCmd.Flags().String("format", "", "Pretty-print search using a Go template")
	searchCmd.Flags().String("limit", "", "Max number of search results")
	searchCmd.Flags().Bool("no-trunc", false, "Don't truncate output")
	rootCmd.AddCommand(searchCmd)

	carapace.Gen(searchCmd).FlagCompletion(carapace.ActionMap{
		"filter": carapace.ActionMultiPartsN("=", 2, func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionValuesDescribed(
					"is-official", "Filter by official images",
					"stars", "Filter by star count",
				).Suffix("=")
			default:
				switch c.Parts[0] {
				case "is-official":
					return carapace.ActionValues("true", "false").StyleF(style.ForKeyword)
				default:
					return carapace.ActionValues()
				}
			}
		}),
	})
}
