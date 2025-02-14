package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/golang"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "add dependencies to current module and install them",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(getCmd).Standalone()
	getCmd.Flags().SetInterspersed(false)

	getCmd.Flags().BoolS("d", "d", false, "only download the source code needed to build")
	getCmd.Flags().BoolS("insecure", "insecure", false, "permit using insecure schemes such as HTTP")
	getCmd.Flags().BoolS("t", "t", false, "consider modules needed to build tests")
	getCmd.Flags().StringS("tool", "tool", "", "add tool line to for each listed package")
	getCmd.Flags().BoolS("u", "u", false, "update modules providing dependencies")
	addBuildFlags(getCmd)
	rootCmd.AddCommand(getCmd)

	carapace.Gen(getCmd).FlagCompletion(carapace.ActionMap{
		"tool": carapace.Batch(
			golang.ActionModuleSearch(),
			carapace.ActionMultiPartsN("@", 2, func(c carapace.Context) carapace.Action {
				switch len(c.Parts) {
				case 0:
					return carapace.ActionValues()
				default:
					return carapace.ActionStyledValuesDescribed("none", "remove tool line", style.Carapace.KeywordNegative)
				}
			}),
		).ToA(),
	})

	carapace.Gen(getCmd).PositionalCompletion(
		golang.ActionModuleSearch(),
	)
}
