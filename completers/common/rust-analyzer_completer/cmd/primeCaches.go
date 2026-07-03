package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var primeCachesCmd = &cobra.Command{
	Use:   "prime-caches",
	Short: "Prime caches, as rust-analyzer does typically at startup in interactive sessions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(primeCachesCmd).Standalone()

	primeCachesCmd.Flags().Bool("disable-build-scripts", false, "Don't run build scripts or load `OUT_DIR` values by running `cargo check` before analysis.")
	primeCachesCmd.Flags().Bool("disable-proc-macros", false, "Don't use expand proc macros.")
	primeCachesCmd.Flags().String("num-threads", "", "The number of threads to use. Defaults to the number of physical cores.")
	primeCachesCmd.Flags().String("proc-macro-srv", "", "Run the proc-macro-srv binary at the specified path.")
	rootCmd.AddCommand(primeCachesCmd)

	carapace.Gen(primeCachesCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
