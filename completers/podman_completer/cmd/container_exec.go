package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var container_execCmd = &cobra.Command{
	Use:   "exec [options] CONTAINER COMMAND [ARG...]",
	Short: "Run a process in a running container",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_execCmd).Standalone()

	container_execCmd.Flags().BoolP("detach", "d", false, "Run the exec session in detached mode (backgrounded)")
	container_execCmd.Flags().String("detach-keys", "", "Select the key sequence for detaching a container")
	container_execCmd.Flags().StringSliceP("env", "e", []string{}, "Set environment variables")
	container_execCmd.Flags().StringSlice("env-file", []string{}, "Read in a file of environment variables")
	container_execCmd.Flags().BoolP("interactive", "i", false, "Make STDIN available to the contained process")
	container_execCmd.Flags().BoolP("latest", "l", false, "Act on the latest container podman is aware of")
	container_execCmd.Flags().StringSlice("preserve-fd", []string{}, "Pass a list of additional file descriptors to the container")
	container_execCmd.Flags().String("preserve-fds", "", "Pass N additional file descriptors to the container")
	container_execCmd.Flags().Bool("privileged", false, "Give the process extended Linux capabilities inside the container.  The default is false")
	container_execCmd.Flags().BoolP("tty", "t", false, "Allocate a pseudo-TTY. The default is false")
	container_execCmd.Flags().StringP("user", "u", "", "Sets the username or UID used and optionally the groupname or GID for the specified command")
	container_execCmd.Flags().String("wait", "", "Total seconds to wait for container to start")
	container_execCmd.Flags().StringP("workdir", "w", "", "Working directory inside the container")
	container_execCmd.Flag("wait").Hidden = true
	containerCmd.AddCommand(container_execCmd)
}
