package cmd

import (
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "carapace",
	Short: "multi-shell multi-command argument completer",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("help", "h", false, "help for carapace")
	rootCmd.Flags().Bool("list", false, "list completers")
	rootCmd.Flags().BoolP("version", "v", false, "version for carapace")

	carapace.Gen(rootCmd).PositionalCompletion(
		ActionCompleters(),
		carapace.ActionValues("bash", "elvish", "fish", "ion", "nushell", "oil", "powershell", "xonsh", "zsh"),
		carapace.ActionValues("_"),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return carapace.ActionValues(c.Args[0])
		}),
	)
}

func ActionCompleters() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("carapace", "--list")(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			re := regexp.MustCompile("^(?P<name>[^ ]+) +(?P<description>.*)$")

			vals := make([]string, 0)
			for _, line := range lines {
				if re.MatchString(line) {
					matched := re.FindStringSubmatch(line)
					//vals = append(vals, matched[1], matched[2])
					vals = append(vals, matched[1])
				}
			}
			return carapace.ActionValues(vals...)
		})
	})
}
