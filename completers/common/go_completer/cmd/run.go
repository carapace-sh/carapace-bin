package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/go_completer/cmd/common"
	"github.com/carapace-sh/carapace/pkg/util"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "compile and run Go program",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runCmd).Standalone()
	runCmd.Flags().SetInterspersed(false)

	runCmd.Flags().StringS("exec", "exec", "", "invoke the binary using xprog")
	common.AddBuildFlags(runCmd)
	rootCmd.AddCommand(runCmd)

	carapace.Gen(runCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if util.HasPathPrefix(c.Value) {
				return carapace.ActionDirectories()
			}
			return carapace.ActionValues()
		}),
	)
}
