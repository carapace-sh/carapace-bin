package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dscv3Cmd = &cobra.Command{
	Use:   "dscv3",
	Short: "DSC v3 resource commands",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dscv3Cmd).Standalone()

	dscv3Cmd.PersistentFlags().String("manifest", "", "Get the resource manifest")
	dscv3Cmd.PersistentFlags().StringP("output", "o", "", "Directory where the results are to be written")
	rootCmd.AddCommand(dscv3Cmd)

	carapace.Gen(dscv3Cmd).FlagCompletion(carapace.ActionMap{
		"manifest": carapace.ActionFiles(),
		"output":   carapace.ActionFiles(),
	})
}
