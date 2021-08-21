package cmd

import (
	"os"
	"strconv"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:                "gcloud",
	Short:              "manage Google Cloud Platform resources and developer workflow",
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			// TODO patch user@instance and --flag=optarg as in gcloud completion script

			args := c.Args
			current := c.CallbackValue
			if c.CallbackValue == "-" {
				return carapace.ActionValues("--").NoSpace() // seems shorthand flags aren't completed anyway so expand to longhand first
			}

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

			compLine := "gcloud " + strings.Join(append(args, current), " ") // TODO escape/quote special characters
			os.Setenv("COMP_LINE", compLine)
			os.Setenv("COMP_POINT", strconv.Itoa(len(compLine)))
			os.Setenv("_ARGCOMPLETE", "1")
			os.Setenv("CLOUDSDK_COMPONENT_MANAGER_DISABLE_UPDATE_CHECK", "1")
			nospace := false
			a := carapace.ActionExecCommand("sh", "-c", "gcloud 8>&1 9>&2 1>/dev/null 2>/dev/null")(func(output []byte) carapace.Action {
				lines := strings.Split(string(output), "\v")
				if lines[0] == "" {
					return carapace.ActionValues()
				}

				for _, line := range lines[:len(lines)-1] {
					if strings.HasSuffix(line, "=") ||
						strings.HasSuffix(line, "/") ||
						strings.HasSuffix(line, ",") {
						nospace = true
					}
				}

				return carapace.ActionValues(lines[:len(lines)-1]...) // TODO nospace
			}).Invoke(c).Prefix(prefix).ToA() // re-add optarg prefix
			if nospace {
				return a.NoSpace()
			}
			return a
		}),
	)
}
