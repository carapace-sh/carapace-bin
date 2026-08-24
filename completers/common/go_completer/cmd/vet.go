package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/go_completer/cmd/common"
	"github.com/spf13/cobra"
)

var vetCmd = &cobra.Command{
	Use:   "vet",
	Short: "report likely mistakes in packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vetCmd).Standalone()
	vetCmd.Flags().SetInterspersed(false)

	vetCmd.Flags().StringS("c", "c", "", "display offending line with this many lines of context")
	vetCmd.Flags().BoolS("diff", "diff", false, "print diffs")
	vetCmd.Flags().BoolS("fix", "fix", false, "apply the first fix (if any) for each diagnostic")
	vetCmd.Flags().BoolS("json", "json", false, "emit JSON output")
	vetCmd.Flags().StringS("vettool", "vettool", "", "select a different analysis tool")
	common.AddPackageBuildFlags(vetCmd)
	rootCmd.AddCommand(vetCmd)

	carapace.Gen(vetCmd).FlagCompletion(carapace.ActionMap{
		"vettool": carapace.ActionFiles(),
	})
}
