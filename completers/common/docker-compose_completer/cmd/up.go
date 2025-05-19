package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/docker-compose_completer/cmd/action"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var upCmd = &cobra.Command{
	Use:   "up [OPTIONS] [SERVICE...]",
	Short: "Create and start containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upCmd).Standalone()

	upCmd.Flags().Bool("abort-on-container-exit", false, "Stops all containers if any container was stopped. Incompatible with -d")
	upCmd.Flags().Bool("abort-on-container-failure", false, "Stops all containers if any container exited with failure. Incompatible with -d")
	upCmd.Flags().Bool("always-recreate-deps", false, "Recreate dependent containers. Incompatible with --no-recreate.")
	upCmd.Flags().StringSlice("attach", nil, "Restrict attaching to the specified services. Incompatible with --attach-dependencies.")
	upCmd.Flags().Bool("attach-dependencies", false, "Automatically attach to log output of dependent services")
	upCmd.Flags().Bool("build", false, "Build images before starting containers")
	upCmd.Flags().BoolP("detach", "d", false, "Detached mode: Run containers in the background")
	upCmd.Flags().String("exit-code-from", "", "Return the exit code of the selected service container. Implies --abort-on-container-exit")
	upCmd.Flags().Bool("force-recreate", false, "Recreate containers even if their configuration and image haven't changed")
	upCmd.Flags().Bool("menu", false, "Enable interactive shortcuts when running attached. Incompatible with --detach. Can also be enable/disable by setting COMPOSE_MENU environment var.")
	upCmd.Flags().StringSlice("no-attach", nil, "Do not attach (stream logs) to the specified services")
	upCmd.Flags().Bool("no-build", false, "Don't build an image, even if it's policy")
	upCmd.Flags().Bool("no-color", false, "Produce monochrome output")
	upCmd.Flags().Bool("no-deps", false, "Don't start linked services")
	upCmd.Flags().Bool("no-log-prefix", false, "Don't print prefix in logs")
	upCmd.Flags().Bool("no-recreate", false, "If containers already exist, don't recreate them. Incompatible with --force-recreate.")
	upCmd.Flags().Bool("no-start", false, "Don't start the services after creating them")
	upCmd.Flags().String("pull", "", "Pull image before running (\"always\"|\"missing\"|\"never\")")
	upCmd.Flags().Bool("quiet-pull", false, "Pull without printing progress information")
	upCmd.Flags().Bool("remove-orphans", false, "Remove containers for services not defined in the Compose file")
	upCmd.Flags().BoolP("renew-anon-volumes", "V", false, "Recreate anonymous volumes instead of retrieving data from the previous containers")
	upCmd.Flags().StringSlice("scale", nil, "Scale SERVICE to NUM instances. Overrides the `scale` setting in the Compose file if present.")
	upCmd.Flags().StringP("timeout", "t", "", "Use this timeout in seconds for container shutdown when attached or when containers are already running")
	upCmd.Flags().Bool("timestamps", false, "Show timestamps")
	upCmd.Flags().Bool("wait", false, "Wait for services to be running|healthy. Implies detached mode.")
	upCmd.Flags().String("wait-timeout", "", "Maximum duration in seconds to wait for the project to be running|healthy")
	upCmd.Flags().BoolP("watch", "w", false, "Watch source code and rebuild/refresh containers when files are updated.")
	upCmd.Flags().BoolP("yes", "y", false, "Assume \"yes\" as answer to all prompts and run non-interactively")
	rootCmd.AddCommand(upCmd)

	carapace.Gen(upCmd).FlagCompletion(carapace.ActionMap{
		"attach":         action.ActionServices(upCmd),
		"exit-code-from": action.ActionServices(upCmd),
		"no-attach":      action.ActionServices(upCmd).UniqueList(","),
		"pull":           carapace.ActionValues("always", "missing", "never").StyleF(style.ForKeyword),
		"scale": carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return action.ActionServices(upCmd).Invoke(c).Suffix("=").ToA()
			default:
				return carapace.ActionValues()
			}
		}),
	})

	carapace.Gen(upCmd).PositionalAnyCompletion(
		action.ActionServices(upCmd).FilterArgs(),
	)
}
