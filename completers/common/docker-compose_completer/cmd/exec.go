package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/docker-compose_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/env"
	"github.com/spf13/cobra"
)

var execCmd = &cobra.Command{
	Use:   "exec [OPTIONS] SERVICE COMMAND [ARGS...]",
	Short: "Execute a command in a running container",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(execCmd).Standalone()

	execCmd.Flags().BoolP("detach", "d", false, "Detached mode: Run command in the background")
	execCmd.Flags().StringSliceP("env", "e", nil, "Set environment variables")
	execCmd.Flags().String("index", "", "Index of the container if service has multiple replicas")
	execCmd.Flags().BoolP("interactive", "i", false, "Keep STDIN open even if not attached")
	execCmd.Flags().BoolP("no-TTY", "T", false, "Disable pseudo-TTY allocation. By default `docker compose exec` allocates a TTY.")
	execCmd.Flags().Bool("privileged", false, "Give extended privileges to the process")
	execCmd.Flags().BoolP("tty", "t", false, "Allocate a pseudo-TTY")
	execCmd.Flags().StringP("user", "u", "", "Run the command as this user")
	execCmd.Flags().StringP("workdir", "w", "", "Path to workdir directory for this command")
	execCmd.Flag("interactive").Hidden = true
	execCmd.Flag("tty").Hidden = true
	rootCmd.AddCommand(execCmd)

	// TODO workdir completion
	// TODO index
	carapace.Gen(execCmd).FlagCompletion(carapace.ActionMap{
		"env": env.ActionNameValues(false),
		"user": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 0 {
				if index, err := execCmd.Flags().GetInt("index"); err != nil {
					return carapace.ActionMessage(err.Error())
				} else {
					return action.ActionContainerUsers(execCmd, c.Args[0], index)
				}
			}
			return carapace.ActionValues()
		}),
		"workdir": carapace.ActionDirectories(),
	})

	carapace.Gen(execCmd).PositionalCompletion(
		action.ActionServices(execCmd),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if index, err := execCmd.Flags().GetInt("index"); err != nil {
				return carapace.ActionMessage(err.Error())
			} else {
				return action.ActionFiles(execCmd, c.Args[0], index)
			}
		}),
	)
}
