package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/systemctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var setEnvironmentCmd = &cobra.Command{
	Use:     "set-environment",
	Short:   "Set one or more environment variables",
	GroupID: "environment",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setEnvironmentCmd).Standalone()

	rootCmd.AddCommand(setEnvironmentCmd)

	carapace.Gen(setEnvironmentCmd).PositionalAnyCompletion(
		carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return action.ActionEnvironmentVariables(setEnvironmentCmd).Invoke(c).Suffix("=").ToA()
			default:
				return carapace.ActionValues()
			}
		}),
	)
}
