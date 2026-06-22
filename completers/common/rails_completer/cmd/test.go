package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rails_completer/cmd/common"
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Run tests",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(testCmd).Standalone()

	testCmd.Flags().BoolP("help", "h", false, "Show help")
	testCmd.Flags().StringP("name", "n", "", "Filter by test name (regex)")
	testCmd.Flags().Bool("no-plugins", false, "Disable plugins")
	testCmd.Flags().Int("seed", 0, "Random seed")
	testCmd.Flags().BoolP("verbose", "v", false, "Verbose output")
	testCmd.Flags().BoolP("warnings", "w", false, "Run with Ruby warnings")

	common.AddEnvironmentFlag(testCmd, "test")

	carapace.Gen(testCmd).PositionalAnyCompletion(
		carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionFiles("*_test.rb", "*_spec.rb").NoSpace(':')
			case 1:
				return carapace.ActionValues().Usage("line number or range (e.g. 27 or 10-20)")
			default:
				return carapace.ActionValues()
			}
		}),
	)
}
