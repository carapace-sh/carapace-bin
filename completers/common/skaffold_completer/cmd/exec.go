package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "Execute a custom action",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(execCmd).Standalone()

	execCmd.Flags().Bool("assume-yes", false, "If true, skaffold will skip yes/no confirmation from the user and default to yes")
	execCmd.Flags().StringP("build-artifacts", "a", "", "File containing pre-built images to use instead of rebuilding artifacts. A sample file looks like the following:")
	execCmd.Flags().StringP("default-repo", "d", "", "Default repository value (overrides global config)")
	execCmd.Flags().String("docker-network", "", "Name of an existing docker network to use when running the verify tests. If not specified, Skaffold will create a new network to use of the form 'skaffold-network-<uuid>'")
	execCmd.Flags().String("env-file", "", "File containing env var key-value pairs that will be set in all verify container envs")
	execCmd.Flags().String("event-log-file", "", "Save Skaffold events to the provided file after skaffold has finished executing, requires --rpc-port or --rpc-http-port")
	execCmd.Flags().StringP("filename", "f", "skaffold.yaml", "Path or URL to the Skaffold config file")
	execCmd.Flags().StringSliceP("module", "m", nil, "Filter Skaffold configs to only the provided named modules")
	execCmd.Flags().StringP("namespace", "n", "", "Runs deployments in the specified namespace. When used with 'render' command, renders manifests contain the namespace")
	execCmd.Flags().StringSlice("port-forward", []string{"off"}, "Port-forward exposes service ports and container ports within pods and other resources (off, user, services, debug, pods)")
	execCmd.Flags().StringSliceP("profile", "p", nil, "Activate profiles by name (prefixed with `-` to disable a profile)")
	execCmd.Flags().Bool("profile-auto-activation", true, "Set to false to disable profile auto activation")
	execCmd.Flags().Bool("propagate-profiles", true, "Setting '--propagate-profiles=false' disables propagating profiles set by the '--profile' flag across config dependencies. This mean that only profiles defined directly in the target 'skaffold.yaml' file are activated.")
	execCmd.Flags().String("remote-cache-dir", "", "Specify the location of the remote cache (default $HOME/.skaffold/remote-cache)")
	execCmd.Flags().String("rpc-http-port", "", "tcp port to expose the Skaffold API over HTTP REST")
	execCmd.Flags().String("rpc-port", "", "tcp port to expose the Skaffold API over gRPC")
	execCmd.Flags().String("sync-remote-cache", "always", "Controls how Skaffold manages the remote config cache (see `remote-cache-dir`). One of `always` (default), `missing`, or `never`. `always` syncs remote repositories to latest on access. `missing` only clones remote repositories if they do not exist locally. `never` means the user takes responsibility for updating remote repositories.")
	execCmd.Flags().StringP("user", "u", "", "")
	execCmd.Flag("event-log-file").Hidden = true
	execCmd.Flag("port-forward").NoOptDefVal = " "
	execCmd.Flag("user").Hidden = true
	applyCommonFlagCompletion(execCmd)
	rootCmd.AddCommand(execCmd)
}
