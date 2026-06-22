package common

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/rails"
	"github.com/spf13/cobra"
)

func AddEnvironmentFlag(cmd *cobra.Command, defaultEnv string) {
	cmd.Flags().StringP("environment", "e", defaultEnv, "Rails environment")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"environment": rails.ActionEnvironments(),
	})
}
