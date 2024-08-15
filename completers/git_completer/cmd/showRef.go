package cmd

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var showRefCmd = &cobra.Command{
	Use:     "show-ref",
	Short:   "List references in a local repository",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_interrogator].ID,
}

func init() {
	carapace.Gen(showRefCmd).Standalone()

	showRefCmd.Flags().String("abbrev", "", "abbreviate the object name")
	showRefCmd.Flags().Bool("branches", false, "limit to local branches")
	showRefCmd.Flags().BoolP("dereference", "d", false, "dereference tags into object IDs as well")
	showRefCmd.Flags().String("exclude-existing", "", "make git show-ref act as a filter that reads refs from stdin")
	showRefCmd.Flags().Bool("exists", false, "check whether the given reference exists")
	showRefCmd.Flags().StringP("hash", "s", "", "only show the OID, not the reference name")
	showRefCmd.Flags().Bool("head", false, "show the HEAD reference")
	showRefCmd.Flags().BoolP("quiet", "q", false, "do not print any results to stdout")
	showRefCmd.Flags().Bool("tags", false, "limit to local tags")
	showRefCmd.Flags().Bool("verify", false, "enable stricter reference checking by requiring an exact ref path")
	rootCmd.AddCommand(showRefCmd)

	showRefCmd.Flag("abbrev").NoOptDefVal = " "
	showRefCmd.Flag("exclude-existing").NoOptDefVal = " "
	showRefCmd.Flag("hash").NoOptDefVal = " "

	carapace.Gen(showRefCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			switch {
			case showRefCmd.Flag("exclude-existing").Changed,
				showRefCmd.Flag("exists").Changed && len(c.Args) > 0:
				return carapace.ActionValues()
			}
			return carapace.ActionExecCommand("git", "show-ref")(func(output []byte) carapace.Action {
				lines := strings.Split(string(output), "\n")
				vals := make([]string, 0)

				for _, line := range lines {
					if id, name, ok := strings.Cut(line, " "); ok {
						vals = append(vals, name, id) // TODO limit to branches, tags when flags are set
					}
				}
				return carapace.ActionValuesDescribed(vals...)
			}).MultiParts("/")
		}),
	)
}
