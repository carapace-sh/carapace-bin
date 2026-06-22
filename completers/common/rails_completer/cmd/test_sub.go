package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rails_completer/cmd/common"
	"github.com/spf13/cobra"
)

var testSubCmd = &cobra.Command{
	Use:   "test",
	Short: "Test subcommands (test:all, test:system, etc.)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(testSubCmd).Standalone()

	testSubCmd.Flags().BoolP("help", "h", false, "Show help")
	testSubCmd.Flags().StringP("name", "n", "", "Filter by test name (regex)")
	testSubCmd.Flags().Bool("no-plugins", false, "Disable plugins")
	testSubCmd.Flags().Int("seed", 0, "Random seed")
	testSubCmd.Flags().BoolP("verbose", "v", false, "Verbose output")
	testSubCmd.Flags().BoolP("warnings", "w", false, "Run with Ruby warnings")

	common.AddEnvironmentFlag(testSubCmd, "test")

	carapace.Gen(testSubCmd).PositionalAnyCompletion(
		carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionFiles("*_test.rb", "*_spec.rb").NoSpace(':')
			case 1:
				return carapace.ActionValues().Usage("line number or range")
			default:
				return carapace.ActionValues()
			}
		}),
	)
}
