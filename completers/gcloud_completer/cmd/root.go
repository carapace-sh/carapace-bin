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

			compLine := "gcloud " + strings.Join(append(c.Args, c.CallbackValue), " ") // TODO escape/quote special characters
			os.Setenv("COMP_LINE", compLine)
			os.Setenv("COMP_POINT", strconv.Itoa(len(compLine)))
			os.Setenv("COMP_WORDBREAKS", " ")
			os.Setenv("_ARGCOMPLETE", "1")
			os.Setenv("CLOUDSDK_COMPONENT_MANAGER_DISABLE_UPDATE_CHECK", "1")
			//os.Setenv("IFS", "\n")
			return carapace.ActionExecCommand("sh", "-c", "gcloud 8>&1 9>&2 1>/dev/null 2>/dev/null")(func(output []byte) carapace.Action {
				lines := strings.Split(string(output), "\v")
				if lines[0] == "" {
					return carapace.ActionValues()
				}
				return carapace.ActionValues(lines[:len(lines)-1]...) // TODO nospace
			})
		}),
	)
}
