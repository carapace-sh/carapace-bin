package cmd

import (
	"os/exec"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var toolCmd = &cobra.Command{
	Use:   "tool",
	Short: "run specified go tool",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(toolCmd).Standalone()

	toolCmd.Flags().BoolS("n", "n", false, "only print the command that would be executed")
	rootCmd.AddCommand(toolCmd)

	carapace.Gen(toolCmd).PositionalCompletion(
		ActionTools(),
	)
}

func ActionTools() carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		if output, err := exec.Command("go", "tool").Output(); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			lines := strings.Split(string(output), "\n")
			return carapace.ActionValues(lines[:len(lines)-1]...)
		}
	})
}
