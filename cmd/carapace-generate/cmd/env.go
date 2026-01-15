package cmd

import (
	"errors"
	"os"
	"os/exec"
	"slices"
	"strings"

	"github.com/spf13/cobra"
)

var envCmd = &cobra.Command{
	Use:   "env",
	Short: "",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		i := slices.IndexFunc(args, func(s string) bool { return !strings.Contains(s, "=") })
		if i < 0 {
			return errors.New("missing command")
		}

		command := exec.Command(args[i], args[i+1:]...)
		command.Env = append(os.Environ(), args[:i]...)
		command.Stdin = os.Stdin
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr
		return command.Run()
	},
}

func init() {
	envCmd.Flags().SetInterspersed(false)
	rootCmd.AddCommand(envCmd)
}
