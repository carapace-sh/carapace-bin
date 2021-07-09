package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/docker-compose_completer/cmd/action"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create services",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(createCmd).Standalone()

	createCmd.Flags().Bool("build", false, "Build images before creating containers.")
	createCmd.Flags().Bool("force-recreate", false, "Recreate containers even if their configuration and")
	createCmd.Flags().Bool("no-build", false, "Don't build an image, even if it's missing.")
	createCmd.Flags().Bool("no-recreate", false, "If containers already exist, don't recreate them.")
	rootCmd.AddCommand(createCmd)

	carapace.Gen(createCmd).PositionalAnyCompletion(
		action.ActionServices(createCmd),
	)
}
