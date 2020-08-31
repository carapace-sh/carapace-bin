package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Create and start containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	upCmd.Flags().Bool("abort-on-container-exit", false, "Stops all containers if any container was")
	upCmd.Flags().Bool("abort-on-container-exit.", false, "")
	upCmd.Flags().Bool("always-recreate-deps", false, "Recreate dependent containers.")
	upCmd.Flags().Bool("attach-dependencies", false, "Attach to dependent containers")
	upCmd.Flags().Bool("build", false, "Build images before starting containers.")
	upCmd.Flags().BoolP("detach", "d", false, "Detached mode: Run containers in the background,")
	upCmd.Flags().String("exit-code-from", "", "Return the exit code of the selected service")
	upCmd.Flags().Bool("force-recreate", false, "Recreate containers even if their configuration")
	upCmd.Flags().Bool("no-build", false, "Don't build an image, even if it's missing.")
	upCmd.Flags().Bool("no-color", false, "Produce monochrome output.")
	upCmd.Flags().Bool("no-deps", false, "Don't start linked services.")
	upCmd.Flags().Bool("no-recreate", false, "If containers already exist, don't recreate")
	upCmd.Flags().Bool("no-start", false, "Don't start the services after creating them.")
	upCmd.Flags().Bool("quiet-pull", false, "Pull without printing progress information")
	upCmd.Flags().Bool("remove-orphans", false, "Remove containers for services not defined")
	upCmd.Flags().BoolP("renew-anon-volumes", "V", false, "Recreate anonymous volumes instead of retrieving")
	upCmd.Flags().String("scale", "", "Scale SERVICE to NUM instances. Overrides the")
	upCmd.Flags().StringP("timeout", "t", "", "Use this timeout in seconds for container")
	rootCmd.AddCommand(upCmd)

	carapace.Gen(upCmd).FlagCompletion(carapace.ActionMap{
		"exit-code-from": ActionServices(),
	})

	carapace.Gen(upCmd).PositionalAnyCompletion(ActionServices())
}
