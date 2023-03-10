package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Show help on command or topic",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(helpCmd).Standalone()

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
				"buildmode", "build modes",
				"c", "calling between Go and C",
				"cache", "build and test caching",
				"environment", "environment variables",
				"filetype", "file types",
				"go.mod", "the go.mod file",
				"gopath", "GOPATH environment variable",
				"gopath-get", "legacy GOPATH go get",
				"goproxy", "module proxy protocol",
				"importpath", "import path syntax",
				"modules", "modules, module versions, and more",
				"module-get", "module-aware go get",
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
