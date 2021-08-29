package cmd

import (
	"os"
	"strconv"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:                "az",
	Short:              "Azure Command-Line Interface",
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		// TODO extract to general argcomplete action
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			args := c.Args
			current := c.CallbackValue
			//if c.CallbackValue == "-" {
			// TODO
			// return carapace.ActionValues("--").NoSpace() // seems shorthand flags aren't completed anyway so expand to longhand first
			//}

			prefix := ""
			if strings.HasPrefix(current, "--") {
				if strings.Contains(current, "=") { // optarg flag which is handled as normal arg by the completer
					splitted := strings.SplitN(current, "=", 2)
					prefix = splitted[0] + "="
					args = append(args, splitted[0]) // add flag as arg
					current = ""                     // seem partial optarg value isn't completed

				} else {
					current = "--" // seems partial flag names aren't completed so get all
				}
			}

			compLine := "az " + strings.Join(append(args, current), " ") // TODO escape/quote special characters
			os.Setenv("_ARGCOMPLETE", "1")
			os.Setenv("_ARGCOMPLETE_DFS", "\t")
			os.Setenv("_ARGCOMPLETE_IFS", "\n")
			os.Setenv("_ARGCOMPLETE_SHELL", "fish")
			os.Setenv("_ARGCOMPLETE_SUPPRESS_SPACE", "1") // TODO needed?
			os.Setenv("_ARGCOMPLETE", "1")
			os.Setenv("COMP_LINE", compLine)
			os.Setenv("COMP_POINT", strconv.Itoa(len(compLine)))
			nospace := false
			a := carapace.ActionExecCommand("sh", "-c", "az 8>&1 9>&2 1>/dev/null 2>/dev/null")(func(output []byte) carapace.Action {
				lines := strings.Split(string(output), "\n")
				if lines[0] == "" {
					return carapace.ActionValues()
				}

				vals := make([]string, 0)
				isFlag := strings.HasPrefix(c.CallbackValue, "-")
				for _, line := range lines[:len(lines)-1] {
					if !isFlag && strings.HasPrefix(line, "-") {
						continue
					}
					if strings.HasSuffix(line, "=") ||
						strings.HasSuffix(line, "/") ||
						strings.HasSuffix(line, ",") {
						nospace = true
					}
					splitted := strings.SplitN(line, "\t", 2)
					vals = append(vals, splitted...)
					if len(splitted) < 2 {
						vals = append(vals, "")
					}
				}

				return carapace.ActionValuesDescribed(vals...)
			}).Invoke(c).Prefix(prefix).ToA() // re-add optarg prefix
			if nospace {
				return a.NoSpace()
			}
			return a
		}),
	)
}
