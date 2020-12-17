package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var fmtCmd = &cobra.Command{
	Use:   "fmt",
	Short: "format all bin and lib files of the current crate",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fmtCmd).Standalone()

	fmtCmd.Flags().Bool("all", false, "Format all packages (only usable in workspaces)")
	fmtCmd.Flags().BoolP("help", "h", false, "Prints help information")
	fmtCmd.Flags().String("manifest-path", "", "Specify path to Cargo.toml")
	fmtCmd.Flags().String("message-format", "", "Specify message-format: short|json|human")
	fmtCmd.Flags().StringP("package", "p", "", "Specify package to format (only usable in workspaces)")
	fmtCmd.Flags().BoolP("quiet", "q", false, "No output printed to stdout")
	fmtCmd.Flags().BoolP("verbose", "v", false, "Use verbose output")
	fmtCmd.Flags().Bool("version", false, "Print rustfmt version and exit")
	rootCmd.AddCommand(fmtCmd)

	carapace.Gen(fmtCmd).FlagCompletion(carapace.ActionMap{
		"manifest-path":  carapace.ActionFiles(".toml"),
		"message-format": carapace.ActionValues("short", "json", "human"),
	})
}
