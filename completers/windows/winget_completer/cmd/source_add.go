package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var source_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new source",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(source_addCmd).Standalone()

	source_addCmd.Flags().Bool("accept-source-agreements", false, "Accept all source agreements during source operations")
	source_addCmd.Flags().StringP("arg", "a", "", "Argument given to the source")
	source_addCmd.Flags().Bool("explicit", false, "Excludes a source from discovery unless specified")
	source_addCmd.Flags().String("header", "", "Optional Windows-Package-Manager REST source HTTP header")
	source_addCmd.Flags().StringP("name", "n", "", "Name of the source")
	source_addCmd.Flags().String("trust-level", "", "Trust level of the source")
	source_addCmd.Flags().StringP("type", "t", "", "Type of the source")
	sourceCmd.AddCommand(source_addCmd)

	// TODO flag completion
	carapace.Gen(source_addCmd).FlagCompletion(carapace.ActionMap{
		"arg":         carapace.ActionValues(),
		"header":      carapace.ActionValues(),
		"name":        carapace.ActionValues(),
		"trust-level": carapace.ActionValues("none", "trusted"),
		"type":        carapace.ActionValues(),
	})
}
