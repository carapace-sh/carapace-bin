package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var unresolvedReferencesCmd = &cobra.Command{
	Use:   "unresolved-references",
	Short: "Report unresolved references",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unresolvedReferencesCmd).Standalone()

	unresolvedReferencesCmd.Flags().Bool("disable-build-scripts", false, "Don't run build scripts or load `OUT_DIR` values by running `cargo check` before analysis.")
	unresolvedReferencesCmd.Flags().Bool("disable-proc-macros", false, "Don't use expand proc macros.")
	unresolvedReferencesCmd.Flags().String("proc-macro-srv", "", "Run the proc-macro-srv binary at the specified path.")
	rootCmd.AddCommand(unresolvedReferencesCmd)

	carapace.Gen(unresolvedReferencesCmd).FlagCompletion(carapace.ActionMap{
		"proc-macro-srv": carapace.ActionFiles(),
	})

	carapace.Gen(unresolvedReferencesCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
