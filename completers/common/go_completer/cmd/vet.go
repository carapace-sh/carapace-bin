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

	vetCmd.Flags().StringS("vettool", "vettool", "", "select a different analysis tool")
	common.AddPackageBuildFlags(vetCmd)
	rootCmd.AddCommand(vetCmd)

	carapace.Gen(vetCmd).FlagCompletion(carapace.ActionMap{
		"vettool": carapace.ActionFiles(),
	})
}
