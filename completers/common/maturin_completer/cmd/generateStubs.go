package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/maturin_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/completers/common/maturin_completer/cmd/common"
	"github.com/spf13/cobra"
)

var generateStubsCmd = &cobra.Command{
	Use:   "generate-stubs",
	Short: "Autogenerate type stubs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(generateStubsCmd).Standalone()

	generateStubsCmd.Flags().StringP("bindings", "b", "", "Which kind of bindings to use")
	generateStubsCmd.Flags().BoolP("find-interpreter", "f", false, "Find interpreters from the host machine")
	generateStubsCmd.Flags().StringArrayP("interpreter", "i", nil, "The python versions to build wheels for, given as the executables of interpreters")
	generateStubsCmd.Flags().StringP("out", "o", "", "The directory to store the type stubs in")
	common.AddCargoFlags(generateStubsCmd)
	rootCmd.AddCommand(generateStubsCmd)

	if err := generateStubsCmd.MarkFlagRequired("out"); err != nil {
		panic(err)
	}

	carapace.Gen(generateStubsCmd).FlagCompletion(carapace.ActionMap{
		"bindings":    action.ActionBindings(),
		"interpreter": carapace.ActionFiles(),
		"out":         carapace.ActionDirectories(),
	})
}
