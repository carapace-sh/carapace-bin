package cmd

import (
	_ "embed"
	"encoding/json"
	"os"
	"os/exec"
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
			path, err := exec.LookPath("aws_completer")
			if err != nil {
				return carapace.ActionMessage(err.Error())
			}

			info, err := os.Stat(path)
			if err != nil {
				return carapace.ActionMessage(err.Error())
			}
			if info.Size() > 100000 { // python version is ~4 kb and compiled is >4 MB
				return actionBinaryCompleter()
			} else {
				return actionPythonCompleter()
			}
		}),
	)
}

func actionBinaryCompleter() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		current := c.CallbackValue
		if c.CallbackValue == "-" {
			current = "--"
		}
		os.Setenv("COMP_LINE", "aws "+strings.Join(append(c.Args, current), " ")) // TODO escape/quote special characters
		return carapace.ActionExecCommand("aws_completer")(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			a := carapace.ActionValues(lines[:len(lines)-1]...)
			if strings.HasPrefix(current, "file://") ||
				strings.HasPrefix(current, "fileb://") {
				return a.NoSpace()
			}
			return a
		})
	})
}

func actionPythonCompleter() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
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

			prefix := ""
			if index := strings.LastIndexAny(c.CallbackValue, "=,/"); index > -1 {
				prefix = c.CallbackValue[:index+1]
			}

			nospace := false
			vals := make([]string, 0, len(completionResults))
			for _, c := range completionResults {
				vals = append(vals, strings.TrimPrefix(c.Name, prefix), c.HelpText)
				nospace = nospace || strings.ContainsAny(c.Name, "=,/")
			}

			if strings.HasPrefix(current, "file://") ||
				strings.HasPrefix(current, "fileb://") {
				return carapace.ActionValuesDescribed(vals...).NoSpace()
			}
			a := carapace.ActionValuesDescribed(vals...).Invoke(c).Prefix(prefix).ToA()
			if nospace {
				return a.NoSpace()
			}
			return a
		})
	})
}
