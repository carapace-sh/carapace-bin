package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/docker-compose_completer/cmd/action"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create [OPTIONS] [SERVICE...]",
	Short: "Creates containers for a service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(createCmd).Standalone()

	createCmd.Flags().Bool("build", false, "Build images before starting containers")
	createCmd.Flags().Bool("force-recreate", false, "Recreate containers even if their configuration and image haven't changed")
	createCmd.Flags().Bool("no-build", false, "Don't build an image, even if it's policy")
	createCmd.Flags().Bool("no-recreate", false, "If containers already exist, don't recreate them. Incompatible with --force-recreate.")
	createCmd.Flags().String("pull", "", "Pull image before running (\"always\"|\"missing\"|\"never\"|\"build\")")
	createCmd.Flags().Bool("quiet-pull", false, "Pull without printing progress information")
	createCmd.Flags().Bool("remove-orphans", false, "Remove containers for services not defined in the Compose file")
	createCmd.Flags().StringSlice("scale", []string{}, "Scale SERVICE to NUM instances. Overrides the `scale` setting in the Compose file if present.")
	createCmd.Flags().BoolP("y", "y", false, "Assume \"yes\" as answer to all prompts and run non-interactively")
	rootCmd.AddCommand(createCmd)

	carapace.Gen(createCmd).FlagCompletion(carapace.ActionMap{
		"pull": carapace.ActionValues("always", "missing", "never", "build").StyleF(style.ForKeyword),
	})

	carapace.Gen(createCmd).PositionalAnyCompletion(
		action.ActionServices(createCmd).FilterArgs(),
	)
}
