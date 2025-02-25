package cmd

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var cherryCmd = &cobra.Command{
	Use:     "cherry",
	Short:   "Find commits yet to be applied to upstream",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_interrogator].ID,
}

func init() {
	carapace.Gen(cherryCmd).Standalone()

	cherryCmd.Flags().String("abbrev", "", "use <n> digits to display object names")
	cherryCmd.Flags().BoolP("verbose", "v", false, "be verbose")
	rootCmd.AddCommand(cherryCmd)

	cherryCmd.Flag("abbrev").NoOptDefVal = " "

	carapace.Gen(cherryCmd).PositionalCompletion(
		git.ActionRefs(git.RefOption{}.Default()),
		git.ActionRefs(git.RefOption{}.Default()),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return carapace.ActionExecCommand("git", "cherry", "--abbrev", "--verbose", c.Args[0], c.Args[1])(func(output []byte) carapace.Action {
				lines := strings.Split(string(output), "\n")
				vals := make([]string, 0)
				for _, line := range lines {
					if fields := strings.Fields(line); len(fields) > 2 {
						vals = append(vals, fields[1], strings.Join(fields[2:], " "))
					}
				}
				return carapace.ActionValuesDescribed(vals...)
			})
		}),
	)

	carapace.Gen(cherryCmd).DashAnyCompletion(
		carapace.ActionPositional(cherryCmd),
	)
}
