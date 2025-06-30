package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/docker-compose_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/env"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run [OPTIONS] SERVICE [COMMAND] [ARGS...]",
	Short: "Run a one-off command on a service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runCmd).Standalone()

	runCmd.Flags().Bool("build", false, "Build image before starting container")
	runCmd.Flags().String("cap-add", "", "Add Linux capabilities")
	runCmd.Flags().String("cap-drop", "", "Drop Linux capabilities")
	runCmd.Flags().BoolP("detach", "d", false, "Run container in background and print container ID")
	runCmd.Flags().String("entrypoint", "", "Override the entrypoint of the image")
	runCmd.Flags().StringSliceP("env", "e", nil, "Set environment variables")
	runCmd.Flags().StringSlice("env-from-file", nil, "Set environment variables from file")
	runCmd.Flags().BoolP("interactive", "i", false, "Keep STDIN open even if not attached")
	runCmd.Flags().StringSliceP("label", "l", nil, "Add or override a label")
	runCmd.Flags().String("name", "", "Assign a name to the container")
	runCmd.Flags().BoolP("no-TTY", "T", false, "Disable pseudo-TTY allocation (default: auto-detected)")
	runCmd.Flags().Bool("no-deps", false, "Don't start linked services")
	runCmd.Flags().StringSliceP("publish", "p", nil, "Publish a container's port(s) to the host")
	runCmd.Flags().String("pull", "", "Pull image before running (\"always\"|\"missing\"|\"never\")")
	runCmd.Flags().BoolP("quiet", "q", false, "Don't print anything to STDOUT")
	runCmd.Flags().Bool("quiet-build", false, "Suppress progress output from the build process")
	runCmd.Flags().Bool("quiet-pull", false, "Pull without printing progress information")
	runCmd.Flags().Bool("remove-orphans", false, "Remove containers for services not defined in the Compose file")
	runCmd.Flags().Bool("rm", false, "Automatically remove the container when it exits")
	runCmd.Flags().BoolP("service-ports", "P", false, "Run command with all service's ports enabled and mapped to the host")
	runCmd.Flags().BoolP("tty", "t", false, "Allocate a pseudo-TTY")
	runCmd.Flags().Bool("use-aliases", false, "Use the service's network useAliases in the network(s) the container connects to")
	runCmd.Flags().StringP("user", "u", "", "Run as specified username or uid")
	runCmd.Flags().StringSliceP("volume", "v", nil, "Bind mount a volume")
	runCmd.Flags().StringP("workdir", "w", "", "Working directory inside the container")
	runCmd.Flag("tty").Hidden = true
	rootCmd.AddCommand(runCmd)

	// TODO flag completion
	carapace.Gen(runCmd).FlagCompletion(carapace.ActionMap{
		"env":           env.ActionNameValues(false),
		"env-from-file": carapace.ActionFiles(),
		"pull":          carapace.ActionValues("always", "missing", "never").StyleF(style.ForKeyword),
		"volume":        action.ActionVolumes(runCmd),
	})

	carapace.Gen(runCmd).PositionalCompletion(
		action.ActionServices(runCmd),
	)
}
