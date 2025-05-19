package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/env"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var execCmd = &cobra.Command{
	Use:     "exec [OPTIONS] CONTAINER COMMAND [ARG...]",
	Short:   "Execute a command in a running container",
	GroupID: "common",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(execCmd).Standalone()

	execCmd.Flags().BoolP("detach", "d", false, "Detached mode: run command in the background")
	execCmd.Flags().String("detach-keys", "", "Override the key sequence for detaching a container")
	execCmd.Flags().StringP("env", "e", "", "Set environment variables")
	execCmd.Flags().String("env-file", "", "Read in a file of environment variables")
	execCmd.Flags().BoolP("interactive", "i", false, "Keep STDIN open even if not attached")
	execCmd.Flags().Bool("privileged", false, "Give extended privileges to the command")
	execCmd.Flags().BoolP("tty", "t", false, "Allocate a pseudo-TTY")
	execCmd.Flags().StringP("user", "u", "", "Username or UID (format: \"<name|uid>[:<group|gid>]\")")
	execCmd.Flags().StringP("workdir", "w", "", "Working directory inside the container")
	rootCmd.AddCommand(execCmd)

	carapace.Gen(execCmd).FlagCompletion(carapace.ActionMap{
		"detach-keys": docker.ActionDetachKeys(),
		"env":         env.ActionNameValues(false),
		"env-file":    carapace.ActionFiles(),
	})
}
