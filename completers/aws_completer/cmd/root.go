package cmd

import (
	_ "embed"
	"encoding/json"
	"os"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:                "aws",
	Short:              "Universal Command Line Interface for Amazon Web Services",
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func Execute() error {
	return rootCmd.Execute()
}

//go:embed complete.py
var complete string

type completionResult struct {
	Name     string
	HelpText string `json:"help_text"`
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			current := c.CallbackValue
			if c.CallbackValue == "-" {
				current = "--"
			}
			os.Setenv("COMP_LINE", "aws "+strings.Join(append(c.Args, current), " ")) // TODO escape/quote special characters
			return carapace.ActionExecCommand("python", "-c", complete)(func(output []byte) carapace.Action {
				var completionResults []completionResult
				if err := json.Unmarshal(output, &completionResults); err != nil {
					return carapace.ActionMessage(err.Error())
				}
				vals := make([]string, 0, len(completionResults))
				for _, c := range completionResults {
					vals = append(vals, c.Name, c.HelpText)
				}

				if strings.HasPrefix(current, "file://") ||
					strings.HasPrefix(current, "fileb://") {
					return carapace.ActionValuesDescribed(vals...).NoSpace()
				}
				return carapace.ActionValuesDescribed(vals...)
			})
		}),
	)
}
