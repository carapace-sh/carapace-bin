package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "Execute a command in a running container",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(execCmd).Standalone()

	execCmd.Flags().BoolS("T", "T", false, "Disable pseudo-tty allocation. By default `docker-compose exec`")
	execCmd.Flags().BoolP("detach", "d", false, "Detached mode: Run command in the background.")
	execCmd.Flags().String("index", "", "index of the container if there are multiple")
	execCmd.Flags().Bool("privileged", false, "Give extended privileges to the process.")
	execCmd.Flags().StringP("user", "u", "", "Run the command as this user.")
	rootCmd.AddCommand(execCmd)

	carapace.Gen(execCmd).FlagCompletion(carapace.ActionMap{
		"user": carapace.ActionUsers(),
	})

	carapace.Gen(execCmd).PositionalCompletion(
		ActionServices(),
	)
}
