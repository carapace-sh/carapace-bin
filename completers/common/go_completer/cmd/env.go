package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/env"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/golang"
	"github.com/spf13/cobra"
)

var envCmd = &cobra.Command{
	Use:   "env",
	Short: "print Go environment information",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(envCmd).Standalone()
	envCmd.Flags().SetInterspersed(false)

	envCmd.Flags().StringS("C", "C", "", "Change to dir before running the command")
	envCmd.Flags().BoolS("json", "json", false, "print the environment in JSON format")
	envCmd.Flags().BoolS("u", "u", false, "unsets the default setting for the named environment variables")
	envCmd.Flags().BoolS("w", "w", false, "changes the default settings of the named environment variables")
	rootCmd.AddCommand(envCmd)

	envCmd.MarkFlagsMutuallyExclusive("u", "w")

	carapace.Gen(envCmd).FlagCompletion(carapace.ActionMap{
		"C": carapace.ActionDirectories(),
	})

	carapace.Gen(envCmd).PositionalAnyCompletion(
		carapace.ActionMultiPartsN("=", 2, func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				if envCmd.Flag("w").Changed {
					return golang.ActionEnvironmentVariables().Suffix("=")
				}
				return golang.ActionEnvironmentVariables()
			default:
				if envCmd.Flag("w").Changed {
					return env.ActionValues(c.Parts[0])
				}
				return carapace.ActionValues()
			}
		}),
	)
}
