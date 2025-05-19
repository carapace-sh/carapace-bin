package cmd

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/git_completer/cmd/common"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace/pkg/traverse"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:     "show",
	Short:   "show branches and their commits",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_main].ID,
}

func init() {
	carapace.Gen(showCmd).Standalone()

	showCmd.Flags().StringS("L", "L", "", "Process line range n,m in file, counting from 1")
	showCmd.Flags().String("decorate", "", "decorate options")
	showCmd.Flags().String("decorate-refs", "", "only decorate refs that match <pattern>")
	showCmd.Flags().String("decorate-refs-exclude", "", "do not decorate refs that match <pattern>")
	showCmd.Flags().Bool("mailmap", false, "alias of --use-mailmap")
	showCmd.Flags().Bool("source", false, "show source")
	showCmd.Flags().Bool("use-mailmap", false, "Use mail map file")
	common.AddBisectionHelperOptions(showCmd)
	common.AddCommitFormattingOptions(showCmd)
	common.AddCommitLimitingOptions(showCmd)
	common.AddCommitOrderingOptions(showCmd)
	common.AddDiffFlags(showCmd)
	common.AddHistorySimplificationOptions(showCmd)
	common.AddObjectTraversalOptions(showCmd)
	rootCmd.AddCommand(showCmd)

	showCmd.Flag("decorate").NoOptDefVal = "short"

	carapace.Gen(showCmd).FlagCompletion(carapace.ActionMap{
		"decorate": carapace.ActionValuesDescribed(
			"short", "do not print ref prefixes",
			"full", "print ref prefixes",
			"auto", "short format when output to terminal",
			"no", "no decoration",
		),
	})

	carapace.Gen(showCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			delimiter := ":"
			if strings.Contains(c.Value, "{") {
				delimiter = "}:"
			}

			return carapace.ActionMultiPartsN(delimiter, 2, func(c carapace.Context) carapace.Action {
				switch len(c.Parts) {
				case 0:
					return git.ActionRefs(git.RefOption{}.Default()).NoSpace()
				default:
					ref := c.Parts[0] + strings.TrimSuffix(delimiter, ":")
					return git.ActionRefFiles(ref).ChdirF(traverse.GitWorkTree)
				}
			})
		}),
	)
}
