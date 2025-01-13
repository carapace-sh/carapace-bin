package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/docker-compose_completer/cmd/action"
	"github.com/spf13/cobra"
)

var scaleCmd = &cobra.Command{
	Use:   "scale [SERVICE=REPLICAS...]",
	Short: "Scale services ",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(scaleCmd).Standalone()

	scaleCmd.Flags().Bool("no-deps", false, "Don't start linked services")
	rootCmd.AddCommand(scaleCmd)

	carapace.Gen(scaleCmd).PositionalAnyCompletion(
		carapace.ActionMultiPartsN("=", 2, func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return action.ActionServices(scaleCmd).Suffix("=")
			default:
				return carapace.ActionValues()
			}
		}),
	)
}
