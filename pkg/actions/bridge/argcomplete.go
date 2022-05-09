package bridge

import (
	"os/exec"
	"strconv"
	"strings"

	"github.com/rsteube/carapace"
)

// ActionArgcomplete bridges kislyuk/argcomplete
//   var rootCmd = &cobra.Command{
//   	Use:                "az",
//   	Short:              "Azure Command-Line Interface",
//   	Run:                func(cmd *cobra.Command, args []string) {},
//   	DisableFlagParsing: true,
//   }
//
//   func Execute() error {
//   	return rootCmd.Execute()
//   }
//
//   func init() {
//   	carapace.Gen(rootCmd).Standalone()
//
//   	carapace.Gen(rootCmd).PositionalAnyCompletion(
//   		argcomplete.ActionArgcomplete("az"),
//   	)
//   }
func ActionArgcomplete(command string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if _, err := exec.LookPath(command); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		args := c.Args
		current := c.CallbackValue

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

		compLine := command + " " + strings.Join(append(args, current), " ") // TODO escape/quote special characters
		c.Setenv("_ARGCOMPLETE", "1")
		c.Setenv("_ARGCOMPLETE_DFS", "\t")
		c.Setenv("_ARGCOMPLETE_IFS", "\n")
		c.Setenv("_ARGCOMPLETE_SHELL", "fish")
		c.Setenv("_ARGCOMPLETE_SUPPRESS_SPACE", "1") // TODO needed? relevant for nospace detection?
		// c.Setenv("_ARGCOMPLETE_COMP_WORDBREAKS", " ") // TODO set to space-only for multiparts?
		c.Setenv("_ARGCOMPLETE", "1")
		c.Setenv("COMP_LINE", compLine)
		c.Setenv("COMP_POINT", strconv.Itoa(len(compLine)))
		nospace := false
		a := carapace.ActionExecCommand("sh", "-c", command+" 8>&1 9>&2 1>/dev/null 2>/dev/null")(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
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
				if splitted := strings.SplitN(line, "\t", 2); splitted[0] != "" {
					vals = append(vals, splitted...)
					if len(splitted) < 2 {
						vals = append(vals, "")
					}
				}
			}

			if len(vals) == 0 {
				// fallback to file completions when no values returned
				if index := strings.Index(c.CallbackValue, "="); index > -1 {
					return carapace.ActionFiles().Invoke(carapace.Context{CallbackValue: c.CallbackValue[index+1:]}).ToA()
				}
				return carapace.ActionFiles()
			}
			return carapace.ActionValuesDescribed(vals...)
		}).Invoke(c).Prefix(prefix).ToA() // re-add optarg prefix
		if nospace {
			return a.NoSpace()
		}
		return a
	})
}
