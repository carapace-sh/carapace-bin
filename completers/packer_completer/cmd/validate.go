package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/packer_completer/cmd/action"
	"github.com/spf13/cobra"
)

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "check that a template is valid",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(validateCmd).Standalone()

	validateCmd.Flags().String("except", "", "Validate all builds other than these.")
	validateCmd.Flags().Bool("machine-readable", false, "Produce machine-readable output.")
	validateCmd.Flags().String("only", "", "Validate only these builds.")
	validateCmd.Flags().Bool("syntax-only", false, "Only check syntax. Do not verify config of the template.")
	validateCmd.Flags().StringArray("var", []string{}, "Variable for templates, can be used multiple times.")
	validateCmd.Flags().String("var-file", "", "JSON or HCL2 file containing user variables.")
	rootCmd.AddCommand(validateCmd)

	carapace.Gen(validateCmd).FlagCompletion(carapace.ActionMap{
		// TODO except/only
		"var": carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
			if len(c.Parts) == 0 && len(c.Args) > 0 {
				return action.ActionVariables(c.Args[0]).Invoke(c).Suffix("=").ToA()
			}
			return carapace.ActionValues()
		}),
		"var-file": carapace.ActionFiles(".json", ".hcl"),
	})

	carapace.Gen(validateCmd).PositionalCompletion(
		carapace.ActionFiles(".json", ".hcl"),
	)
}
