package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/go_completer/cmd/common"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/golang"
	"github.com/carapace-sh/carapace/pkg/util"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "compile packages and dependencies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(buildCmd).Standalone()
	buildCmd.Flags().SetInterspersed(false)

	buildCmd.Flags().BoolS("json", "json", false, "Emit build output in JSON suitable for automated processing")
	buildCmd.Flags().StringS("o", "o", "", "set output file or directory")
	common.AddBuildFlags(buildCmd)
	rootCmd.AddCommand(buildCmd)

	carapace.Gen(buildCmd).FlagCompletion(carapace.ActionMap{
		"o": carapace.ActionFiles(),
	})

	carapace.Gen(buildCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if util.HasPathPrefix(c.Value) {
				return carapace.ActionDirectories()
			}
			return golang.ActionModules(golang.ModuleOpts{}.Default())
		}),
	)
}
