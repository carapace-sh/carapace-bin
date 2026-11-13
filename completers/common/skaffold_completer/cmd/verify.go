package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Run verification tests against skaffold deployments",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(verifyCmd).Standalone()

	verifyCmd.Flags().Bool("assume-yes", false, "If true, skaffold will skip yes/no confirmation from the user and default to yes")
	verifyCmd.Flags().StringP("build-artifacts", "a", "", "File containing pre-built images to use instead of rebuilding artifacts. A sample file looks like the following:")
	verifyCmd.Flags().StringP("default-repo", "d", "", "Default repository value (overrides global config)")
	verifyCmd.Flags().String("docker-network", "", "Name of an existing docker network to use when running the verify tests. If not specified, Skaffold will create a new network to use of the form 'skaffold-network-<uuid>'")
	verifyCmd.Flags().String("env-file", "", "File containing env var key-value pairs that will be set in all verify container envs")
	verifyCmd.Flags().String("event-log-file", "", "Save Skaffold events to the provided file after skaffold has finished executing, requires --rpc-port or --rpc-http-port")
	verifyCmd.Flags().StringP("filename", "f", "skaffold.yaml", "Path or URL to the Skaffold config file")
	verifyCmd.Flags().StringSliceP("module", "m", nil, "Filter Skaffold configs to only the provided named modules")
	verifyCmd.Flags().StringP("namespace", "n", "", "Runs deployments in the specified namespace. When used with 'render' command, renders manifests contain the namespace")
	verifyCmd.Flags().StringSlice("port-forward", []string{"off"}, "Port-forward exposes service ports and container ports within pods and other resources (off, user, services, debug, pods)")
	verifyCmd.Flags().StringSliceP("profile", "p", nil, "Activate profiles by name (prefixed with `-` to disable a profile)")
	verifyCmd.Flags().Bool("profile-auto-activation", true, "Set to false to disable profile auto activation")
	verifyCmd.Flags().Bool("propagate-profiles", true, "Setting '--propagate-profiles=false' disables propagating profiles set by the '--profile' flag across config dependencies. This mean that only profiles defined directly in the target 'skaffold.yaml' file are activated.")
	verifyCmd.Flags().String("remote-cache-dir", "", "Specify the location of the remote cache (default $HOME/.skaffold/remote-cache)")
	verifyCmd.Flags().String("rpc-http-port", "", "tcp port to expose the Skaffold API over HTTP REST")
	verifyCmd.Flags().String("rpc-port", "", "tcp port to expose the Skaffold API over gRPC")
	verifyCmd.Flags().Bool("status-check", false, "Wait for deployed resources to stabilize")
	verifyCmd.Flags().String("sync-remote-cache", "always", "Controls how Skaffold manages the remote config cache (see `remote-cache-dir`). One of `always` (default), `missing`, or `never`. `always` syncs remote repositories to latest on access. `missing` only clones remote repositories if they do not exist locally. `never` means the user takes responsibility for updating remote repositories.")
	verifyCmd.Flags().StringP("user", "u", "", "")
	verifyCmd.Flag("event-log-file").Hidden = true
	verifyCmd.Flag("port-forward").NoOptDefVal = " "
	verifyCmd.Flag("user").Hidden = true
	applyCommonFlagCompletion(verifyCmd)
	rootCmd.AddCommand(verifyCmd)
}
