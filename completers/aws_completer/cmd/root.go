package cmd

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:                "aws",
	Short:              "Universal Command Line Interface for Amazon Web Services",
	Long:               "https://aws.amazon.com/cli/",
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
			if c.Value == "-" {
				c.Value += "-" // no shorthand flags so expand to longhand first (which is needed for the completer)
			}

			c.Setenv("COMP_LINE", "aws "+strings.Join(append(c.Args, c.Value), " ")) // TODO escape/quote special characters
			return carapace.ActionExecCommand("aws_completer")(func(output []byte) carapace.Action {
				lines := strings.Split(string(output), "\n")
				if lines[0] == "" {
					return carapace.ActionValues()
				}
				for index, line := range lines {
					if strings.HasSuffix(line, " ") {
						lines[index] = strings.TrimSuffix(line, " ") // v1 has space suffix
					}
				}
				a := carapace.ActionValues(lines[:len(lines)-1]...)
				switch {
				case strings.HasPrefix(c.Value, "file://"):
					a = a.NoSpace('/').StyleF(func(s string, sc style.Context) string {
						return style.ForPath(strings.TrimPrefix(s, "file://"), sc)
					})
				case strings.HasPrefix(c.Value, "fileb://"):
					a = a.NoSpace('/').StyleF(func(s string, sc style.Context) string {
						return style.ForPath(strings.TrimPrefix(s, "fileb://"), sc)
					})
				}
				return a
			}).Invoke(c).ToA()
		}),
	)
}
