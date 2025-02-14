package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Show help on command or topic",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(helpCmd).Standalone()
	helpCmd.Flags().SetInterspersed(false)

	rootCmd.AddCommand(helpCmd)

	carapace.Gen(helpCmd).PositionalCompletion(
		carapace.Batch(
			carapace.ActionCallback(func(c carapace.Context) carapace.Action {
				vals := make([]string, 0)
				for _, cmd := range rootCmd.Commands() {
					if !cmd.Hidden && cmd != helpCmd {
						vals = append(vals, cmd.Name(), cmd.Short)
					}
				}
				return carapace.ActionValuesDescribed(vals...)
			}).Style(style.Blue).Tag("commands"),
			carapace.ActionValuesDescribed(
				"buildconstraint", "build constraints",
				"buildjson", "build -json encoding",
				"buildmode", "build modes",
				"c", "calling between Go and C",
				"cache", "build and test caching",
				"environment", "environment variables",
				"filetype", "file types",
				"goauth", "GOAUTH environment variable",
				"go.mod", "the go.mod file",
				"gopath", "GOPATH environment variable",
				"goproxy", "module proxy protocol",
				"importpath", "import path syntax",
				"modules", "modules, module versions, and more",
				"module-auth", "module authentication using go.sum",
				"packages", "package lists and patterns",
				"private", "configuration for downloading non-public code",
				"testflag", "testing flags",
				"testfunc", "testing functions",
				"vcs", "controlling version control with GOVCS",
			).Tag("topics"),
		).ToA(),
	)
}
