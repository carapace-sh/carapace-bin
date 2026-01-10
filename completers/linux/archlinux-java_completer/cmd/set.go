package cmd

import (
	"regexp"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Force <JAVA_ENV> as default",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setCmd).Standalone()

	rootCmd.AddCommand(setCmd)

	carapace.Gen(setCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return carapace.ActionExecCommand("archlinux-java", "status")(func(output []byte) carapace.Action {
				lines := strings.Split(string(output), "\n")
				r := regexp.MustCompile(`^\s+(\S+)`)

				vals := make([]string, 0)

				for _, line := range lines {
					if matches := r.FindStringSubmatch(line); matches != nil {
						vals = append(vals, matches[1])
					}
				}
				return carapace.ActionValues(vals...)
			})
		}),
	)
}
