package cmd

import (
	"os/exec"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "carapace",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Root()
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("help", "h", false, "help for carapace")
	rootCmd.Flags().Bool("list", false, "list completers")

	carapace.Gen(rootCmd).PositionalCompletion(
		ActionCompleters(),
		carapace.ActionValues("bash", "elvish", "fish", "oil", "powershell", "xonsh", "zsh"),
	)
}

func ActionCompleters() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if output, err := exec.Command("carapace", "--list").Output(); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			lines := strings.Split(string(output), "\n")
			return carapace.ActionValues(lines[:len(lines)-1]...)
		}
	})
}
