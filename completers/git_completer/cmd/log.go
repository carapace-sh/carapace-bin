package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/rsteube/carapace-bin/pkg/util"
	"github.com/spf13/cobra"
)

var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Show commit logs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logCmd).Standalone()

	logCmd.Flags().StringS("L", "L", "", "Process line range n,m in file, counting from 1")
	logCmd.Flags().String("decorate", "", "decorate options")
	logCmd.Flags().String("decorate-refs", "", "only decorate refs that match <pattern>")
	logCmd.Flags().String("decorate-refs-exclude", "", "do not decorate refs that match <pattern>")
	logCmd.Flags().Bool("mailmap", false, "alias of --use-mailmap")
	logCmd.Flags().BoolP("quiet", "q", false, "suppress diff output")
	logCmd.Flags().Bool("source", false, "show source")
	logCmd.Flags().Bool("use-mailmap", false, "Use mail map file")
	rootCmd.AddCommand(logCmd)

	logCmd.Flag("decorate").NoOptDefVal = "short"

	carapace.Gen(logCmd).FlagCompletion(carapace.ActionMap{
		"decorate": carapace.ActionValuesDescribed(
			"short", "do not print ref prefixes",
			"full", "print ref prefixes",
			"auto", "short format when output to terminal",
			"no", "no decoration",
		),
	})

	carapace.Gen(logCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if util.HasPathPrefix(c.CallbackValue) {
				return carapace.ActionFiles()
			} else {
				return carapace.ActionMultiParts("...", func(c carapace.Context) carapace.Action {
					if len(c.Parts) < 2 {
						return git.ActionRefs(git.RefOption{}.Default())
					} else {
						return carapace.ActionValues()
					}
				})
			}
		}),
	)
}
