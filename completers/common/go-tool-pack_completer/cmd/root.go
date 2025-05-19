package cmd

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-tool-pack",
	Short: "Pack is a simple version of the traditional Unix ar tool",
	Long:  "https://pkg.go.dev/cmd/pack",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"c", "append files (from the file system) to a new archive",
			"p", "print files from the archive",
			"r", "append files (from the file system) to the archive",
			"t", "list files from the archive",
			"x", "extract files from the archive",
		),
		carapace.ActionFiles(".a"),
	)

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			switch c.Args[0] {
			case "c", "r":
				return carapace.ActionFiles()
			case "x", "p":
				return carapace.ActionExecCommand("go", "tool", "pack", "t", c.Args[1])(func(output []byte) carapace.Action {
					lines := strings.Split(string(output), "\n")
					return carapace.ActionValues(lines[:len(lines)-1]...).StyleF(style.ForPathExt).MultiParts("/")
				})
			}
			return carapace.ActionValues()
		}),
	)
}
