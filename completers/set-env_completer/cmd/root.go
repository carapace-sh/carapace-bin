package cmd

import (
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/env"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:                "set-env <name> <value>",
	Short:              "set environment variable",
	DisableFlagParsing: true,
	Run:                func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.Batch(
			carapace.ActionCallback(func(c carapace.Context) carapace.Action {
				alreadySet := make([]string, 0)
				for _, e := range c.Env {
					alreadySet = append(alreadySet, strings.SplitN(e, "=", 2)[0])
				}
				a := env.ActionKnownEnvironmentVariables().Filter(alreadySet...)
				if !strings.Contains(c.Value, "_") {
					return a.MultiParts("_") // only do multipart completion for first underscore
				}
				return a
			}),
			os.ActionEnvironmentVariables().Style(style.Blue),
		).ToA(),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return env.ActionEnvironmentVariableValues(c.Args[0])
		}),
	)
}
